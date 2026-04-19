package ifc

import (
	"errors"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseFiles(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name       string
		path       string
		version    Version
		schemaName string
	}{
		{"IFC2X3", filepath.Join("testdata", "minimal_ifc2x3.ifc"), VersionIFC2X3TC1, "IFC2X3"},
		{"IFC4", filepath.Join("testdata", "minimal_ifc4.ifc"), VersionIFC4ADD2TC1, "IFC4"},
		{"IFC4X3", filepath.Join("testdata", "minimal_ifc4x3.ifc"), VersionIFC4X3ADD2, "IFC4X3_ADD2"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			doc, err := ParseFile(tc.path)
			if err != nil {
				t.Fatalf("ParseFile(%s): %v", tc.path, err)
			}

			if doc.Version != tc.version {
				t.Fatalf("doc.Version = %s, want %s", doc.Version, tc.version)
			}

			if got := doc.RawSchemaIdentifiers; !reflect.DeepEqual(got, []string{tc.schemaName}) {
				t.Fatalf("doc.RawSchemaIdentifiers = %#v, want %#v", got, []string{tc.schemaName})
			}

			if got := len(doc.Instances); got != 5 {
				t.Fatalf("len(doc.Instances) = %d, want 5", got)
			}

			property, ok := doc.Instance(4)
			if !ok {
				t.Fatal("instance #4 not found")
			}
			if property.Entity != "IFCPROPERTYSINGLEVALUE" {
				t.Fatalf("instance #4 entity = %q", property.Entity)
			}

			named, ok := property.Arguments[2].(NamedParameter)
			if !ok {
				t.Fatalf("argument[2] type = %T, want NamedParameter", property.Arguments[2])
			}
			if named.Name != "IFCTEXT" {
				t.Fatalf("named.Name = %q, want IFCTEXT", named.Name)
			}
			if len(named.Arguments) != 1 {
				t.Fatalf("len(named.Arguments) = %d, want 1", len(named.Arguments))
			}

			text, ok := named.Arguments[0].(StringParameter)
			if !ok {
				t.Fatalf("IFCTEXT argument type = %T, want StringParameter", named.Arguments[0])
			}
			if text.Decoded != "Ä" {
				t.Fatalf("decoded text = %q, want Ä", text.Decoded)
			}

			schema, err := doc.Schema()
			if err != nil {
				t.Fatalf("doc.Schema(): %v", err)
			}
			placement, ok := doc.Instance(5)
			if !ok {
				t.Fatal("instance #5 not found")
			}
			bound, err := schema.Bind(placement)
			if err != nil {
				t.Fatalf("schema.Bind(#5): %v", err)
			}
			if got := []string{bound[0].Attribute.Name, bound[1].Attribute.Name}; !reflect.DeepEqual(got, []string{"Location", "RefDirection"}) {
				t.Fatalf("bound names = %#v", got)
			}
			if _, ok := bound[1].Value.(OmittedParameter); !ok {
				t.Fatalf("bound[1].Value type = %T, want OmittedParameter", bound[1].Value)
			}
		})
	}
}

func TestDecodeSTEPString(t *testing.T) {
	t.Parallel()

	cases := []struct {
		raw  string
		want string
	}{
		{`plain`, "plain"},
		{`O''Brien`, "O'Brien"},
		{`\S\D`, "Ä"},
		{`\X\C4`, "Ä"},
		{`\X2\00C4\X0\`, "Ä"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.raw, func(t *testing.T) {
			t.Parallel()
			if got := decodeSTEPString(tc.raw); got != tc.want {
				t.Fatalf("decodeSTEPString(%q) = %q, want %q", tc.raw, got, tc.want)
			}
		})
	}
}

func TestParseRejectsDuplicateInstanceIDs(t *testing.T) {
	t.Parallel()

	_, err := ParseBytes([]byte(`
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCORGANIZATION($,$,$,$,$);
#1=IFCORGANIZATION($,$,$,$,$);
ENDSEC;
END-ISO-10303-21;
`))
	if !errors.Is(err, ErrDuplicateInstanceID) {
		t.Fatalf("ParseBytes error = %v, want ErrDuplicateInstanceID", err)
	}
}

func TestParseWithVersionKeepsHeaderSchemaIdentifiers(t *testing.T) {
	t.Parallel()

	doc, err := ParseFile(filepath.Join("testdata", "minimal_ifc4.ifc"), WithVersion(VersionIFC2X3TC1))
	if err != nil {
		t.Fatal(err)
	}

	if doc.Version != VersionIFC2X3TC1 {
		t.Fatalf("doc.Version = %s, want %s", doc.Version, VersionIFC2X3TC1)
	}
	if got := doc.RawSchemaIdentifiers; !reflect.DeepEqual(got, []string{"IFC4"}) {
		t.Fatalf("doc.RawSchemaIdentifiers = %#v, want %#v", got, []string{"IFC4"})
	}
}

func TestParseComplexParameterValues(t *testing.T) {
	t.Parallel()

	doc, err := ParseBytes([]byte(`
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCFAKE("0FAB",((1,2),(3,4)),IFCTEXT('ok'),IFCLABEL('\X\C4'),.ARCHITECT.);
ENDSEC;
END-ISO-10303-21;
`))
	if err != nil {
		t.Fatal(err)
	}

	instance, ok := doc.Instance(1)
	if !ok {
		t.Fatal("instance #1 not found")
	}
	if _, ok := instance.Arguments[0].(BinaryParameter); !ok {
		t.Fatalf("argument[0] type = %T, want BinaryParameter", instance.Arguments[0])
	}

	aggregate, ok := instance.Arguments[1].(AggregateParameter)
	if !ok {
		t.Fatalf("argument[1] type = %T, want AggregateParameter", instance.Arguments[1])
	}
	if len(aggregate.Elements) != 2 {
		t.Fatalf("len(argument[1].Elements) = %d, want 2", len(aggregate.Elements))
	}
	if nested, ok := aggregate.Elements[0].(AggregateParameter); !ok || len(nested.Elements) != 2 {
		t.Fatalf("nested aggregate = %#v", aggregate.Elements[0])
	}

	label, ok := instance.Arguments[3].(NamedParameter)
	if !ok {
		t.Fatalf("argument[3] type = %T, want NamedParameter", instance.Arguments[3])
	}
	if label.Name != "IFCLABEL" {
		t.Fatalf("label.Name = %q, want IFCLABEL", label.Name)
	}
	text, ok := label.Arguments[0].(StringParameter)
	if !ok {
		t.Fatalf("label argument type = %T, want StringParameter", label.Arguments[0])
	}
	if text.Decoded != "Ä" {
		t.Fatalf("decoded label = %q, want Ä", text.Decoded)
	}

	if enum, ok := instance.Arguments[4].(EnumerationParameter); !ok || enum.Value != "ARCHITECT" {
		t.Fatalf("enum = %#v", instance.Arguments[4])
	}
}
