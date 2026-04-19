package ids

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"slices"
	"strings"
)

func (document *Document) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	return unmarshalDocumentStrict(document, decoder, start)
}

func (versions *IFCVersions) UnmarshalXMLAttr(attr xml.Attr) error {
	raw := strings.Fields(attr.Value)
	out := make(IFCVersions, len(raw))
	for i, value := range raw {
		out[i] = IFCVersion(value)
	}
	*versions = out
	return nil
}

func (applicability *Applicability) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*applicability = Applicability{}
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			return fmt.Errorf("%w: unexpected applicability attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
		switch attr.Name.Local {
		case "minOccurs":
			value := attr.Value
			applicability.MinOccurs = &value
		case "maxOccurs":
			value := attr.Value
			applicability.MaxOccurs = &value
		default:
			return fmt.Errorf("%w: unexpected applicability attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
	}
	lastIndex := -1
	seenEntity := false

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				return io.ErrUnexpectedEOF
			}
			return err
		}

		switch token := token.(type) {
		case xml.StartElement:
			index, ok := applicabilityFacetOrder(token.Name)
			if !ok || index < lastIndex {
				return fmt.Errorf("%w: malformed applicability sequence", ErrInvalidDocument)
			}
			if token.Name.Local == string(FacetEntity) {
				if seenEntity {
					return fmt.Errorf("%w: applicability.entity may appear at most once", ErrInvalidDocument)
				}
				seenEntity = true
			}
			lastIndex = index
			facet, err := decodeApplicabilityFacet(decoder, token)
			if err != nil {
				return err
			}
			applicability.Facets = append(applicability.Facets, facet)
		case xml.EndElement:
			if token.Name == start.Name {
				return nil
			}
		case xml.CharData:
			if !isWhitespace(token) {
				return fmt.Errorf("%w: unexpected text in applicability", ErrInvalidDocument)
			}
		}
	}
}

func (requirements *Requirements) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*requirements = Requirements{}
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			return fmt.Errorf("%w: unexpected requirements attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
		switch attr.Name.Local {
		case "description":
			value := attr.Value
			requirements.Description = &value
		default:
			return fmt.Errorf("%w: unexpected requirements attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				return io.ErrUnexpectedEOF
			}
			return err
		}

		switch token := token.(type) {
		case xml.StartElement:
			facet, err := decodeRequirementFacet(decoder, token)
			if err != nil {
				return err
			}
			requirements.Facets = append(requirements.Facets, facet)
		case xml.EndElement:
			if token.Name == start.Name {
				return nil
			}
		case xml.CharData:
			if !isWhitespace(token) {
				return fmt.Errorf("%w: unexpected text in requirements", ErrInvalidDocument)
			}
		}
	}
}

func (value *Value) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*value = Value{}

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				return io.ErrUnexpectedEOF
			}
			return err
		}

		switch token := token.(type) {
		case xml.StartElement:
			switch {
			case isIDSElement(token.Name, "simpleValue"):
				if value.simpleValue != nil || value.restriction != nil {
					return fmt.Errorf("%w: idsValue contains multiple values", ErrInvalidDocument)
				}
				var text string
				if err := decoder.DecodeElement(&text, &token); err != nil {
					return err
				}
				value.kind = ValueKindSimple
				value.simpleValue = &text
			case isXSElement(token.Name, "restriction"):
				if value.simpleValue != nil || value.restriction != nil {
					return fmt.Errorf("%w: idsValue contains multiple values", ErrInvalidDocument)
				}
				var restriction Restriction
				if err := decoder.DecodeElement(&restriction, &token); err != nil {
					return err
				}
				value.kind = ValueKindRestriction
				value.restriction = &restriction
			default:
				return fmt.Errorf("%w: unsupported idsValue child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
		case xml.EndElement:
			if token.Name == start.Name {
				if value.simpleValue == nil && value.restriction == nil {
					return fmt.Errorf("%w: idsValue is empty", ErrInvalidDocument)
				}
				return nil
			}
		case xml.CharData:
			if !isWhitespace(token) {
				return fmt.Errorf("%w: unexpected text in idsValue", ErrInvalidDocument)
			}
		}
	}
}

func (restriction *Restriction) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*restriction = Restriction{}
	seenBase := false
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			return fmt.Errorf("%w: unexpected restriction attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
		switch attr.Name.Local {
		case "base":
			restriction.Base = attr.Value
			restriction.BaseName = resolveQName(attr.Value)
			seenBase = true
		default:
			return fmt.Errorf("%w: unexpected restriction attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
	}
	if !seenBase {
		return fmt.Errorf("%w: restriction.base is required", ErrInvalidDocument)
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				return io.ErrUnexpectedEOF
			}
			return err
		}

		switch token := token.(type) {
		case xml.StartElement:
			if token.Name.Space != xmlSchemaNamespace {
				return fmt.Errorf("%w: unsupported restriction child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
			facet := RestrictionFacet{
				Name:       token.Name,
				Type:       RestrictionFacetType(token.Name.Local),
				Attributes: cloneAttrs(token.Attr),
			}
			seenValue := false
			for _, attr := range token.Attr {
				if attr.Name.Space != "" {
					return fmt.Errorf("%w: unexpected restriction facet attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
				}
				switch attr.Name.Local {
				case "value":
					facet.Value = attr.Value
					seenValue = true
				default:
					return fmt.Errorf("%w: unexpected restriction facet attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
				}
			}
			if !seenValue {
				return fmt.Errorf("%w: restriction facet value is required", ErrInvalidDocument)
			}
			if err := decoder.Skip(); err != nil {
				return err
			}
			restriction.Facets = append(restriction.Facets, facet)
		case xml.EndElement:
			if token.Name == start.Name {
				return nil
			}
		case xml.CharData:
			if !isWhitespace(token) {
				return fmt.Errorf("%w: unexpected text in restriction", ErrInvalidDocument)
			}
		}
	}
}

func decodeApplicabilityFacet(decoder *xml.Decoder, start xml.StartElement) (ApplicabilityFacet, error) {
	if !isIDSElement(start.Name, start.Name.Local) {
		return ApplicabilityFacet{}, fmt.Errorf("%w: unsupported applicability facet %q", ErrInvalidDocument, qualifiedName(start.Name))
	}

	switch start.Name.Local {
	case string(FacetEntity):
		var facet EntityFacet
		if err := decoder.DecodeElement(&facet, &start); err != nil {
			return ApplicabilityFacet{}, err
		}
		return ApplicabilityFacet{kind: FacetEntity, entity: &facet}, nil
	case string(FacetPartOf):
		var facet PartOfFacet
		if err := decoder.DecodeElement(&facet, &start); err != nil {
			return ApplicabilityFacet{}, err
		}
		return ApplicabilityFacet{kind: FacetPartOf, partOf: &facet}, nil
	case string(FacetClassification):
		var facet ClassificationFacet
		if err := decoder.DecodeElement(&facet, &start); err != nil {
			return ApplicabilityFacet{}, err
		}
		return ApplicabilityFacet{kind: FacetClassification, classification: &facet}, nil
	case string(FacetAttribute):
		var facet AttributeFacet
		if err := decoder.DecodeElement(&facet, &start); err != nil {
			return ApplicabilityFacet{}, err
		}
		return ApplicabilityFacet{kind: FacetAttribute, attribute: &facet}, nil
	case string(FacetProperty):
		var facet PropertyFacet
		if err := decoder.DecodeElement(&facet, &start); err != nil {
			return ApplicabilityFacet{}, err
		}
		return ApplicabilityFacet{kind: FacetProperty, property: &facet}, nil
	case string(FacetMaterial):
		var facet MaterialFacet
		if err := decoder.DecodeElement(&facet, &start); err != nil {
			return ApplicabilityFacet{}, err
		}
		return ApplicabilityFacet{kind: FacetMaterial, material: &facet}, nil
	default:
		return ApplicabilityFacet{}, fmt.Errorf("%w: unsupported applicability facet %q", ErrInvalidDocument, qualifiedName(start.Name))
	}
}

func decodeRequirementFacet(decoder *xml.Decoder, start xml.StartElement) (RequirementFacet, error) {
	if !isIDSElement(start.Name, start.Name.Local) {
		return RequirementFacet{}, fmt.Errorf("%w: unsupported requirement facet %q", ErrInvalidDocument, qualifiedName(start.Name))
	}

	switch start.Name.Local {
	case string(FacetEntity):
		var facet RequirementEntityFacet
		if err := decoder.DecodeElement(&facet, &start); err != nil {
			return RequirementFacet{}, err
		}
		return RequirementFacet{kind: FacetEntity, entity: &facet}, nil
	case string(FacetPartOf):
		var facet RequirementPartOfFacet
		if err := decoder.DecodeElement(&facet, &start); err != nil {
			return RequirementFacet{}, err
		}
		return RequirementFacet{kind: FacetPartOf, partOf: &facet}, nil
	case string(FacetClassification):
		var facet RequirementClassificationFacet
		if err := decoder.DecodeElement(&facet, &start); err != nil {
			return RequirementFacet{}, err
		}
		return RequirementFacet{kind: FacetClassification, classification: &facet}, nil
	case string(FacetAttribute):
		var facet RequirementAttributeFacet
		if err := decoder.DecodeElement(&facet, &start); err != nil {
			return RequirementFacet{}, err
		}
		return RequirementFacet{kind: FacetAttribute, attribute: &facet}, nil
	case string(FacetProperty):
		var facet RequirementPropertyFacet
		if err := decoder.DecodeElement(&facet, &start); err != nil {
			return RequirementFacet{}, err
		}
		return RequirementFacet{kind: FacetProperty, property: &facet}, nil
	case string(FacetMaterial):
		var facet RequirementMaterialFacet
		if err := decoder.DecodeElement(&facet, &start); err != nil {
			return RequirementFacet{}, err
		}
		return RequirementFacet{kind: FacetMaterial, material: &facet}, nil
	default:
		return RequirementFacet{}, fmt.Errorf("%w: unsupported requirement facet %q", ErrInvalidDocument, qualifiedName(start.Name))
	}
}

func parseSchemaLocations(attrs []xml.Attr) []SchemaLocation {
	for _, attr := range attrs {
		if attr.Name.Space == xmlSchemaInstanceNamespace && attr.Name.Local == "schemaLocation" {
			fields := strings.Fields(attr.Value)
			locations := make([]SchemaLocation, 0, len(fields)/2)
			for i := 0; i+1 < len(fields); i += 2 {
				locations = append(locations, SchemaLocation{
					Namespace: fields[i],
					Location:  fields[i+1],
				})
			}
			return locations
		}
	}
	return nil
}

func cloneAttrs(attrs []xml.Attr) []xml.Attr {
	return slices.Clone(attrs)
}

func isWhitespace(data xml.CharData) bool {
	return len(bytes.TrimSpace(data)) == 0
}

func isIDSElement(name xml.Name, local string) bool {
	return name.Local == local && name.Space == idsNamespace
}

func isXSElement(name xml.Name, local string) bool {
	return name.Local == local && name.Space == xmlSchemaNamespace
}

func qualifiedName(name xml.Name) string {
	if name.Space == "" {
		return name.Local
	}
	return name.Space + ":" + name.Local
}

func resolveQName(raw string) xml.Name {
	prefix, local, ok := strings.Cut(raw, ":")
	if !ok {
		return xml.Name{Local: raw}
	}

	switch prefix {
	case "xs", "xsd":
		return xml.Name{Space: xmlSchemaNamespace, Local: local}
	default:
		return xml.Name{Local: raw}
	}
}

func applicabilityFacetOrder(name xml.Name) (int, bool) {
	switch {
	case isIDSElement(name, string(FacetEntity)):
		return 0, true
	case isIDSElement(name, string(FacetPartOf)):
		return 1, true
	case isIDSElement(name, string(FacetClassification)):
		return 2, true
	case isIDSElement(name, string(FacetAttribute)):
		return 3, true
	case isIDSElement(name, string(FacetProperty)):
		return 4, true
	case isIDSElement(name, string(FacetMaterial)):
		return 5, true
	default:
		return 0, false
	}
}
