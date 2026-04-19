package ifc

import (
	"os"
	"testing"
)

// TestParseExternalIFC is an opt-in integration test for real-world IFC files.
//
// Usage:
//
//	GO_BIM_EXTERNAL_IFC_PATH=/absolute/path/to/model.ifc go test -run TestParseExternalIFC -v ./ifc
func TestParseExternalIFC(t *testing.T) {
	t.Helper()

	path := os.Getenv("GO_BIM_EXTERNAL_IFC_PATH")
	if path == "" {
		t.Skip("set GO_BIM_EXTERNAL_IFC_PATH to run this integration test")
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat %s: %v", path, err)
	}

	doc, err := ParseFile(path)
	if err != nil {
		t.Fatalf("ParseFile(%s): %v", path, err)
	}

	t.Logf("file=%s", path)
	t.Logf("size_bytes=%d", info.Size())
	t.Logf("version=%s", doc.Version)
	t.Logf("raw_schema=%v", doc.RawSchemaIdentifiers)
	t.Logf("data_sections=%d", len(doc.Data))
	t.Logf("instances=%d", len(doc.Instances))

	if len(doc.Instances) == 0 {
		t.Fatal("expected at least one instance")
	}
	if len(doc.RawSchemaIdentifiers) == 0 {
		t.Fatal("expected FILE_SCHEMA to be parsed")
	}

	report, err := Validate(doc)
	if err != nil {
		t.Fatalf("Validate(%s): %v", path, err)
	}
	t.Logf("validation_findings=%d", len(report.Findings))
	if !report.Valid() {
		t.Fatalf("expected validation to pass, findings = %#v", report.Findings)
	}

	counts := map[string]int{}
	for _, id := range doc.Order {
		counts[doc.Instances[id].Entity]++
	}

	for _, entity := range []string{
		"IFCPROJECT",
		"IFCSITE",
		"IFCBUILDING",
		"IFCBUILDINGSTOREY",
		"IFCWALL",
		"IFCDOOR",
		"IFCWINDOW",
		"IFCSPACE",
		"IFCBEAM",
		"IFCSLAB",
		"IFCROOF",
		"IFCSTAIR",
		"IFCRELAGGREGATES",
		"IFCRELCONTAINEDINSPATIALSTRUCTURE",
	} {
		if count := counts[entity]; count > 0 {
			t.Logf("count[%s]=%d", entity, count)
		}
	}

	if counts["IFCPROJECT"] == 0 {
		t.Fatal("expected at least one IFCPROJECT")
	}
	if counts["IFCSITE"] == 0 && counts["IFCBUILDING"] == 0 {
		t.Fatal("expected at least one spatial structure entity")
	}
}
