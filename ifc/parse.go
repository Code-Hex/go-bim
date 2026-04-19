package ifc

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type parseConfig struct {
	version Version
}

// ParseOption customizes Parse.
type ParseOption func(*parseConfig)

// WithVersion forces the document schema version instead of inferring it from FILE_SCHEMA.
func WithVersion(version Version) ParseOption {
	return func(cfg *parseConfig) {
		cfg.version = version
	}
}

// Parse reads an IFC STEP physical file from r.
func Parse(r io.Reader, opts ...ParseOption) (*Document, error) {
	if r == nil {
		return nil, ErrEmptyInput
	}

	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("ifc: read input: %w", err)
	}
	return ParseBytes(data, opts...)
}

// ParseBytes reads an IFC STEP physical file from data.
func ParseBytes(data []byte, opts ...ParseOption) (*Document, error) {
	if len(bytes.TrimSpace(data)) == 0 {
		return nil, ErrEmptyInput
	}

	cfg := parseConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}

	doc, err := parseSPF(data)
	if err != nil {
		return nil, err
	}

	doc.RawSchemaIdentifiers = doc.Header.SchemaIdentifiers()

	if cfg.version != VersionUnknown {
		doc.Version = cfg.version
		return doc, nil
	}
	for _, schemaID := range doc.RawSchemaIdentifiers {
		version := VersionFromSchemaIdentifier(schemaID)
		if version != VersionUnknown {
			doc.Version = version
			break
		}
	}

	return doc, nil
}

// ParseFile reads an IFC STEP physical file from path.
func ParseFile(path string, opts ...ParseOption) (*Document, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ifc: read %s: %w", path, err)
	}
	return ParseBytes(data, opts...)
}
