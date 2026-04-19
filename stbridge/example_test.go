package stbridge_test

import (
	"fmt"
	"slices"

	"github.com/Code-Hex/go-bim/stbridge"
)

func ExampleParseBytes() {
	doc, err := stbridge.ParseBytes([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<ST_BRIDGE xmlns="https://www.building-smart.or.jp/dl" version="2.0.2">
  <StbCommon project_name="demo" app_name="authoring-app" app_version="1.0.0" />
  <StbModel />
  <StbExtensions />
  <StbCalData />
  <StbAnaModels />
</ST_BRIDGE>`))
	if err != nil {
		panic(err)
	}

	projectName, _ := doc.View().Common().ProjectName()
	fmt.Println(doc.Version)
	fmt.Println(projectName)

	// Output:
	// 2.0.2
	// demo
}

func ExampleDocument_View() {
	doc := mustParseExampleDocument(`<?xml version="1.0" encoding="UTF-8"?>
<ST_BRIDGE xmlns="https://www.building-smart.or.jp/dl" version="2.1.0">
  <StbCommon project_name="demo" app_name="authoring-app" app_version="1.0.0" />
  <StbModel>
    <StbNodes />
  </StbModel>
</ST_BRIDGE>`)

	view := doc.View()
	projectName, _ := view.Common().ProjectName()

	fmt.Println(projectName)
	fmt.Println(view.Model().Nodes() != nil)

	// Output:
	// demo
	// true
}

func ExampleNode_Float64Attr() {
	doc := mustParseExampleDocument(`<?xml version="1.0" encoding="UTF-8"?>
<ST_BRIDGE xmlns="https://www.building-smart.or.jp/dl" version="2.1.0">
  <StbCommon project_name="demo" app_name="authoring-app" app_version="1.0.0" global_offset_X="12.5" />
</ST_BRIDGE>`)

	value, ok, err := doc.Common().Float64Attr(stbridge.Attr_global_offset_X)
	fmt.Println(value, ok, err == nil)

	// Output:
	// 12.5 true true
}

func ExampleSchema_Element() {
	schema, err := stbridge.SchemaFor(stbridge.Version210)
	if err != nil {
		panic(err)
	}

	root, _ := schema.Element(stbridge.Element_ST_BRIDGE)
	common, _ := schema.Element(stbridge.Element_StbCommon)
	projectName, _ := common.Attribute(stbridge.Attr_project_name)

	fmt.Println(root.ChildNames()[0])
	fmt.Println(projectName.Name)

	// Output:
	// StbCommon
	// project_name
}

func ExampleKnownElementNamesFor() {
	fmt.Println(slices.Contains(
		stbridge.KnownElementNamesFor(stbridge.Version202),
		stbridge.Element_StbExportInformation,
	))
	fmt.Println(slices.Contains(
		stbridge.KnownElementNamesFor(stbridge.Version210),
		stbridge.Element_StbExportInformation,
	))

	// Output:
	// false
	// true
}

func mustParseExampleDocument(src string) *stbridge.Document {
	doc, err := stbridge.ParseBytes([]byte(src))
	if err != nil {
		panic(err)
	}
	return doc
}
