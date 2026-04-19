package ifc_test

import (
	"fmt"

	"github.com/Code-Hex/go-bim/ifc"
)

func ExampleParseBytes() {
	doc, err := ifc.ParseBytes([]byte(`
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCORGANIZATION($,'Code-Hex',$,$,$);
ENDSEC;
END-ISO-10303-21;
`))
	if err != nil {
		panic(err)
	}

	fmt.Println(doc.Version)
	fmt.Println(doc.Header.SchemaIdentifiers())
	fmt.Println(doc.EntityIDs())

	// Output:
	// IFC4_ADD2_TC1
	// [IFC4]
	// [1]
}

func ExampleValidate() {
	doc, err := ifc.ParseBytes([]byte(`
ISO-10303-21;
HEADER;
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCORGANIZATION($,'Code-Hex',$,$,$);
ENDSEC;
END-ISO-10303-21;
`))
	if err != nil {
		panic(err)
	}

	report, err := ifc.Validate(doc)
	if err != nil {
		panic(err)
	}

	fmt.Println(report.Valid())
	fmt.Println(len(report.Findings))

	// Output:
	// true
	// 0
}
