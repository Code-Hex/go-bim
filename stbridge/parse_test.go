package stbridge

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"unicode/utf8"
)

func TestParseBytesVersionDetection(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		xml            string
		wantVersion    Version
		wantNamespace  string
		wantCalData    bool
		wantExportInfo bool
	}{
		{
			name: "2.0.0",
			xml: `<?xml version="1.0" encoding="UTF-8"?>
<ST_BRIDGE xmlns="https://www.building-smart.or.jp/dl" version="2.0">
  <StbCommon project_name="demo" app_name="authoring-app" />
  <StbModel />
  <StbExtensions />
</ST_BRIDGE>`,
			wantVersion:   Version200,
			wantNamespace: "https://www.building-smart.or.jp/dl",
		},
		{
			name: "2.0.1",
			xml: `<?xml version="1.0" encoding="UTF-8"?>
<ST_BRIDGE xmlns="https://www.building-smart.or.jp/dl" version="2.0.1">
  <StbCommon project_name="demo" app_name="authoring-app" />
  <StbModel />
  <StbExtensions />
</ST_BRIDGE>`,
			wantVersion:   Version201,
			wantNamespace: "https://www.building-smart.or.jp/dl",
		},
		{
			name: "2.0.2",
			xml: `<?xml version="1.0" encoding="UTF-8"?>
<ST_BRIDGE xmlns="https://www.building-smart.or.jp/dl" version="2.0.2">
  <StbCommon project_name="demo" app_name="authoring-app" app_version="1.0.0" />
  <StbModel />
  <StbExtensions />
  <StbCalData />
  <StbAnaModels />
</ST_BRIDGE>`,
			wantVersion:   Version202,
			wantNamespace: "https://www.building-smart.or.jp/dl",
			wantCalData:   true,
		},
		{
			name: "2.1.0",
			xml: `<?xml version="1.0" encoding="UTF-8"?>
<ST_BRIDGE xmlns="https://www.building-smart.or.jp/dl" version="2.1.0">
  <StbCommon project_name="demo" app_name="authoring-app" app_version="1.0.0" />
  <StbModel />
  <StbExtensions />
  <StbExportInformation />
  <StbCalData />
  <StbAnaModels />
</ST_BRIDGE>`,
			wantVersion:    Version210,
			wantNamespace:  "https://www.building-smart.or.jp/dl",
			wantCalData:    true,
			wantExportInfo: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			doc, err := ParseBytes([]byte(tt.xml))
			if err != nil {
				t.Fatalf("ParseBytes() error = %v", err)
			}

			if doc.Version != tt.wantVersion {
				t.Fatalf("doc.Version = %s, want %s", doc.Version, tt.wantVersion)
			}
			if doc.Namespace != tt.wantNamespace {
				t.Fatalf("doc.Namespace = %q, want %q", doc.Namespace, tt.wantNamespace)
			}
			if doc.Common() == nil {
				t.Fatal("doc.Common() = nil")
			}
			if doc.Model() == nil {
				t.Fatal("doc.Model() = nil")
			}
			if got := doc.CalData() != nil; got != tt.wantCalData {
				t.Fatalf("doc.CalData() presence = %v, want %v", got, tt.wantCalData)
			}
			if got := doc.ExportInformation() != nil; got != tt.wantExportInfo {
				t.Fatalf("doc.ExportInformation() presence = %v, want %v", got, tt.wantExportInfo)
			}
		})
	}
}

func TestParsePreservesPreambleAndWhitespace(t *testing.T) {
	t.Parallel()

	doc, err := ParseBytes([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<!--before-root-->
<ST_BRIDGE xmlns="https://www.building-smart.or.jp/dl" version="2.0.2">
  <StbCommon project_name="demo" app_name="authoring-app" app_version="1.0.0" />
  <StbModel />
</ST_BRIDGE>
<!--after-root-->`))
	if err != nil {
		t.Fatalf("ParseBytes() error = %v", err)
	}

	if !containsNodeType(doc.Leading, NodeTypeComment) {
		t.Fatalf("doc.Leading does not preserve the leading comment: %#v", doc.Leading)
	}
	if !containsNodeType(doc.Trailing, NodeTypeComment) {
		t.Fatalf("doc.Trailing does not preserve the trailing comment: %#v", doc.Trailing)
	}
	if len(doc.Root.Children) == len(doc.Root.ElementChildren()) {
		t.Fatal("root whitespace text nodes were dropped")
	}
}

func TestParseFileLegacyShiftJIS(t *testing.T) {
	t.Parallel()

	path := filepath.Join("testdata", "samples", "20171030", "legacy-rc-20150226.stb")
	doc, err := ParseFile(path)
	if err != nil {
		t.Fatalf("ParseFile(%q) error = %v", path, err)
	}

	if doc.Version != VersionLegacy1X {
		t.Fatalf("doc.Version = %s, want %s", doc.Version, VersionLegacy1X)
	}
	if doc.RawVersion != "1.1.00" {
		t.Fatalf("doc.RawVersion = %q, want %q", doc.RawVersion, "1.1.00")
	}
	if doc.Common() == nil {
		t.Fatal("doc.Common() = nil")
	}
	if doc.Model() == nil {
		t.Fatal("doc.Model() = nil")
	}

	projectName, ok := doc.Common().Attr("project_name")
	if !ok {
		t.Fatal("StbCommon.project_name is missing")
	}
	if projectName == "" {
		t.Fatal("StbCommon.project_name is empty")
	}
	if !utf8.ValidString(projectName) {
		t.Fatalf("project_name is not valid UTF-8: %q", projectName)
	}
	if projectName != "ＲＣ標準モデル_02" {
		t.Fatalf("project_name = %q, want %q", projectName, "ＲＣ標準モデル_02")
	}
}

func TestParseRejectsNonSTBridgeRoot(t *testing.T) {
	t.Parallel()

	_, err := ParseBytes([]byte(`<?xml version="1.0" encoding="UTF-8"?><root/>`))
	if err == nil {
		t.Fatal("ParseBytes() error = nil, want error")
	}
	if !errors.Is(err, ErrInvalidRoot) {
		t.Fatalf("ParseBytes() error = %v, want %v", err, ErrInvalidRoot)
	}
}

func TestParseRejectsMultipleRoots(t *testing.T) {
	t.Parallel()

	_, err := ParseBytes([]byte(`<?xml version="1.0" encoding="UTF-8"?><ST_BRIDGE version="2.0"></ST_BRIDGE><extra/>`))
	if err == nil {
		t.Fatal("ParseBytes() error = nil, want error")
	}
}

func TestParseRejectsInvalidNamespaceFor2X(t *testing.T) {
	t.Parallel()

	_, err := ParseBytes([]byte(`<?xml version="1.0" encoding="UTF-8"?><ST_BRIDGE version="2.0.2"></ST_BRIDGE>`))
	if err == nil {
		t.Fatal("ParseBytes() error = nil, want error")
	}
	if !errors.Is(err, ErrInvalidNamespace) {
		t.Fatalf("ParseBytes() error = %v, want %v", err, ErrInvalidNamespace)
	}
}

func TestDocumentSchemaLegacyUnsupported(t *testing.T) {
	t.Parallel()

	path := filepath.Join("testdata", "samples", "20171030", "legacy-s-20160616.stb")
	doc, err := ParseFile(path)
	if err != nil {
		t.Fatalf("ParseFile(%q) error = %v", path, err)
	}

	_, err = doc.Schema()
	if err == nil {
		t.Fatal("doc.Schema() error = nil, want error")
	}
	if !errors.Is(err, ErrSchemaUnavailable) {
		t.Fatalf("doc.Schema() error = %v, want %v", err, ErrSchemaUnavailable)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func containsNodeType(nodes []*Node, want NodeType) bool {
	for _, node := range nodes {
		if node.Type == want {
			return true
		}
	}
	return false
}
