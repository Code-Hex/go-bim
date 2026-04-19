// Package ifc reads IFC STEP physical files and the official EXPRESS schemas
// used to check them.
//
// The package stays close to the source material. Parsed documents keep header
// records, entity instances, schema identifiers, and raw STEP parameter values
// instead of forcing everything into a higher-level object model up front. That
// makes it useful both as a parser and as a base layer for IFC tooling that
// wants to keep control over interpretation.
//
// The usual entry points are Parse, ParseBytes, and ParseFile. SchemaFor loads
// one of the embedded official EXPRESS schemas, Bind lines up an instance with
// its inherited explicit attributes, and Validate checks a parsed document
// against the embedded schema metadata.
//
// Supported schema releases are:
//
//   - IFC2X3_TC1
//   - IFC4_ADD2_TC1
//   - IFC4X3_ADD2
//
// The vendored EXPRESS files under ifc/reference/express come from
// buildingSMART and are excluded from the repository's MIT license.
package ifc
