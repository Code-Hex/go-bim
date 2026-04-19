package ifc

import (
	"strings"
)

// Version identifies an official IFC schema release supported by this package.
type Version string

const (
	VersionUnknown     Version = ""
	VersionIFC2X3TC1   Version = "IFC2X3_TC1"
	VersionIFC4ADD2TC1 Version = "IFC4_ADD2_TC1"
	VersionIFC4X3ADD2  Version = "IFC4X3_ADD2"
)

// KnownVersions reports the supported official IFC schema releases.
func KnownVersions() []Version {
	return []Version{
		VersionIFC2X3TC1,
		VersionIFC4ADD2TC1,
		VersionIFC4X3ADD2,
	}
}

// HasSchema reports whether the version has embedded official EXPRESS source.
func (v Version) HasSchema() bool {
	switch v {
	case VersionIFC2X3TC1, VersionIFC4ADD2TC1, VersionIFC4X3ADD2:
		return true
	default:
		return false
	}
}

// SchemaName reports the EXPRESS schema name used in FILE_SCHEMA and SCHEMA declarations.
func (v Version) SchemaName() string {
	switch v {
	case VersionIFC2X3TC1:
		return "IFC2X3"
	case VersionIFC4ADD2TC1:
		return "IFC4"
	case VersionIFC4X3ADD2:
		return "IFC4X3_ADD2"
	default:
		return ""
	}
}

func (v Version) String() string {
	return string(v)
}

// VersionFromSchemaIdentifier maps FILE_SCHEMA / SCHEMA identifiers to the supported official versions.
func VersionFromSchemaIdentifier(raw string) Version {
	switch normalizeSchemaIdentifier(raw) {
	case "IFC2X3", "IFC2X3TC1", "IFC2X3_TC1":
		return VersionIFC2X3TC1
	case "IFC4", "IFC4ADD2TC1", "IFC4_ADD2_TC1":
		return VersionIFC4ADD2TC1
	case "IFC4X3ADD2", "IFC4X3_ADD2", "IFC4.3ADD2", "IFC4.3_ADD2":
		return VersionIFC4X3ADD2
	default:
		return VersionUnknown
	}
}

func normalizeSchemaIdentifier(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}

	replacer := strings.NewReplacer(" ", "", "-", "", "/", "", ":", "")
	raw = replacer.Replace(raw)
	raw = strings.ToUpper(raw)
	return raw
}
