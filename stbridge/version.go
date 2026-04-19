package stbridge

import "strings"

// Version is the normalized ST-Bridge version family for a parsed document or schema.
type Version string

const (
	VersionUnknown  Version = ""
	VersionLegacy1X Version = "1.x"
	Version200      Version = "2.0.0"
	Version201      Version = "2.0.1"
	Version202      Version = "2.0.2"
	Version210      Version = "2.1.0"
)

// KnownVersions reports the supported version families.
func KnownVersions() []Version {
	return []Version{
		VersionLegacy1X,
		Version200,
		Version201,
		Version202,
		Version210,
	}
}

// KnownSchemaVersions reports the versions that have embedded XSD metadata.
func KnownSchemaVersions() []Version {
	return []Version{
		Version200,
		Version201,
		Version202,
		Version210,
	}
}

// HasSchema reports whether the version has embedded XSD metadata.
func (v Version) HasSchema() bool {
	switch v {
	case Version200, Version201, Version202, Version210:
		return true
	default:
		return false
	}
}

func (v Version) String() string {
	return string(v)
}

// NormalizeVersion canonicalizes raw ST-Bridge version strings.
func NormalizeVersion(raw string) Version {
	switch strings.TrimSpace(strings.ToLower(raw)) {
	case "2.0", "2.0.0":
		return Version200
	case "2.0.1":
		return Version201
	case "2.0.2":
		return Version202
	case "2.1", "2.1.0":
		return Version210
	}

	if strings.HasPrefix(strings.TrimSpace(raw), "1.") {
		return VersionLegacy1X
	}

	return VersionUnknown
}
