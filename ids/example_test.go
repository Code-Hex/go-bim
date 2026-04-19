package ids_test

import (
	"fmt"

	"github.com/Code-Hex/go-bim/ids"
)

func ExampleParseBytes() {
	doc, err := ids.ParseBytes([]byte(`<?xml version="1.0" encoding="utf-8"?>
<ids xmlns="http://standards.buildingsmart.org/IDS"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://standards.buildingsmart.org/IDS http://standards.buildingsmart.org/IDS/1.0/ids.xsd">
  <info>
    <title>Walls</title>
  </info>
  <specifications>
    <specification name="Walls" ifcVersion="IFC4">
      <applicability maxOccurs="unbounded">
        <entity>
          <name><simpleValue>IFCWALL</simpleValue></name>
        </entity>
      </applicability>
      <requirements>
        <attribute>
          <name><simpleValue>Name</simpleValue></name>
          <value><simpleValue>External Wall</simpleValue></value>
        </attribute>
      </requirements>
    </specification>
  </specifications>
</ids>`))
	if err != nil {
		panic(err)
	}

	fmt.Println(doc.Version)
	fmt.Println(doc.Info.Title)
	fmt.Println(doc.Specifications[0].Name)

	// Output:
	// 1.0.0
	// Walls
	// Walls
}
