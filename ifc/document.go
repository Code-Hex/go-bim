package ifc

import "slices"

// HeaderRecord preserves a header entity such as FILE_NAME.
type HeaderRecord struct {
	Keyword   string
	Arguments []Parameter
}

// Header preserves the SPF HEADER section.
type Header struct {
	Records []HeaderRecord
}

// Record returns the first header record with the given keyword.
func (h *Header) Record(keyword string) (*HeaderRecord, bool) {
	if h == nil {
		return nil, false
	}
	for i := range h.Records {
		if h.Records[i].Keyword == keyword {
			return &h.Records[i], true
		}
	}
	return nil, false
}

// SchemaIdentifiers returns the FILE_SCHEMA values as decoded strings.
func (h *Header) SchemaIdentifiers() []string {
	record, ok := h.Record("FILE_SCHEMA")
	if !ok || len(record.Arguments) == 0 {
		return nil
	}

	values := []string{}
	for _, argument := range record.Arguments {
		values = append(values, collectStringParameters(argument)...)
	}
	return values
}

// Instance preserves a single data section entity instance.
type Instance struct {
	ID        int64
	Entity    string
	Arguments []Parameter
}

// DataSection preserves a DATA section.
type DataSection struct {
	Instances []*Instance
}

// Document preserves a parsed IFC STEP physical file.
type Document struct {
	Header               Header
	Data                 []*DataSection
	Instances            map[int64]*Instance
	Order                []int64
	RawSchemaIdentifiers []string
	Version              Version
}

// Instance returns the referenced entity instance.
func (d *Document) Instance(id int64) (*Instance, bool) {
	if d == nil {
		return nil, false
	}
	instance, ok := d.Instances[id]
	return instance, ok
}

// Schema returns the embedded official schema that matches the document.
func (d *Document) Schema() (*Schema, error) {
	if d == nil {
		return nil, ErrNilDocument
	}
	return SchemaFor(d.Version)
}

// EntityIDs returns the parsed instance identifiers in file order.
func (d *Document) EntityIDs() []int64 {
	if d == nil {
		return nil
	}
	out := slices.Clone(d.Order)
	return out
}

func collectStringParameters(parameter Parameter) []string {
	switch parameter := parameter.(type) {
	case StringParameter:
		return []string{parameter.Decoded}
	case AggregateParameter:
		values := []string{}
		for _, element := range parameter.Elements {
			values = append(values, collectStringParameters(element)...)
		}
		return values
	default:
		return nil
	}
}
