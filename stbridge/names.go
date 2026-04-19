package stbridge

//go:generate go run ./internal/cmd/genstbridgenames

// ElementName is a raw ST-Bridge element name.
//
// You can still pass plain strings around, but the generated Element_*
// constants are much nicer in an editor because completion can do the work.
type ElementName = string

// AttributeName is a raw ST-Bridge attribute name.
//
// The generated Attr_* constants exist for the same reason as Element_*:
// they make it easier to keep typing without cross-checking the spec.
type AttributeName = string

// SimpleTypeName is a named XSD simple type defined by a ST-Bridge schema.
type SimpleTypeName = string

// SectionName is a top-level ST_BRIDGE child element.
type SectionName = string

const (
	// Section_StbCommon is the ST_BRIDGE section that stores project-wide metadata.
	Section_StbCommon SectionName = "StbCommon"
	// Section_StbModel is the ST_BRIDGE section that stores the physical model.
	Section_StbModel SectionName = "StbModel"
	// Section_StbExtensions is the optional section for application-defined extensions.
	Section_StbExtensions SectionName = "StbExtensions"
	// Section_StbExportInformation is the optional export log section added in 2.1.
	Section_StbExportInformation SectionName = "StbExportInformation"
	// Section_StbCalData is the optional calculation-data section added in 2.0.2.
	Section_StbCalData SectionName = "StbCalData"
	// Section_StbAnaModels is the optional analysis-model section added in 2.0.2.
	Section_StbAnaModels SectionName = "StbAnaModels"
)

var knownSectionNames = []SectionName{
	Section_StbCommon,
	Section_StbModel,
	Section_StbExtensions,
	Section_StbExportInformation,
	Section_StbCalData,
	Section_StbAnaModels,
}

// KnownSectionNames returns the union of top-level section names across the
// supported ST-Bridge versions.
//
// When the version boundary matters, use KnownSectionNamesFor instead.
func KnownSectionNames() []SectionName {
	return append([]SectionName(nil), knownSectionNames...)
}

// KnownSectionNamesFor returns the top-level section names defined for version.
//
// Legacy 1.x files do not ship with an embedded XSD in this repository, but
// the parser still exposes the familiar top-level sections that appear in the
// historical samples.
func KnownSectionNamesFor(version Version) []SectionName {
	switch version {
	case VersionLegacy1X, Version200, Version201:
		return []SectionName{
			Section_StbCommon,
			Section_StbModel,
			Section_StbExtensions,
		}
	case Version202:
		return []SectionName{
			Section_StbCommon,
			Section_StbModel,
			Section_StbExtensions,
			Section_StbCalData,
			Section_StbAnaModels,
		}
	case Version210:
		return []SectionName{
			Section_StbCommon,
			Section_StbModel,
			Section_StbExtensions,
			Section_StbExportInformation,
			Section_StbCalData,
			Section_StbAnaModels,
		}
	default:
		return nil
	}
}
