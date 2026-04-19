package ifc

import (
	"reflect"
	"testing"
)

func TestSchemaForOfficialCounts(t *testing.T) {
	t.Parallel()

	cases := []struct {
		version       Version
		name          string
		typeCount     int
		entityCount   int
		functionCount int
		ruleCount     int
	}{
		{VersionIFC2X3TC1, "IFC2X3", 327, 653, 38, 2},
		{VersionIFC4ADD2TC1, "IFC4", 397, 776, 47, 2},
		{VersionIFC4X3ADD2, "IFC4X3_ADD2", 436, 876, 48, 2},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.version.String(), func(t *testing.T) {
			t.Parallel()

			schema, err := SchemaFor(tc.version)
			if err != nil {
				t.Fatalf("SchemaFor(%s): %v", tc.version, err)
			}

			if schema.Name != tc.name {
				t.Fatalf("schema.Name = %q, want %q", schema.Name, tc.name)
			}
			if got := len(schema.Types); got != tc.typeCount {
				t.Fatalf("len(schema.Types) = %d, want %d", got, tc.typeCount)
			}
			if got := len(schema.Entities); got != tc.entityCount {
				t.Fatalf("len(schema.Entities) = %d, want %d", got, tc.entityCount)
			}
			if got := len(schema.Functions); got != tc.functionCount {
				t.Fatalf("len(schema.Functions) = %d, want %d", got, tc.functionCount)
			}
			if got := len(schema.Rules); got != tc.ruleCount {
				t.Fatalf("len(schema.Rules) = %d, want %d", got, tc.ruleCount)
			}
		})
	}
}

func TestSchemaAllExplicitAttributes(t *testing.T) {
	t.Parallel()

	for _, version := range KnownVersions() {
		version := version
		t.Run(version.String(), func(t *testing.T) {
			t.Parallel()

			schema, err := SchemaFor(version)
			if err != nil {
				t.Fatalf("SchemaFor(%s): %v", version, err)
			}

			attrs, err := schema.AllExplicitAttributes("IfcAxis2Placement2D")
			if err != nil {
				t.Fatalf("AllExplicitAttributes: %v", err)
			}

			names := make([]string, 0, len(attrs))
			for _, attr := range attrs {
				names = append(names, attr.Name)
			}

			want := []string{"Location", "RefDirection"}
			if !reflect.DeepEqual(names, want) {
				t.Fatalf("attribute names = %#v, want %#v", names, want)
			}
		})
	}
}

func TestSchemaSpotChecks(t *testing.T) {
	t.Parallel()

	t.Run("IFC4_ADD2_TC1", func(t *testing.T) {
		t.Parallel()

		schema, err := SchemaFor(VersionIFC4ADD2TC1)
		if err != nil {
			t.Fatal(err)
		}

		typ, ok := schema.Type("IfcPropertySetDefinitionSet")
		if !ok {
			t.Fatal("IfcPropertySetDefinitionSet not found")
		}
		if typ.Definition.Kind != TypeKindAggregate || typ.Definition.Name != "SET" {
			t.Fatalf("IfcPropertySetDefinitionSet kind = %#v", typ.Definition)
		}
		if typ.Definition.Lower != "1" || typ.Definition.Upper != "?" {
			t.Fatalf("aggregate bounds = [%s:%s], want [1:?]", typ.Definition.Lower, typ.Definition.Upper)
		}
		if typ.Definition.Element == nil || typ.Definition.Element.Name != "IfcPropertySetDefinition" {
			t.Fatalf("aggregate element = %#v", typ.Definition.Element)
		}

		selectType, ok := schema.Type("IfcWarpingStiffnessSelect")
		if !ok {
			t.Fatal("IfcWarpingStiffnessSelect not found")
		}
		if selectType.Definition.Kind != TypeKindSelect {
			t.Fatalf("IfcWarpingStiffnessSelect kind = %#v", selectType.Definition)
		}
		if !reflect.DeepEqual(selectType.Definition.Items, []string{"IfcBoolean", "IfcWarpingMomentMeasure"}) {
			t.Fatalf("IfcWarpingStiffnessSelect items = %#v", selectType.Definition.Items)
		}

		grid, ok := schema.Entity("IfcGrid")
		if !ok {
			t.Fatal("IfcGrid not found")
		}
		if len(grid.Attributes) == 0 {
			t.Fatal("IfcGrid.Attributes is empty")
		}
		uAxes := grid.Attributes[0]
		if uAxes.Name != "UAxes" {
			t.Fatalf("IfcGrid first attribute = %q, want UAxes", uAxes.Name)
		}
		if !uAxes.Type.Unique || uAxes.Type.Element == nil || uAxes.Type.Element.Name != "IfcGridAxis" {
			t.Fatalf("IfcGrid.UAxes type = %#v", uAxes.Type)
		}

		entity, ok := schema.Entity("IfcAxis2Placement2D")
		if !ok {
			t.Fatal("IfcAxis2Placement2D not found")
		}
		if len(entity.Derived) != 1 || entity.Derived[0].Name != "P" {
			t.Fatalf("derived attrs = %#v", entity.Derived)
		}

		app, ok := schema.Entity("IfcApplication")
		if !ok {
			t.Fatal("IfcApplication not found")
		}
		if len(app.Unique) != 2 {
			t.Fatalf("len(IfcApplication.Unique) = %d, want 2", len(app.Unique))
		}

		fn, ok := schema.Function("IfcBaseAxis")
		if !ok {
			t.Fatal("IfcBaseAxis not found")
		}
		if len(fn.Parameters) != 4 {
			t.Fatalf("len(IfcBaseAxis.Parameters) = %d, want 4", len(fn.Parameters))
		}
		if fn.ReturnType.Kind != TypeKindAggregate || fn.ReturnType.Name != "LIST" {
			t.Fatalf("IfcBaseAxis return type = %#v", fn.ReturnType)
		}

		listToArray, ok := schema.Function("IfcListToArray")
		if !ok {
			t.Fatal("IfcListToArray not found")
		}
		if listToArray.ReturnType.Kind != TypeKindAggregate || listToArray.ReturnType.Name != "ARRAY" {
			t.Fatalf("IfcListToArray return type = %#v", listToArray.ReturnType)
		}
		if listToArray.ReturnType.Element == nil || listToArray.ReturnType.Element.Kind != TypeKindGeneric || listToArray.ReturnType.Element.GenericParameter != "T" {
			t.Fatalf("IfcListToArray return element = %#v", listToArray.ReturnType.Element)
		}
		if len(listToArray.Parameters) == 0 || listToArray.Parameters[0].Type.Kind != TypeKindAggregate {
			t.Fatalf("IfcListToArray first parameter = %#v", listToArray.Parameters)
		}
		if listToArray.Parameters[0].Type.Element == nil || listToArray.Parameters[0].Type.Element.Kind != TypeKindGeneric || listToArray.Parameters[0].Type.Element.GenericParameter != "T" {
			t.Fatalf("IfcListToArray first parameter element = %#v", listToArray.Parameters[0].Type.Element)
		}

		booleanChoose, ok := schema.Function("IfcBooleanChoose")
		if !ok {
			t.Fatal("IfcBooleanChoose not found")
		}
		if got := booleanChoose.Parameters[1].Type; got.Kind != TypeKindGeneric || got.GenericParameter != "Item" {
			t.Fatalf("IfcBooleanChoose generic parameter = %#v", got)
		}
		if got := booleanChoose.ReturnType; got.Kind != TypeKindGeneric || got.GenericParameter != "Item" {
			t.Fatalf("IfcBooleanChoose return type = %#v", got)
		}

		rule, ok := schema.Rule("IfcSingleProjectInstance")
		if !ok {
			t.Fatal("IfcSingleProjectInstance not found")
		}
		if !reflect.DeepEqual(rule.AppliesTo, []string{"IfcProject"}) {
			t.Fatalf("rule appliesTo = %#v", rule.AppliesTo)
		}
	})

	t.Run("IFC4X3_ADD2", func(t *testing.T) {
		t.Parallel()

		schema, err := SchemaFor(VersionIFC4X3ADD2)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := schema.Entity("IfcBridge"); !ok {
			t.Fatal("IfcBridge not found")
		}
		if _, ok := schema.Type("IfcBridgePartTypeEnum"); !ok {
			t.Fatal("IfcBridgePartTypeEnum not found")
		}
		loadConfig, ok := schema.Entity("IfcStructuralLoadConfiguration")
		if !ok {
			t.Fatal("IfcStructuralLoadConfiguration not found")
		}
		if len(loadConfig.Attributes) < 2 {
			t.Fatalf("IfcStructuralLoadConfiguration attrs = %#v", loadConfig.Attributes)
		}
		locations := loadConfig.Attributes[1]
		if locations.Name != "Locations" || !locations.Optional {
			t.Fatalf("Locations attr = %#v", locations)
		}
		if !locations.Type.Unique || locations.Type.Element == nil || locations.Type.Element.Kind != TypeKindAggregate {
			t.Fatalf("Locations type = %#v", locations.Type)
		}
		if locations.Type.Element.Element == nil || locations.Type.Element.Element.Name != "IfcLengthMeasure" {
			t.Fatalf("Locations nested element = %#v", locations.Type.Element)
		}
	})

	t.Run("IFC2X3_TC1", func(t *testing.T) {
		t.Parallel()

		schema, err := SchemaFor(VersionIFC2X3TC1)
		if err != nil {
			t.Fatal(err)
		}
		typ, ok := schema.Type("IfcGloballyUniqueId")
		if !ok {
			t.Fatal("IfcGloballyUniqueId not found")
		}
		if typ.Definition.Kind != TypeKindBuiltin || typ.Definition.Name != "STRING" || !typ.Definition.Fixed {
			t.Fatalf("IfcGloballyUniqueId definition = %#v", typ.Definition)
		}
		if typ.Definition.WidthLower != "22" || typ.Definition.WidthUpper != "22" {
			t.Fatalf("IfcGloballyUniqueId width = [%s:%s], want [22:22]", typ.Definition.WidthLower, typ.Definition.WidthUpper)
		}
	})
}

func TestSchemaForReturnsDetachedCopies(t *testing.T) {
	t.Parallel()

	schema1, err := SchemaFor(VersionIFC4ADD2TC1)
	if err != nil {
		t.Fatal(err)
	}
	delete(schema1.Entities, "IfcWall")

	schema2, err := SchemaFor(VersionIFC4ADD2TC1)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := schema2.Entity("IfcWall"); !ok {
		t.Fatal("SchemaFor returned a shared mutable schema")
	}
}

func TestParseSchemaRejectsInvalidTypeExpr(t *testing.T) {
	t.Parallel()

	_, err := ParseSchemaBytes([]byte(`
SCHEMA BAD;
TYPE Broken = LIST [1:?] IfcLabel;
END_TYPE;
END_SCHEMA;
`))
	if err == nil {
		t.Fatal("ParseSchemaBytes succeeded on invalid type expression")
	}
}
