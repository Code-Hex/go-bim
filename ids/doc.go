// Package ids parses buildingSMART Information Delivery Specification (IDS)
// XML documents into plain Go structs.
//
// The parser is intentionally strict. It checks namespace-qualified element
// names, XSD ordering rules, required attributes, requirement cardinality, and
// IDS-specific value rules such as property data types and supported
// xs:restriction facets. Parsed values keep both simpleValue payloads and
// xs:restriction data so callers can inspect the document without losing detail.
//
// Parse, ParseBytes, and ParseFile are the main entry points. By default the
// package infers the IDS version from xsi:schemaLocation, but WithVersion can
// override that when a caller needs to pin the interpretation.
//
// The package supports the official IDS 1.0.0 release. It also recognizes the
// older 0.9.7 candidate schema location because upstream example files still
// reference it.
//
// The vendored reference files under ids/reference/official come from
// buildingSMART and are excluded from the repository's MIT license.
package ids
