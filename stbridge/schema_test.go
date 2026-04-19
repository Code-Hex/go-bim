package stbridge

import (
	"reflect"
	"testing"
)

func TestSchemaForKnownVersions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		version          Version
		wantElementCount int
		wantSimpleTypes  int
		wantRootChildren []string
		wantHasAppVer    bool
		wantHasCalData   bool
		wantHasExport    bool
	}{
		{
			name:             "2.0.0",
			version:          Version200,
			wantElementCount: 300,
			wantRootChildren: []string{"StbCommon", "StbModel", "StbExtensions"},
		},
		{
			name:             "2.0.1",
			version:          Version201,
			wantElementCount: 301,
			wantRootChildren: []string{"StbCommon", "StbModel", "StbExtensions"},
		},
		{
			name:             "2.0.2",
			version:          Version202,
			wantElementCount: 595,
			wantSimpleTypes:  6,
			wantRootChildren: []string{"StbCommon", "StbModel", "StbExtensions", "StbCalData", "StbAnaModels"},
			wantHasAppVer:    true,
			wantHasCalData:   true,
		},
		{
			name:             "2.1.0",
			version:          Version210,
			wantElementCount: 829,
			wantSimpleTypes:  8,
			wantRootChildren: []string{"StbCommon", "StbModel", "StbExtensions", "StbExportInformation", "StbCalData", "StbAnaModels"},
			wantHasAppVer:    true,
			wantHasCalData:   true,
			wantHasExport:    true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			schema, err := SchemaFor(tt.version)
			if err != nil {
				t.Fatalf("SchemaFor(%s) error = %v", tt.version, err)
			}

			if got := len(schema.Elements()); got != tt.wantElementCount {
				t.Fatalf("len(schema.Elements()) = %d, want %d", got, tt.wantElementCount)
			}
			if tt.wantSimpleTypes > 0 {
				if got := len(schema.SimpleTypes()); got != tt.wantSimpleTypes {
					t.Fatalf("len(schema.SimpleTypes()) = %d, want %d", got, tt.wantSimpleTypes)
				}
			}

			root, ok := schema.Element("ST_BRIDGE")
			if !ok {
				t.Fatal(`schema.Element("ST_BRIDGE") = missing`)
			}
			if root.ContentModel == nil {
				t.Fatal("root.ContentModel = nil")
			}
			gotChildren := make([]string, 0, len(root.ContentModel.Children))
			for _, child := range root.ContentModel.Children {
				gotChildren = append(gotChildren, child.Name)
			}
			if !reflect.DeepEqual(gotChildren, tt.wantRootChildren) {
				t.Fatalf("root children = %#v, want %#v", gotChildren, tt.wantRootChildren)
			}

			common, ok := schema.Element("StbCommon")
			if !ok {
				t.Fatal(`schema.Element("StbCommon") = missing`)
			}
			_, hasAppVersion := common.Attribute("app_version")
			if hasAppVersion != tt.wantHasAppVer {
				t.Fatalf("StbCommon app_version presence = %v, want %v", hasAppVersion, tt.wantHasAppVer)
			}
			_, hasCalData := root.Child("StbCalData")
			if hasCalData != tt.wantHasCalData {
				t.Fatalf("ST_BRIDGE child StbCalData presence = %v, want %v", hasCalData, tt.wantHasCalData)
			}
			_, hasExport := root.Child("StbExportInformation")
			if hasExport != tt.wantHasExport {
				t.Fatalf("ST_BRIDGE child StbExportInformation presence = %v, want %v", hasExport, tt.wantHasExport)
			}
		})
	}
}

func TestSchemaForHandlesPrefixed202XSD(t *testing.T) {
	t.Parallel()

	schema, err := SchemaFor(Version202)
	if err != nil {
		t.Fatalf("SchemaFor(%s) error = %v", Version202, err)
	}

	if schema.TargetNamespace != "https://www.building-smart.or.jp/dl" {
		t.Fatalf("schema.TargetNamespace = %q", schema.TargetNamespace)
	}
}

func TestSchemaForRepresentativeAttributes(t *testing.T) {
	t.Parallel()

	schema, err := SchemaFor(Version210)
	if err != nil {
		t.Fatalf("SchemaFor(%s) error = %v", Version210, err)
	}

	girder, ok := schema.Element("StbGirder")
	if !ok {
		t.Fatal(`schema.Element("StbGirder") = missing`)
	}
	if _, ok := girder.Attribute("id"); !ok {
		t.Fatal(`StbGirder attribute "id" = missing`)
	}
	if _, ok := girder.Attribute("guid"); !ok {
		t.Fatal(`StbGirder attribute "guid" = missing`)
	}
}

func TestSchemaForPreservesIdentityConstraints(t *testing.T) {
	t.Parallel()

	for _, version := range []Version{Version202, Version210} {
		version := version
		t.Run(version.String(), func(t *testing.T) {
			t.Parallel()

			schema, err := SchemaFor(version)
			if err != nil {
				t.Fatalf("SchemaFor(%s) error = %v", version, err)
			}

			root, ok := schema.Element("ST_BRIDGE")
			if !ok {
				t.Fatal(`schema.Element("ST_BRIDGE") = missing`)
			}
			if len(root.IdentityConstraints) == 0 {
				t.Fatal("root.IdentityConstraints is empty")
			}
			if root.IdentityConstraints[0].Kind != IdentityConstraintUnique {
				t.Fatalf("root.IdentityConstraints[0].Kind = %q, want %q", root.IdentityConstraints[0].Kind, IdentityConstraintUnique)
			}
			if root.IdentityConstraints[0].Selector != ".//*" {
				t.Fatalf("root.IdentityConstraints[0].Selector = %q, want %q", root.IdentityConstraints[0].Selector, ".//*")
			}
			if !reflect.DeepEqual(root.IdentityConstraints[0].Fields, []string{"@guid"}) {
				t.Fatalf("root.IdentityConstraints[0].Fields = %#v", root.IdentityConstraints[0].Fields)
			}

			nodes, ok := schema.Element("StbNodes")
			if !ok {
				t.Fatal(`schema.Element("StbNodes") = missing`)
			}
			if len(nodes.IdentityConstraints) == 0 {
				t.Fatal("nodes.IdentityConstraints is empty")
			}
			if nodes.IdentityConstraints[0].Kind != IdentityConstraintKey {
				t.Fatalf("nodes.IdentityConstraints[0].Kind = %q, want %q", nodes.IdentityConstraints[0].Kind, IdentityConstraintKey)
			}
			if !reflect.DeepEqual(nodes.IdentityConstraints[0].Fields, []string{"@id"}) {
				t.Fatalf("nodes.IdentityConstraints[0].Fields = %#v", nodes.IdentityConstraints[0].Fields)
			}
		})
	}
}

func TestSchemaForPreservesChoiceStructure(t *testing.T) {
	t.Parallel()

	tests := []struct {
		version            Version
		wantFigureChildren []string
	}{
		{
			version:            Version202,
			wantFigureChildren: []string{"StbSecColumn_RC_Rect", "StbSecColumn_RC_Circle"},
		},
		{
			version:            Version210,
			wantFigureChildren: []string{"StbSecColumnRect", "StbSecColumnCircle"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.version.String(), func(t *testing.T) {
			t.Parallel()

			schema, err := SchemaFor(tt.version)
			if err != nil {
				t.Fatalf("SchemaFor(%s) error = %v", tt.version, err)
			}

			figure, ok := schema.Element("StbSecFigureColumn_RC")
			if !ok {
				t.Fatal(`schema.Element("StbSecFigureColumn_RC") = missing`)
			}
			if figure.ContentModel == nil {
				t.Fatal("figure.ContentModel = nil")
			}
			if figure.ContentModel.Kind != ParticleChoice {
				t.Fatalf("figure.ContentModel.Kind = %q, want %q", figure.ContentModel.Kind, ParticleChoice)
			}
			gotFigureChildren := []string{
				figure.ContentModel.Children[0].Name,
				figure.ContentModel.Children[1].Name,
			}
			if !reflect.DeepEqual(gotFigureChildren, tt.wantFigureChildren) {
				t.Fatalf("figure choice children = %#v, want %#v", gotFigureChildren, tt.wantFigureChildren)
			}

			arrangement, ok := schema.Element("StbSecBarArrangementColumn_RC")
			if !ok {
				t.Fatal(`schema.Element("StbSecBarArrangementColumn_RC") = missing`)
			}
			if arrangement.ContentModel == nil {
				t.Fatal("arrangement.ContentModel = nil")
			}
			if arrangement.ContentModel.Kind != ParticleChoice {
				t.Fatalf("arrangement.ContentModel.Kind = %q, want %q", arrangement.ContentModel.Kind, ParticleChoice)
			}
			if arrangement.ContentModel.Children[0].Kind != ParticleSequence {
				t.Fatalf("first branch kind = %q, want %q", arrangement.ContentModel.Children[0].Kind, ParticleSequence)
			}
		})
	}
}

func TestSchemaForPreservesTypeFacets(t *testing.T) {
	t.Parallel()

	schema, err := SchemaFor(Version210)
	if err != nil {
		t.Fatalf("SchemaFor(%s) error = %v", Version210, err)
	}

	node, ok := schema.Element("StbNode")
	if !ok {
		t.Fatal(`schema.Element("StbNode") = missing`)
	}
	kind, ok := node.Attribute("kind")
	if !ok {
		t.Fatal(`StbNode attribute "kind" = missing`)
	}
	if got := kind.Type.Enumerations; len(got) != 8 || got[0] != "ON_GIRDER" {
		t.Fatalf("kind.Type.Enumerations = %#v", got)
	}

	guid, ok := schema.SimpleType("guid")
	if !ok {
		t.Fatal(`schema.SimpleType("guid") = missing`)
	}
	if len(guid.Patterns) != 1 || guid.Patterns[0] != "[0-9a-f]{32}" {
		t.Fatalf("guid.Patterns = %#v", guid.Patterns)
	}

	monoLength, ok := schema.SimpleType("monolist_length")
	if !ok {
		t.Fatal(`schema.SimpleType("monolist_length") = missing`)
	}
	if !monoLength.IsList() || monoLength.ListItemType != "length" {
		t.Fatalf("monolist_length = %#v", monoLength)
	}

	monoID, ok := schema.SimpleType("monolist_id")
	if !ok {
		t.Fatal(`schema.SimpleType("monolist_id") = missing`)
	}
	if monoID.ListItemType != "positiveInteger" {
		t.Fatalf("monolist_id.ListItemType = %q, want %q", monoID.ListItemType, "positiveInteger")
	}
	if monoID.MinLength != "3" {
		t.Fatalf("monolist_id.MinLength = %q, want %q", monoID.MinLength, "3")
	}
}

func TestSchemaForReturnsDetachedClones(t *testing.T) {
	t.Parallel()

	first, err := SchemaFor(Version210)
	if err != nil {
		t.Fatalf("SchemaFor(%s) error = %v", Version210, err)
	}
	root, ok := first.Element("ST_BRIDGE")
	if !ok {
		t.Fatal(`first.Element("ST_BRIDGE") = missing`)
	}
	root.Attributes = append(root.Attributes, AttributeSpec{Name: "broken"})

	second, err := SchemaFor(Version210)
	if err != nil {
		t.Fatalf("SchemaFor(%s) error = %v", Version210, err)
	}
	rootAgain, ok := second.Element("ST_BRIDGE")
	if !ok {
		t.Fatal(`second.Element("ST_BRIDGE") = missing`)
	}
	if _, ok := rootAgain.Attribute("broken"); ok {
		t.Fatal(`shared schema cache leaked caller mutation for attribute "broken"`)
	}
}
