// Package stbridge reads ST-Bridge documents and exposes schema metadata for
// the published 2.x specifications.
//
// The parser accepts legacy 1.x files as plain ST-Bridge documents and
// normalizes them to VersionLegacy1X. For the published 2.x specifications,
// SchemaFor returns version-specific metadata extracted from the XSD files
// copied into this repository.
//
// The package keeps the raw XML tree available through Document and Node, and
// it also adds editor-friendly helpers on top:
//
//   - Document.View gives you typed wrappers for the well-known top-level
//     sections and the attributes people usually reach for first.
//   - Generated Element_*, Attr_*, SimpleType_*, and Section_* constants make
//     it easier to discover legal names through completion instead of hunting
//     through the spec.
//   - KnownElementNamesFor, KnownAttributeNamesFor, KnownSimpleTypeNamesFor,
//     and KnownSectionNamesFor let you inspect the published vocabulary for a
//     specific ST-Bridge version.
//
// The typed view is intentionally a union of convenience wrappers across the
// supported versions. Version-specific sections and children simply come back
// empty when they are not available in the parsed document. When the version
// boundary matters up front, start from the Known*For helpers.
//
// The package preserves the parsed XML structure so callers can inspect the
// document directly. It does not try to be a full XML Schema validator.
package stbridge
