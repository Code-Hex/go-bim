package ids_test

import (
	"errors"
	"io/fs"
	"path/filepath"
	"reflect"
	"slices"
	"testing"

	"github.com/Code-Hex/go-bim/ids"
)

const idsNamespace = "http://standards.buildingsmart.org/IDS"

func TestOfficialVersions(t *testing.T) {
	t.Parallel()

	if got := ids.OfficialVersions(); !reflect.DeepEqual(got, []ids.Version{ids.Version1_0_0}) {
		t.Fatalf("OfficialVersions() = %#v, want %#v", got, []ids.Version{ids.Version1_0_0})
	}
}

func TestVersionFromSchemaLocation(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		raw  string
		want ids.Version
	}{
		{
			name: "official 1.0",
			raw:  "http://standards.buildingsmart.org/IDS/1.0/ids.xsd",
			want: ids.Version1_0_0,
		},
		{
			name: "candidate 0.9.7",
			raw:  "http://standards.buildingsmart.org/IDS/0.9.7/ids.xsd",
			want: ids.Version0_9_7Candidate,
		},
		{
			name: "unknown",
			raw:  "https://example.com/ids.xsd",
			want: ids.VersionUnknown,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := ids.VersionFromSchemaLocation(tc.raw); got != tc.want {
				t.Fatalf("VersionFromSchemaLocation(%q) = %s, want %s", tc.raw, got, tc.want)
			}
		})
	}
}

func TestParseCapturesAllIDSFields(t *testing.T) {
	t.Parallel()

	doc, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xs="http://www.w3.org/2001/XMLSchema"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info>
    <title>Everything</title>
    <copyright>Copyright</copyright>
    <version>Authoring version</version>
    <description>Info description</description>
    <author>ids@example.com</author>
    <date>2024-06-03</date>
    <purpose>Purpose</purpose>
    <milestone>M3</milestone>
  </info>
  <specifications>
    <specification
      name="Walls"
      ifcVersion="IFC2X3 IFC4 IFC4X3_ADD2"
      identifier="SPEC-001"
      description="Spec description"
      instructions="Top level instructions">
      <applicability minOccurs="0" maxOccurs="0">
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
          <predefinedType>
            <xs:restriction base="xs:string">
              <xs:enumeration value="SOLIDWALL"/>
              <xs:pattern value="S.*"/>
            </xs:restriction>
          </predefinedType>
        </entity>
        <partOf relation="IFCRELAGGREGATES">
          <entity>
            <name><simpleValue>IFCBUILDING</simpleValue></name>
          </entity>
        </partOf>
        <classification>
          <value><simpleValue>EF_25_10</simpleValue></value>
          <system><simpleValue>Uniclass 2015</simpleValue></system>
        </classification>
        <attribute>
          <name><simpleValue>Name</simpleValue></name>
          <value><simpleValue>W01</simpleValue></value>
        </attribute>
        <property dataType="IFCLABEL">
          <propertySet><simpleValue>Pset_WallCommon</simpleValue></propertySet>
          <baseName><simpleValue>FireRating</simpleValue></baseName>
          <value>
            <xs:restriction base="xs:string">
              <xs:minLength value="2"/>
              <xs:maxLength value="5"/>
            </xs:restriction>
          </value>
        </property>
        <material>
          <value><simpleValue>concrete</simpleValue></value>
        </material>
      </applicability>
      <requirements description="Requirement description">
        <entity instructions="Keep it a wall">
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
        <partOf relation="IFCRELCONTAINEDINSPATIALSTRUCTURE" cardinality="prohibited" instructions="No spaces">
          <entity>
            <name><simpleValue>IFCSPACE</simpleValue></name>
          </entity>
        </partOf>
        <classification uri="https://example.com/class" cardinality="optional" instructions="Optional classification">
          <value>
            <xs:restriction base="xs:string">
              <xs:enumeration value="EF_25_10"/>
            </xs:restriction>
          </value>
          <system><simpleValue>Uniclass 2015</simpleValue></system>
        </classification>
        <attribute cardinality="required" instructions="Author the name">
          <name><simpleValue>Name</simpleValue></name>
          <value>
            <xs:restriction base="xs:string">
              <xs:length value="3"/>
            </xs:restriction>
          </value>
        </attribute>
        <property dataType="IFCLABEL" uri="https://example.com/property" cardinality="optional" instructions="Author fire rating">
          <propertySet><simpleValue>Pset_WallCommon</simpleValue></propertySet>
          <baseName><simpleValue>FireRating</simpleValue></baseName>
          <value>
            <xs:restriction base="xs:string">
              <xs:enumeration value="REI30"/>
              <xs:minInclusive value="1"/>
              <xs:maxExclusive value="99"/>
            </xs:restriction>
          </value>
        </property>
        <material uri="https://example.com/material" cardinality="prohibited" instructions="No steel">
          <value><simpleValue>steel</simpleValue></value>
        </material>
      </requirements>
    </specification>
  </specifications>
</ids>`))
	if err != nil {
		t.Fatalf("ParseBytes(): %v", err)
	}

	if doc.Version != ids.Version1_0_0 {
		t.Fatalf("doc.Version = %s, want %s", doc.Version, ids.Version1_0_0)
	}
	if doc.Info.Title != "Everything" {
		t.Fatalf("doc.Info.Title = %q, want Everything", doc.Info.Title)
	}
	if got := *doc.Info.Author; got != "ids@example.com" {
		t.Fatalf("doc.Info.Author = %q, want ids@example.com", got)
	}

	location, ok := doc.SchemaLocation(idsNamespace)
	if !ok {
		t.Fatal("schema location for IDS namespace not found")
	}
	if location.Location != "http://standards.buildingsmart.org/IDS/1.0/ids.xsd" {
		t.Fatalf("schema location = %q", location.Location)
	}

	if got := len(doc.Specifications); got != 1 {
		t.Fatalf("len(doc.Specifications) = %d, want 1", got)
	}

	spec := doc.Specifications[0]
	if got := spec.IFCVersions.Strings(); !reflect.DeepEqual(got, []string{"IFC2X3", "IFC4", "IFC4X3_ADD2"}) {
		t.Fatalf("spec.IFCVersions = %#v", got)
	}
	if got := *spec.Identifier; got != "SPEC-001" {
		t.Fatalf("spec.Identifier = %q, want SPEC-001", got)
	}

	if got := []ids.FacetKind{
		spec.Applicability.Facets[0].Kind(),
		spec.Applicability.Facets[1].Kind(),
		spec.Applicability.Facets[2].Kind(),
		spec.Applicability.Facets[3].Kind(),
		spec.Applicability.Facets[4].Kind(),
		spec.Applicability.Facets[5].Kind(),
	}; !reflect.DeepEqual(got, []ids.FacetKind{
		ids.FacetEntity,
		ids.FacetPartOf,
		ids.FacetClassification,
		ids.FacetAttribute,
		ids.FacetProperty,
		ids.FacetMaterial,
	}) {
		t.Fatalf("applicability facet kinds = %#v", got)
	}

	entity := spec.Applicability.Facets[0].Entity()
	name, ok := entity.Name.SimpleValue()
	if !ok || name != "IFCWALL" {
		t.Fatalf("entity name = %q, want IFCWALL", name)
	}
	if entity.PredefinedType == nil {
		t.Fatal("entity predefined type missing")
	}
	predefinedRestriction, ok := entity.PredefinedType.Restriction()
	if !ok {
		t.Fatal("entity predefined type restriction missing")
	}
	if got := predefinedRestriction.Facets[1]; got.Type != ids.RestrictionPattern || got.Value != "S.*" {
		t.Fatalf("predefinedType restriction facet = %#v", got)
	}

	property := spec.Applicability.Facets[4].Property()
	if property.DataType == nil || *property.DataType != "IFCLABEL" {
		t.Fatalf("property dataType = %#v", property.DataType)
	}
	if property.Value == nil {
		t.Fatal("property value missing")
	}
	propertyRestriction, ok := property.Value.Restriction()
	if !ok {
		t.Fatal("property restriction missing")
	}
	if got := []ids.RestrictionFacetType{
		propertyRestriction.Facets[0].Type,
		propertyRestriction.Facets[1].Type,
	}; !reflect.DeepEqual(got, []ids.RestrictionFacetType{ids.RestrictionMinLength, ids.RestrictionMaxLength}) {
		t.Fatalf("property restriction types = %#v", got)
	}

	if spec.Requirements == nil {
		t.Fatal("spec.Requirements is nil")
	}
	if got := len(spec.Requirements.Facets); got != 6 {
		t.Fatalf("len(spec.Requirements.Facets) = %d, want 6", got)
	}

	classification := spec.Requirements.Facets[2].Classification()
	if classification.URI == nil || *classification.URI != "https://example.com/class" {
		t.Fatalf("classification.URI = %#v", classification.URI)
	}
	if classification.Cardinality == nil || *classification.Cardinality != ids.CardinalityOptional {
		t.Fatalf("classification.Cardinality = %#v", classification.Cardinality)
	}
	if classification.Value == nil {
		t.Fatal("classification value missing")
	}
	if _, ok := classification.Value.Restriction(); !ok {
		t.Fatal("classification restriction missing")
	}

	requirementProperty := spec.Requirements.Facets[4].Property()
	if requirementProperty.URI == nil || *requirementProperty.URI != "https://example.com/property" {
		t.Fatalf("requirementProperty.URI = %#v", requirementProperty.URI)
	}
	if requirementProperty.Cardinality == nil || *requirementProperty.Cardinality != ids.CardinalityOptional {
		t.Fatalf("requirementProperty.Cardinality = %#v", requirementProperty.Cardinality)
	}
	restriction, ok := requirementProperty.Value.Restriction()
	if !ok {
		t.Fatal("requirementProperty restriction missing")
	}
	if got := []ids.RestrictionFacetType{
		restriction.Facets[0].Type,
		restriction.Facets[1].Type,
		restriction.Facets[2].Type,
	}; !reflect.DeepEqual(got, []ids.RestrictionFacetType{
		ids.RestrictionEnumeration,
		ids.RestrictionMinInclusive,
		ids.RestrictionMaxExclusive,
	}) {
		t.Fatalf("requirementProperty restriction types = %#v", got)
	}
	if got := []string{
		restriction.Facets[0].Value,
		restriction.Facets[1].Value,
		restriction.Facets[2].Value,
	}; !reflect.DeepEqual(got, []string{"REI30", "1", "99"}) {
		t.Fatalf("requirementProperty restriction values = %#v", got)
	}
}

func TestParseSupportsCandidateSchemaLocation(t *testing.T) {
	t.Parallel()

	doc, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xs="http://www.w3.org/2001/XMLSchema"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/0.9.7/ids.xsd">
  <info>
    <title>Candidate fixture</title>
  </info>
  <specifications>
    <specification name="Candidate" ifcVersion="IFC4X3">
      <applicability maxOccurs="unbounded">
        <entity>
          <name><simpleValue>IFCBRIDGE</simpleValue></name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if err != nil {
		t.Fatalf("ParseBytes(): %v", err)
	}

	if doc.Version != ids.Version0_9_7Candidate {
		t.Fatalf("doc.Version = %s, want %s", doc.Version, ids.Version0_9_7Candidate)
	}
	if got := doc.Specifications[0].IFCVersions.Strings(); !reflect.DeepEqual(got, []string{"IFC4X3"}) {
		t.Fatalf("spec.IFCVersions = %#v", got)
	}
}

func TestParseWithVersionOverridesSchemaLocation(t *testing.T) {
	t.Parallel()

	doc, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xs="http://www.w3.org/2001/XMLSchema"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/0.9.7/ids.xsd">
  <info>
    <title>Forced version</title>
  </info>
  <specifications>
    <specification name="Forced" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`), ids.WithVersion(ids.Version1_0_0))
	if err != nil {
		t.Fatalf("ParseBytes(): %v", err)
	}

	if doc.Version != ids.Version1_0_0 {
		t.Fatalf("doc.Version = %s, want %s", doc.Version, ids.Version1_0_0)
	}

	location, ok := doc.SchemaLocation(idsNamespace)
	if !ok {
		t.Fatal("schema location for IDS namespace not found")
	}
	if location.Location != "http://standards.buildingsmart.org/IDS/0.9.7/ids.xsd" {
		t.Fatalf("schema location = %q, want 0.9.7 location to stay intact", location.Location)
	}
}

func TestParseRejectsUnsupportedForcedVersion(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`), ids.WithVersion(ids.Version("1.0")))
	if !errors.Is(err, ids.ErrUnsupportedVersion) {
		t.Fatalf("ParseBytes(unsupported forced version) error = %v, want ErrUnsupportedVersion", err)
	}
}

func TestParseRejectsNilOption(t *testing.T) {
	t.Parallel()

	var opt ids.ParseOption
	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`), opt)
	if !errors.Is(err, ids.ErrInvalidOption) {
		t.Fatalf("ParseBytes(nil option) error = %v, want ErrInvalidOption", err)
	}
}

func TestParseAllVendoredReferenceIDSFiles(t *testing.T) {
	t.Parallel()

	root := filepath.Join("reference", "official", "v1.0.0", "Documentation", "testcases")
	paths := []string{}
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".ids" {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("WalkDir(%s): %v", root, err)
	}
	if len(paths) == 0 {
		t.Fatalf("no .ids files found under %s", root)
	}

	slices.Sort(paths)
	for _, path := range paths {
		path := path
		t.Run(filepath.Base(path), func(t *testing.T) {
			doc, err := ids.ParseFile(path)
			if err != nil {
				t.Fatalf("ParseFile(%s): %v", path, err)
			}
			if len(doc.Specifications) == 0 {
				t.Fatalf("ParseFile(%s) returned zero specifications", path)
			}
		})
	}
}

func TestParseRestoresDefaultRequirementCardinality(t *testing.T) {
	t.Parallel()

	doc, err := ids.ParseFile(filepath.Join(
		"reference", "official", "v1.0.0", "Documentation", "testcases", "property",
		"pass-measures_are_used_to_specify_an_ifc_data_type_2_2.ids",
	))
	if err != nil {
		t.Fatalf("ParseFile(): %v", err)
	}

	property := doc.Specifications[0].Requirements.Facets[0].Property()
	if property.Cardinality != nil {
		t.Fatalf("property.Cardinality = %#v, want nil explicit attribute", property.Cardinality)
	}
	if got := property.EffectiveCardinality(); got != ids.CardinalityRequired {
		t.Fatalf("property.EffectiveCardinality() = %q, want %q", got, ids.CardinalityRequired)
	}
}

func TestParseRejectsUnexpectedRoot(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<nope/>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(unexpected root) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsMissingIDSNamespace(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<ids><info><title>Bad</title></info><specifications><specification name="Bad" ifcVersion="IFC4"><applicability maxOccurs="unbounded"><entity><name><simpleValue>IFCWALL</simpleValue></name></entity></applicability></specification></specifications></ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(missing IDS namespace) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsInvalidOfficialIFCVersion(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4X3">
      <applicability maxOccurs="unbounded">
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(invalid official ifcVersion) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsNamespacedSpecificationAttribute(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xmlns:bad="urn:bad"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification bad:name="Bad" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(namespaced specification attribute) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsNamespacedApplicabilityAttribute(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xmlns:bad="urn:bad"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability bad:maxOccurs="unbounded">
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(namespaced applicability attribute) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsNamespacedRequirementsAttribute(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xmlns:bad="urn:bad"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability>
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
      </applicability>
      <requirements bad:description="nope">
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
      </requirements>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(namespaced requirements attribute) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseUnknownVersionStillValidatesIFCVersion(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="NOPE">
      <applicability maxOccurs="unbounded">
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(unknown version invalid ifcVersion) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsInvalidApplicabilityOccurs(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability minOccurs="1" maxOccurs="1">
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(invalid applicability occurs) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsMissingRequiredInfoTitle(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS">
  <info></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(missing info.title) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsOutOfOrderClassificationChildren(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <classification>
          <system><simpleValue>Uniclass 2015</simpleValue></system>
          <value><simpleValue>EF_25_10</simpleValue></value>
        </classification>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(out of order classification children) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsInvalidSimpleTypes(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info>
    <title>Bad</title>
    <author>not-an-email</author>
  </info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <property dataType="IfcLabel">
          <propertySet><simpleValue>Pset_WallCommon</simpleValue></propertySet>
          <baseName><simpleValue>FireRating</simpleValue></baseName>
        </property>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(invalid simple types) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsInvalidDataTypeValueCompatibility(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <property dataType="IFCINTEGER">
          <propertySet><simpleValue>Pset_WallCommon</simpleValue></propertySet>
          <baseName><simpleValue>Foo</simpleValue></baseName>
          <value><simpleValue>42.3</simpleValue></value>
        </property>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(invalid dataType/value compatibility) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseAcceptsStringPropertyValues(t *testing.T) {
	t.Parallel()

	doc, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>OK</title></info>
  <specifications>
    <specification name="OK" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <property dataType="IFCLABEL">
          <propertySet><simpleValue>Pset_WallCommon</simpleValue></propertySet>
          <baseName><simpleValue>Status</simpleValue></baseName>
          <value><simpleValue>EXISTING</simpleValue></value>
        </property>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if err != nil {
		t.Fatalf("ParseBytes(valid string property) = %v", err)
	}
	if got, ok := doc.Specifications[0].Applicability.Facets[0].Property().Value.SimpleValue(); !ok || got != "EXISTING" {
		t.Fatalf("property simpleValue = %q, %v, want EXISTING, true", got, ok)
	}
}

func TestParseAcceptsSupportedPartOfRelations(t *testing.T) {
	t.Parallel()

	cases := []string{"IFCRELVOIDSELEMENT", "IFCRELFILLSELEMENT"}
	for _, relation := range cases {
		relation := relation
		t.Run(relation, func(t *testing.T) {
			t.Parallel()

			_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>OK</title></info>
  <specifications>
    <specification name="OK" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <partOf relation="` + relation + `">
          <entity>
            <name><simpleValue>IFCOPENINGELEMENT</simpleValue></name>
          </entity>
        </partOf>
      </applicability>
    </specification>
  </specifications>
</ids>`))
			if err != nil {
				t.Fatalf("ParseBytes(valid partOf relation %s) = %v", relation, err)
			}
		})
	}
}

func TestParseRejectsUnsupportedOrInvalidRestrictionFacets(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		raw  string
	}{
		{
			name: "unsupported totalDigits",
			raw: `<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xs="http://www.w3.org/2001/XMLSchema"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <property dataType="IFCINTEGER">
          <propertySet><simpleValue>Pset_WallCommon</simpleValue></propertySet>
          <baseName><simpleValue>Foo</simpleValue></baseName>
          <value>
            <xs:restriction base="xs:integer">
              <xs:totalDigits value="2"/>
            </xs:restriction>
          </value>
        </property>
      </applicability>
    </specification>
  </specifications>
</ids>`,
		},
		{
			name: "enumeration lexical mismatch",
			raw: `<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xs="http://www.w3.org/2001/XMLSchema"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <property dataType="IFCINTEGER">
          <propertySet><simpleValue>Pset_WallCommon</simpleValue></propertySet>
          <baseName><simpleValue>Foo</simpleValue></baseName>
          <value>
            <xs:restriction base="xs:integer">
              <xs:enumeration value="foo"/>
            </xs:restriction>
          </value>
        </property>
      </applicability>
    </specification>
  </specifications>
</ids>`,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := ids.ParseBytes([]byte(tc.raw))
			if !errors.Is(err, ids.ErrInvalidDocument) {
				t.Fatalf("ParseBytes(%s) error = %v, want ErrInvalidDocument", tc.name, err)
			}
		})
	}
}

func TestParseRejectsInvalidRealLexicalValue(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <property dataType="IFCREAL">
          <propertySet><simpleValue>Pset_WallCommon</simpleValue></propertySet>
          <baseName><simpleValue>Foo</simpleValue></baseName>
          <value><simpleValue></simpleValue></value>
        </property>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(invalid real lexical value) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsInvalidFacetConfigurations(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
      </applicability>
      <requirements>
        <material cardinality="optional"></material>
      </requirements>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(invalid facet configuration) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsMalformedApplicabilityOrdering(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <property dataType="IFCLABEL">
          <propertySet><simpleValue>Pset_WallCommon</simpleValue></propertySet>
          <baseName><simpleValue>Foo</simpleValue></baseName>
        </property>
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(malformed applicability ordering) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsDuplicateValueBranches(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xs="http://www.w3.org/2001/XMLSchema"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability>
        <entity>
          <name>
            <simpleValue>IFCWALL</simpleValue>
            <xs:restriction base="xs:string">
              <xs:enumeration value="IFCWALL"/>
            </xs:restriction>
          </name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(duplicate idsValue branches) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsInvalidRestrictionAttributes(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		raw  string
	}{
		{
			name: "namespaced base",
			raw: `<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xs="http://www.w3.org/2001/XMLSchema"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xmlns:bad="urn:bad"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability>
        <entity>
          <name>
            <xs:restriction bad:base="xs:string">
              <xs:enumeration value="IFCWALL"/>
            </xs:restriction>
          </name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`,
		},
		{
			name: "missing facet value",
			raw: `<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xs="http://www.w3.org/2001/XMLSchema"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability>
        <entity>
          <name>
            <xs:restriction base="xs:string">
              <xs:enumeration/>
            </xs:restriction>
          </name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := ids.ParseBytes([]byte(tc.raw))
			if !errors.Is(err, ids.ErrInvalidDocument) {
				t.Fatalf("ParseBytes(%s) error = %v, want ErrInvalidDocument", tc.name, err)
			}
		})
	}
}

func TestParseRejectsUnexpectedText(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability>noise</applicability>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(unexpected text) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseRejectsUnsupportedValueChild(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info><title>Bad</title></info>
  <specifications>
    <specification name="Bad" ifcVersion="IFC4">
      <applicability>
        <entity>
          <name>
            <unexpected/>
          </name>
        </entity>
      </applicability>
    </specification>
  </specifications>
</ids>`))
	if !errors.Is(err, ids.ErrInvalidDocument) {
		t.Fatalf("ParseBytes(unsupported idsValue child) error = %v, want ErrInvalidDocument", err)
	}
}

func TestParseEmptyInput(t *testing.T) {
	t.Parallel()

	_, err := ids.ParseBytes(nil)
	if !errors.Is(err, ids.ErrEmptyInput) {
		t.Fatalf("ParseBytes(nil) error = %v, want ErrEmptyInput", err)
	}
}
