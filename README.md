# go-bim

go-bim is a small collection of Go packages for building exchange formats.
Right now it covers IFC, IDS, and ST-Bridge.

## packages

### `ifc`

Import `github.com/Code-Hex/go-bim/ifc` when you need to read IFC STEP physical
files, keep the raw parameter structure intact, load embedded EXPRESS schemas,
bind instance arguments to schema attributes, or run schema-backed validation.

Supported IFC releases:

- `IFC2X3_TC1`
- `IFC4_ADD2_TC1`
- `IFC4X3_ADD2`

### `ids`

Import `github.com/Code-Hex/go-bim/ids` when you need a strict IDS parser. It
keeps `simpleValue` and `xs:restriction` data intact, preserves
`xsi:schemaLocation`, and checks namespace rules, ordering, cardinality, and
property `dataType` constraints.

The package supports the official IDS `1.0.0` release and still recognizes the
older `0.9.7` candidate schema location because upstream example files still
use it.

### `stbridge`

Import `github.com/Code-Hex/go-bim/stbridge` when you need to read ST-Bridge
documents, including legacy `1.x` `.stb` files, inspect the raw XML tree, or
work with embedded schema metadata for the published `2.x` XSD files.

The package also ships a typed view layer and generated names so common sections
and attributes are easier to reach from an editor.

Supported ST-Bridge versions:

- legacy `1.x`
- `2.0.0`
- `2.0.1`
- `2.0.2`
- `2.1.0`

## getting started

```bash
go get github.com/Code-Hex/go-bim@latest
```

Parse an IFC file and run schema validation:

```go
package main

import (
	"fmt"
	"log"

	"github.com/Code-Hex/go-bim/ifc"
)

func main() {
	doc, err := ifc.ParseFile("model.ifc")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc.Version)

	report, err := ifc.Validate(doc)
	if err != nil {
		log.Fatal(err)
	}

	if !report.Valid() {
		for _, finding := range report.Findings {
			fmt.Printf("#%d %s: %s\n", finding.InstanceID, finding.Code, finding.Message)
		}
	}
}
```

Parse an IDS file and inspect the declared specifications:

```go
package main

import (
	"fmt"
	"log"

	"github.com/Code-Hex/go-bim/ids"
)

func main() {
	doc, err := ids.ParseFile("requirements.ids")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc.Version)
	for _, spec := range doc.Specifications {
		fmt.Println(spec.Name, spec.IFCVersions.Strings())
	}
}
```

Parse a ST-Bridge file and inspect the common section:

```go
package main

import (
	"fmt"
	"log"

	"github.com/Code-Hex/go-bim/stbridge"
)

func main() {
	doc, err := stbridge.ParseFile("model.stb")
	if err != nil {
		log.Fatal(err)
	}

	projectName, _ := doc.View().Common().ProjectName()
	fmt.Println(doc.Version)
	fmt.Println(projectName)

	schema, err := doc.Schema()
	if err == nil {
		root, _ := schema.Element(stbridge.Element_ST_BRIDGE)
		fmt.Println(len(root.ContentModel.Children))
	}
}
```

## reference material

This repository keeps a few reference sets in-tree because the parsers and tests
need stable fixtures.

- `ifc/reference/express/**` contains copied official IFC EXPRESS schemas
- `ids/reference/official/**` contains copied IDS reference text, XSD, and
  selected example files
- `stbridge/reference/**` contains ST-Bridge XSD files, transcription notes,
  and source notes imported with the `stbridge` package
- `stbridge/testdata/samples/**` contains ST-Bridge sample files used by tests

The fetch scripts live in `scripts/`:

- `scripts/fetch_ifc_references.sh`
- `scripts/fetch_ids_references.sh`

## development

The test surface is split by package.

```bash
go test ./ifc ./ids ./stbridge
```

If you touch vendored reference data, rerun the matching fetch script and keep
the source notes up to date. The `stbridge` package keeps its own provenance
notes under `stbridge/reference/sources/` and `stbridge/reference/SOURCES.txt`.

## license

The Go code in this repository is under MIT.

Some reference material is not. These paths are excluded from the repository
MIT license:

- `ifc/reference/express/**`
- `ids/reference/official/**`
- `stbridge/reference/**`
- `stbridge/testdata/samples/**`

See [LICENSE](./LICENSE) and [THIRD_PARTY_NOTICES.md](./THIRD_PARTY_NOTICES.md).
