package stbridge

import (
	"bytes"
	"encoding/xml"
	"io"
	"slices"
	"strings"
	"testing"
)

func TestDocumentViewCommon(t *testing.T) {
	doc, err := ParseBytes([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<ST_BRIDGE xmlns="https://www.building-smart.or.jp/dl" version="2.1.0">
  <StbCommon guid="g-1" project_name="demo" app_name="authoring-app" app_version="1.0.0" convert_app_name="converter" convert_app_version="2.0.0" strength_concrete="Fc30" global_offset_X="1.5" global_offset_Y="-2.5" global_offset_Z="3.0" global_rotation="90">
    <StbReinforcementStrengthList />
    <StbConnectionSpecs />
    <StbWeldCommon />
  </StbCommon>
  <StbModel>
    <StbNodes />
    <StbAxes />
    <StbStories />
    <StbMembers />
    <StbSections />
    <StbJoints />
    <StbConnections />
    <StbWeld />
  </StbModel>
  <StbExtensions />
  <StbExportInformation>
    <StbExportPolicy />
    <StbExportPolicy />
    <StbExportLog />
    <StbExportLog />
  </StbExportInformation>
  <StbCalData>
    <StbCalCommon />
    <StbCalLoad />
    <StbCalCondition />
    <StbCalLoadArrangements />
    <StbCalConditionArrangements />
  </StbCalData>
  <StbAnaModels>
    <StbAnaModel id="1" />
    <StbAnaModel id="2" />
  </StbAnaModels>
</ST_BRIDGE>`))
	if err != nil {
		t.Fatalf("ParseBytes() error = %v", err)
	}

	if got := doc.SectionNames(); !slices.Equal(got, KnownSectionNamesFor(Version210)) {
		t.Fatalf("SectionNames() = %v, want %v", got, KnownSectionNamesFor(Version210))
	}

	common := doc.View().Common()
	if !common.Present() {
		t.Fatal("View().Common().Present() = false, want true")
	}

	if got, ok := common.ProjectName(); !ok || got != "demo" {
		t.Fatalf("ProjectName() = (%q, %t), want (%q, true)", got, ok, "demo")
	}
	if got, ok := common.GUID(); !ok || got != "g-1" {
		t.Fatalf("GUID() = (%q, %t), want (%q, true)", got, ok, "g-1")
	}
	if got, ok := common.ConvertAppName(); !ok || got != "converter" {
		t.Fatalf("ConvertAppName() = (%q, %t), want (%q, true)", got, ok, "converter")
	}

	assertFloatAttr(t, "GlobalOffsetX", common.GlobalOffsetX, 1.5)
	assertFloatAttr(t, "GlobalOffsetY", common.GlobalOffsetY, -2.5)
	assertFloatAttr(t, "GlobalOffsetZ", common.GlobalOffsetZ, 3.0)
	assertFloatAttr(t, "GlobalRotation", common.GlobalRotation, 90)

	if common.ReinforcementStrengthList() == nil {
		t.Fatal("ReinforcementStrengthList() = nil, want node")
	}
	if common.ConnectionSpecs() == nil {
		t.Fatal("ConnectionSpecs() = nil, want node")
	}
	if common.WeldCommon() == nil {
		t.Fatal("WeldCommon() = nil, want node")
	}

	model := doc.View().Model()
	if model.Nodes() == nil || model.Axes() == nil || model.Stories() == nil {
		t.Fatal("Model() wrapper missed required children")
	}
	if model.Connections() == nil || model.Weld() == nil {
		t.Fatal("Model() wrapper missed 2.1 children")
	}

	exportInfo := doc.View().ExportInformation()
	if !exportInfo.Present() {
		t.Fatal("ExportInformation() wrapper is not present")
	}
	if got := len(exportInfo.ExportPolicies()); got != 2 {
		t.Fatalf("len(ExportPolicies()) = %d, want 2", got)
	}
	if got := len(exportInfo.ExportLogs()); got != 2 {
		t.Fatalf("len(ExportLogs()) = %d, want 2", got)
	}

	calData := doc.View().CalData()
	if !calData.Present() || calData.CalCommon() == nil || calData.CalConditionArrangements() == nil {
		t.Fatal("CalData() wrapper missed child nodes")
	}

	anaModels := doc.View().AnaModels()
	if got := len(anaModels.Models()); got != 2 {
		t.Fatalf("len(Models()) = %d, want 2", got)
	}
}

func TestNodeTypedAttrHelpers(t *testing.T) {
	doc, err := ParseBytes([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<ST_BRIDGE xmlns="https://www.building-smart.or.jp/dl" version="2.1.0">
  <StbCommon global_offset_X="bad" isFoundation="true" />
</ST_BRIDGE>`))
	if err != nil {
		t.Fatalf("ParseBytes() error = %v", err)
	}

	common := doc.Common()

	if got, ok := common.StringAttr(Attr_global_offset_X); !ok || got != "bad" {
		t.Fatalf("StringAttr() = (%q, %t), want (%q, true)", got, ok, "bad")
	}

	if _, ok, err := common.Float64Attr(Attr_global_offset_X); !ok || err == nil {
		t.Fatalf("Float64Attr() = (_, %t, %v), want (_, true, error)", ok, err)
	}

	if _, ok, err := common.BoolAttr(Attr_isFoundation); !ok || err != nil {
		t.Fatalf("BoolAttr() = (_, %t, %v), want (_, true, nil)", ok, err)
	}

	if _, ok, err := common.Int64Attr(Attr_id); ok || err != nil {
		t.Fatalf("Int64Attr(missing) = (_, %t, %v), want (_, false, nil)", ok, err)
	}
}

func TestSchemaNameHelpers(t *testing.T) {
	schema, err := SchemaFor(Version210)
	if err != nil {
		t.Fatalf("SchemaFor(%s) error = %v", Version210, err)
	}

	root, ok := schema.Element(Element_ST_BRIDGE)
	if !ok {
		t.Fatalf("Element(%q) = missing", Element_ST_BRIDGE)
	}
	if got := root.ChildNames(); !slices.Equal(got, []ElementName{
		Element_StbCommon,
		Element_StbModel,
		Element_StbExtensions,
		Element_StbExportInformation,
		Element_StbCalData,
		Element_StbAnaModels,
	}) {
		t.Fatalf("ST_BRIDGE.ChildNames() = %v", got)
	}

	common, ok := schema.Element(Element_StbCommon)
	if !ok {
		t.Fatalf("Element(%q) = missing", Element_StbCommon)
	}
	if _, ok := common.Attribute(Attr_project_name); !ok {
		t.Fatalf("Attribute(%q) = missing", Attr_project_name)
	}
	if !slices.Contains(common.AttributeNames(), Attr_guid) {
		t.Fatalf("AttributeNames() does not contain %q", Attr_guid)
	}
	if !slices.Contains(common.ChildNames(), Element_StbWeldCommon) {
		t.Fatalf("ChildNames() does not contain %q", Element_StbWeldCommon)
	}

	if got := schema.ElementNames(); len(got) != len(schema.Elements()) {
		t.Fatalf("len(ElementNames()) = %d, want %d", len(got), len(schema.Elements()))
	}
	if got := schema.SimpleTypeNames(); len(got) != len(schema.SimpleTypes()) {
		t.Fatalf("len(SimpleTypeNames()) = %d, want %d", len(got), len(schema.SimpleTypes()))
	}
}

func TestKnownNamesStayAlignedWithSchema(t *testing.T) {
	if !slices.Contains(KnownElementNames(), Element_ST_BRIDGE) {
		t.Fatalf("KnownElementNames() does not contain %q", Element_ST_BRIDGE)
	}
	if !slices.Contains(KnownAttributeNames(), Attr_project_name) {
		t.Fatalf("KnownAttributeNames() does not contain %q", Attr_project_name)
	}
	if !slices.Contains(KnownSimpleTypeNames(), SimpleType_guid) {
		t.Fatalf("KnownSimpleTypeNames() does not contain %q", SimpleType_guid)
	}

	if slices.Contains(KnownSectionNamesFor(Version202), Section_StbExportInformation) {
		t.Fatalf("KnownSectionNamesFor(%s) unexpectedly contains %q", Version202, Section_StbExportInformation)
	}
	if !slices.Contains(KnownSectionNamesFor(Version210), Section_StbExportInformation) {
		t.Fatalf("KnownSectionNamesFor(%s) does not contain %q", Version210, Section_StbExportInformation)
	}

	for _, version := range KnownSchemaVersions() {
		schema, err := SchemaFor(version)
		if err != nil {
			t.Fatalf("SchemaFor(%s) error = %v", version, err)
		}

		wantElements, wantAttributes, wantSimpleTypes := extractVersionedSchemaNames(t, version)
		if got := KnownElementNamesFor(version); !slices.Equal(got, wantElements) {
			t.Fatalf("KnownElementNamesFor(%s) = %v, want %v", version, got, wantElements)
		}
		if got := KnownAttributeNamesFor(version); !slices.Equal(got, wantAttributes) {
			t.Fatalf("KnownAttributeNamesFor(%s) = %v, want %v", version, got, wantAttributes)
		}
		if got := KnownSimpleTypeNamesFor(version); !slices.Equal(got, wantSimpleTypes) {
			t.Fatalf("KnownSimpleTypeNamesFor(%s) = %v, want %v", version, got, wantSimpleTypes)
		}

		root, ok := schema.Element(Element_ST_BRIDGE)
		if !ok {
			t.Fatalf("SchemaFor(%s) is missing %q", version, Element_ST_BRIDGE)
		}
		if got := root.ChildNames(); !slices.Equal(got, KnownSectionNamesFor(version)) {
			t.Fatalf("ST_BRIDGE.ChildNames(%s) = %v, want %v", version, got, KnownSectionNamesFor(version))
		}
	}
}

func assertFloatAttr(t *testing.T, label string, fn func() (float64, bool, error), want float64) {
	t.Helper()

	got, ok, err := fn()
	if err != nil {
		t.Fatalf("%s() error = %v", label, err)
	}
	if !ok || got != want {
		t.Fatalf("%s() = (%v, %t), want (%v, true)", label, got, ok, want)
	}
}

func extractVersionedSchemaNames(t *testing.T, version Version) ([]ElementName, []AttributeName, []SimpleTypeName) {
	t.Helper()

	filename, ok := schemaByVersion[version]
	if !ok {
		t.Fatalf("schemaByVersion[%s] is missing", version)
	}

	data, err := schemaFiles.ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile(%q) error = %v", filename, err)
	}

	elementSet := map[ElementName]struct{}{}
	attributeSet := map[AttributeName]struct{}{}
	simpleTypeSet := map[SimpleTypeName]struct{}{}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("Token(%q) error = %v", filename, err)
		}

		start, ok := token.(xml.StartElement)
		if !ok {
			continue
		}

		switch start.Name.Local {
		case "element":
			collectSchemaName(start.Attr, "name", elementSet)
			collectSchemaName(start.Attr, "ref", elementSet)
		case "attribute":
			collectSchemaName(start.Attr, "name", attributeSet)
			collectSchemaName(start.Attr, "ref", attributeSet)
		case "simpleType":
			collectSchemaName(start.Attr, "name", simpleTypeSet)
		}
	}

	return sortedNames(elementSet), sortedNames(attributeSet), sortedNames(simpleTypeSet)
}

func collectSchemaName[T ~string](attrs []xml.Attr, local string, out map[T]struct{}) {
	for _, attr := range attrs {
		if attr.Name.Local != local {
			continue
		}

		value := attr.Value
		if idx := strings.IndexByte(value, ':'); idx >= 0 {
			value = value[idx+1:]
		}
		if value != "" {
			out[T(value)] = struct{}{}
		}
	}
}

func sortedNames[T ~string](set map[T]struct{}) []T {
	names := make([]T, 0, len(set))
	for name := range set {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}
