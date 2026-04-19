package ifc

import (
	"errors"
	"path/filepath"
	"reflect"
	"testing"
)

func TestValidateRejectsUnknownVersion(t *testing.T) {
	t.Parallel()

	doc, err := ParseBytes([]byte(`
ISO-10303-21;
HEADER;
FILE_SCHEMA(('NOT_A_REAL_IFC'));
ENDSEC;
DATA;
#1=IFCORGANIZATION($,'Code-Hex',$,$,$);
ENDSEC;
END-ISO-10303-21;
`))
	if err != nil {
		t.Fatal(err)
	}

	_, err = Validate(doc)
	if !errors.Is(err, ErrValidationVersionRequired) {
		t.Fatalf("Validate error = %v, want ErrValidationVersionRequired", err)
	}
}

func TestValidateHeaderSchemaMismatchFinding(t *testing.T) {
	t.Parallel()

	doc, err := ParseFile(filepath.Join("testdata", "minimal_ifc2x3.ifc"))
	if err != nil {
		t.Fatal(err)
	}

	report, err := Validate(doc, WithValidationVersion(VersionIFC4ADD2TC1))
	if err != nil {
		t.Fatal(err)
	}

	if report.Valid() {
		t.Fatal("report.Valid() = true, want false")
	}
	if got := findingCodes(report); !reflect.DeepEqual(got, []string{findingCodeSchemaVersionMismatch}) {
		t.Fatalf("finding codes = %#v", got)
	}
}

func TestValidateUnknownEntity(t *testing.T) {
	t.Parallel()

	report := mustValidateBytes(t, `
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCNOTAREALENTITY();
ENDSEC;
END-ISO-10303-21;
`)

	assertSingleFindingCode(t, report, findingCodeEntityUnknown)
}

func TestValidateAbstractEntity(t *testing.T) {
	t.Parallel()

	report := mustValidateBytes(t, `
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCGEOMETRICREPRESENTATIONITEM();
ENDSEC;
END-ISO-10303-21;
`)

	assertSingleFindingCode(t, report, findingCodeEntityAbstract)
}

func TestValidateAttributeCountMismatch(t *testing.T) {
	t.Parallel()

	report := mustValidateBytes(t, `
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCORGANIZATION($,'Code-Hex',$,$);
ENDSEC;
END-ISO-10303-21;
`)

	assertSingleFindingCode(t, report, findingCodeAttributeCount)
}

func TestValidateRequiredOmitted(t *testing.T) {
	t.Parallel()

	report := mustValidateBytes(t, `
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCACTORROLE($,$,$);
ENDSEC;
END-ISO-10303-21;
`)

	assertSingleFindingCode(t, report, findingCodeAttributeRequiredOmit)
}

func TestValidateUnresolvedReference(t *testing.T) {
	t.Parallel()

	report := mustValidateBytes(t, `
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCAPPLICATION(#99,'1.0','go-bim','GOBIM');
ENDSEC;
END-ISO-10303-21;
`)

	assertSingleFindingCode(t, report, findingCodeReferenceUnresolved)
}

func TestValidateEnumeration(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()

		report := mustValidateBytes(t, `
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCACTORROLE(.ARCHITECT.,$,$);
ENDSEC;
END-ISO-10303-21;
`)

		if !report.Valid() {
			t.Fatalf("report.Valid() = false, findings = %#v", report.Findings)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()

		report := mustValidateBytes(t, `
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCACTORROLE(.NOTREAL.,$,$);
ENDSEC;
END-ISO-10303-21;
`)

		assertSingleFindingCode(t, report, findingCodeEnumInvalid)
	})
}

func TestValidateSelect(t *testing.T) {
	t.Parallel()

	t.Run("valid nested select", func(t *testing.T) {
		t.Parallel()

		report := mustValidateBytes(t, `
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCPROPERTYSINGLEVALUE('Display Name',$,IFCTEXT('ok'),$);
ENDSEC;
END-ISO-10303-21;
`)

		if !report.Valid() {
			t.Fatalf("report.Valid() = false, findings = %#v", report.Findings)
		}
	})

	t.Run("invalid bare scalar", func(t *testing.T) {
		t.Parallel()

		report := mustValidateBytes(t, `
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCPROPERTYSINGLEVALUE('Display Name',$,'ok',$);
ENDSEC;
END-ISO-10303-21;
`)

		assertSingleFindingCode(t, report, findingCodeSelectInvalid)
	})
}

func TestValidateAggregateCardinalityAndUniqueness(t *testing.T) {
	t.Parallel()

	t.Run("cardinality", func(t *testing.T) {
		t.Parallel()

		report := mustValidateBytes(t, `
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCUNITASSIGNMENT(());
ENDSEC;
END-ISO-10303-21;
`)

		assertSingleFindingCode(t, report, findingCodeAggregateCardinality)
	})

	t.Run("uniqueness", func(t *testing.T) {
		t.Parallel()

		report := mustValidateBytes(t, `
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCSIUNIT(*,.LENGTHUNIT.,$,.METRE.);
#2=IFCUNITASSIGNMENT((#1,#1));
ENDSEC;
END-ISO-10303-21;
`)

		assertSingleFindingCode(t, report, findingCodeAggregateUnique)
	})
}

func TestValidateInverseCardinality(t *testing.T) {
	t.Parallel()

	t.Run("valid inherited inverse", func(t *testing.T) {
		t.Parallel()

		report := mustValidateBytes(t, `
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCORGANIZATION($,'Code-Hex',$,$,$);
#2=IFCAPPLICATION(#1,'1.0','go-bim','GOBIM');
#3=IFCPERSON($,'Tester',$,$,$,$,$,$);
#4=IFCPERSONANDORGANIZATION(#3,#1,$);
#5=IFCOWNERHISTORY(#4,#2,$,.ADDED.,$,$,$,1704067200);
#6=IFCCARTESIANPOINT((0.,0.,0.));
#7=IFCDIRECTION((1.,0.,0.));
#8=IFCDIRECTION((0.,0.,1.));
#9=IFCAXIS2PLACEMENT3D(#6,#8,#7);
#10=IFCLOCALPLACEMENT($,#9);
#11=IFCOPENINGELEMENT('8f1A2B3C4D5E6F7G8H9I0J',#5,'Door Opening',$,'Opening for Door',#10,$,$,.OPENING.);
#12=IFCWALL('7f1A2B3C4D5E6F7G8H9I0J',#5,'Wall',$,$,#10,$,$,.STANDARD.);
#13=IFCRELVOIDSELEMENT('9f1A2B3C4D5E6F7G8H9I0J',#5,'Wall void',$,#12,#11);
ENDSEC;
END-ISO-10303-21;
`)

		if !report.Valid() {
			t.Fatalf("report.Valid() = false, findings = %#v", report.Findings)
		}
	})

	t.Run("invalid missing relationship", func(t *testing.T) {
		t.Parallel()

		report := mustValidateBytes(t, `
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCORGANIZATION($,'Code-Hex',$,$,$);
#2=IFCAPPLICATION(#1,'1.0','go-bim','GOBIM');
#3=IFCPERSON($,'Tester',$,$,$,$,$,$);
#4=IFCPERSONANDORGANIZATION(#3,#1,$);
#5=IFCOWNERHISTORY(#4,#2,$,.ADDED.,$,$,$,1704067200);
#6=IFCCARTESIANPOINT((0.,0.,0.));
#7=IFCDIRECTION((1.,0.,0.));
#8=IFCDIRECTION((0.,0.,1.));
#9=IFCAXIS2PLACEMENT3D(#6,#8,#7);
#10=IFCLOCALPLACEMENT($,#9);
#11=IFCOPENINGELEMENT('8f1A2B3C4D5E6F7G8H9I0J',#5,'Door Opening',$,'Opening for Door',#10,$,$,.OPENING.);
ENDSEC;
END-ISO-10303-21;
`)

		assertSingleFindingCode(t, report, findingCodeInverseCardinality)
	})
}

func TestValidateMaxFindings(t *testing.T) {
	t.Parallel()

	doc, err := ParseBytes([]byte(`
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCACTORROLE($,$,$);
#2=IFCAPPLICATION(#99,'1.0','go-bim','GOBIM');
ENDSEC;
END-ISO-10303-21;
`))
	if err != nil {
		t.Fatal(err)
	}

	report, err := Validate(doc, WithMaxFindings(1))
	if err != nil {
		t.Fatal(err)
	}

	if got := len(report.Findings); got != 1 {
		t.Fatalf("len(report.Findings) = %d, want 1", got)
	}
	if report.Findings[0].Code != findingCodeAttributeRequiredOmit {
		t.Fatalf("first finding code = %q", report.Findings[0].Code)
	}
}

func TestValidateReportValid(t *testing.T) {
	t.Parallel()

	doc, err := ParseFile(filepath.Join("testdata", "minimal_ifc4.ifc"))
	if err != nil {
		t.Fatal(err)
	}

	report, err := Validate(doc)
	if err != nil {
		t.Fatal(err)
	}
	if !report.Valid() {
		t.Fatalf("report.Valid() = false, findings = %#v", report.Findings)
	}
}

func mustValidateBytes(t *testing.T, raw string, opts ...ValidateOption) *Report {
	t.Helper()

	doc, err := ParseBytes([]byte(raw))
	if err != nil {
		t.Fatalf("ParseBytes: %v", err)
	}
	report, err := Validate(doc, opts...)
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	return report
}

func assertSingleFindingCode(t *testing.T, report *Report, want string) {
	t.Helper()

	if report.Valid() {
		t.Fatalf("report.Valid() = true, want false")
	}
	if got := findingCodes(report); !reflect.DeepEqual(got, []string{want}) {
		t.Fatalf("finding codes = %#v, want %#v", got, []string{want})
	}
}

func findingCodes(report *Report) []string {
	codes := make([]string, 0, len(report.Findings))
	for _, finding := range report.Findings {
		codes = append(codes, finding.Code)
	}
	return codes
}
