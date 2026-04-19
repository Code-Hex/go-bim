package ids

import "strings"

// Version identifies an IDS schema version.
type Version string

const (
	VersionUnknown        Version = ""
	Version0_9_7Candidate Version = "0.9.7"
	Version1_0_0          Version = "1.0.0"
)

// SupportedVersions reports all IDS schema versions this package can identify.
func SupportedVersions() []Version {
	return []Version{
		Version0_9_7Candidate,
		Version1_0_0,
	}
}

// OfficialVersions reports the official buildingSMART IDS standards supported by this package.
func OfficialVersions() []Version {
	return []Version{
		Version1_0_0,
	}
}

// Official reports whether the version is an official buildingSMART standard.
func (v Version) Official() bool {
	return v == Version1_0_0
}

// Supported reports whether the version is known by this package.
func (v Version) Supported() bool {
	switch v {
	case Version0_9_7Candidate, Version1_0_0:
		return true
	default:
		return false
	}
}

func (v Version) String() string {
	return string(v)
}

// VersionFromSchemaLocation maps an IDS schemaLocation URL to a known IDS version.
func VersionFromSchemaLocation(raw string) Version {
	raw = strings.TrimSpace(raw)
	switch {
	case strings.Contains(raw, "/IDS/1.0/ids.xsd"):
		return Version1_0_0
	case strings.Contains(raw, "/IDS/0.9.7/ids.xsd"):
		return Version0_9_7Candidate
	default:
		return VersionUnknown
	}
}

func versionFromSchemaLocations(locations []SchemaLocation) Version {
	for _, location := range locations {
		if version := VersionFromSchemaLocation(location.Location); version != VersionUnknown {
			return version
		}
	}
	return VersionUnknown
}
