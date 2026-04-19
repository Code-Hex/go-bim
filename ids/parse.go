package ids

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type parseConfig struct {
	version Version
}

// ParseOption customizes Parse.
type ParseOption interface {
	applyParseOption(*parseConfig)
}

type parseOptionFunc func(*parseConfig)

func (fn parseOptionFunc) applyParseOption(cfg *parseConfig) {
	fn(cfg)
}

// WithVersion forces the IDS schema version instead of inferring it from xsi:schemaLocation.
func WithVersion(version Version) ParseOption {
	return parseOptionFunc(func(cfg *parseConfig) {
		cfg.version = version
	})
}

// Parse reads an IDS XML document from r.
func Parse(r io.Reader, opts ...ParseOption) (*Document, error) {
	if r == nil {
		return nil, ErrEmptyInput
	}

	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("ids: read input: %w", err)
	}
	return ParseBytes(data, opts...)
}

// ParseBytes reads an IDS XML document from data.
func ParseBytes(data []byte, opts ...ParseOption) (*Document, error) {
	if len(bytes.TrimSpace(data)) == 0 {
		return nil, ErrEmptyInput
	}

	cfg := parseConfig{}
	for _, opt := range opts {
		if opt == nil {
			return nil, ErrInvalidOption
		}
		opt.applyParseOption(&cfg)
	}
	if cfg.version != VersionUnknown && !cfg.version.Supported() {
		return nil, fmt.Errorf("%w: %q", ErrUnsupportedVersion, cfg.version)
	}

	var doc Document
	decoder := xml.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&doc); err != nil {
		return nil, fmt.Errorf("ids: decode XML: %w", err)
	}
	if !isIDSElement(doc.XMLName, "ids") {
		return nil, fmt.Errorf("%w: unexpected root element %q", ErrInvalidDocument, qualifiedName(doc.XMLName))
	}

	if cfg.version != VersionUnknown {
		doc.Version = cfg.version
	} else {
		doc.Version = versionFromSchemaLocations(doc.SchemaLocations)
	}
	if err := validateDocument(&doc); err != nil {
		return nil, err
	}

	return &doc, nil
}

// ParseFile reads an IDS XML document from path.
func ParseFile(path string, opts ...ParseOption) (*Document, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ids: read %s: %w", path, err)
	}
	return ParseBytes(data, opts...)
}
