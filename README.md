# go-bim

go-bim is a Go library for working with buildingSMART IFC and IDS files.
It currently ships two packages: `ifc` for STEP physical files plus embedded
official EXPRESS schemas, and `ids` for strict IDS XML parsing.

## packages

### `ifc`

Use `github.com/Code-Hex/go-bim/ifc` when you need to:

- parse IFC STEP physical files
- keep header records, instances, and raw parameter values intact
- load embedded official EXPRESS schemas for supported releases
- bind instance arguments to inherited entity attributes
- validate parsed documents against schema metadata

Supported IFC releases:

- `IFC2X3_TC1`
- `IFC4_ADD2_TC1`
- `IFC4X3_ADD2`

### `ids`

Use `github.com/Code-Hex/go-bim/ids` when you need to:

- parse IDS XML into plain Go structs
- inspect applicability and requirement facets
- preserve `xsi:schemaLocation` and root attributes
- validate namespace rules, ordering, cardinality, and `dataType` constraints
- inspect `simpleValue` and `xs:restriction` values without flattening them

The parser supports the official IDS `1.0.0` release and still recognizes the
older `0.9.7` candidate schema location because upstream examples still use it.

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

## reference material

This repository vendors official buildingSMART reference files for two reasons.
Tests need stable fixtures, and the parser code needs a local reference when the
network is not around.

- `ifc/reference/express/**` contains copied official IFC EXPRESS schemas
- `ids/reference/official/**` contains copied IDS reference text, XSD, and
  selected example files

The fetch scripts live in `scripts/`:

- `scripts/fetch_ifc_references.sh`
- `scripts/fetch_ids_references.sh`

## development

The test surface is split by package.

```bash
go test ./ifc ./ids
```

If you touch vendored reference data, rerun the matching fetch script and keep
the source notes up to date.

## license

The Go code in this repository is under MIT.

Vendored buildingSMART reference material is not. These paths stay under
buildingSMART's `CC BY-ND 4.0` terms:

- `ifc/reference/express/**`
- `ids/reference/official/**`

See [LICENSE](./LICENSE) and [THIRD_PARTY_NOTICES.md](./THIRD_PARTY_NOTICES.md).
