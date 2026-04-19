package ids

import (
	"encoding/xml"
	"fmt"
	"io"
)

type specificationList []Specification

func unmarshalDocumentStrict(document *Document, decoder *xml.Decoder, start xml.StartElement) error {
	*document = Document{
		XMLName:         start.Name,
		RootAttributes:  cloneAttrs(start.Attr),
		SchemaLocations: parseSchemaLocations(start.Attr),
	}

	seenInfo := false
	seenSpecifications := false

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
			if !isIDSElement(token.Name, token.Name.Local) {
				return fmt.Errorf("%w: unexpected document child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
			switch token.Name.Local {
			case "info":
				if seenInfo || seenSpecifications {
					return fmt.Errorf("%w: info out of order or duplicated", ErrInvalidDocument)
				}
				seenInfo = true
				if err := decoder.DecodeElement(&document.Info, &token); err != nil {
					return err
				}
			case "specifications":
				if !seenInfo || seenSpecifications {
					return fmt.Errorf("%w: specifications out of order or duplicated", ErrInvalidDocument)
				}
				seenSpecifications = true
				var list specificationList
				if err := decoder.DecodeElement(&list, &token); err != nil {
					return err
				}
				document.Specifications = []Specification(list)
			default:
				return fmt.Errorf("%w: unexpected document child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
		case xml.EndElement:
			if token.Name == start.Name {
				if !seenInfo || !seenSpecifications {
					return fmt.Errorf("%w: missing required document sections", ErrInvalidDocument)
				}
				return nil
			}
		case xml.CharData:
			if !isWhitespace(token) {
				return fmt.Errorf("%w: unexpected text in document", ErrInvalidDocument)
			}
		}
	}
}

func (list *specificationList) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*list = nil

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
			if !isIDSElement(token.Name, "specification") {
				return fmt.Errorf("%w: unexpected specifications child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
			var specification Specification
			if err := decoder.DecodeElement(&specification, &token); err != nil {
				return err
			}
			*list = append(*list, specification)
		case xml.EndElement:
			if token.Name == start.Name {
				if len(*list) == 0 {
					return fmt.Errorf("%w: specifications is empty", ErrInvalidDocument)
				}
				return nil
			}
		case xml.CharData:
			if !isWhitespace(token) {
				return fmt.Errorf("%w: unexpected text in specifications", ErrInvalidDocument)
			}
		}
	}
}

func (info *Info) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*info = Info{}

	seen := map[string]bool{}
	lastIndex := -1

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
			if !isIDSElement(token.Name, token.Name.Local) {
				return fmt.Errorf("%w: unexpected info child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
			index, ok := infoElementOrder(token.Name.Local)
			if !ok || index < lastIndex || seen[token.Name.Local] {
				return fmt.Errorf("%w: malformed info sequence", ErrInvalidDocument)
			}
			lastIndex = index
			seen[token.Name.Local] = true

			var text string
			if err := decoder.DecodeElement(&text, &token); err != nil {
				return err
			}
			switch token.Name.Local {
			case "title":
				info.Title = text
			case "copyright":
				info.Copyright = &text
			case "version":
				info.Version = &text
			case "description":
				info.Description = &text
			case "author":
				info.Author = &text
			case "date":
				info.Date = &text
			case "purpose":
				info.Purpose = &text
			case "milestone":
				info.Milestone = &text
			}
		case xml.EndElement:
			if token.Name == start.Name {
				if !seen["title"] {
					return fmt.Errorf("%w: info.title is required", ErrInvalidDocument)
				}
				return nil
			}
		case xml.CharData:
			if !isWhitespace(token) {
				return fmt.Errorf("%w: unexpected text in info", ErrInvalidDocument)
			}
		}
	}
}

func (specification *Specification) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*specification = Specification{}

	seenName := false
	seenIFCVersion := false
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			return fmt.Errorf("%w: unexpected specification attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
		switch attr.Name.Local {
		case "name":
			specification.Name = attr.Value
			seenName = true
		case "ifcVersion":
			if err := specification.IFCVersions.UnmarshalXMLAttr(attr); err != nil {
				return err
			}
			seenIFCVersion = true
		case "identifier":
			value := attr.Value
			specification.Identifier = &value
		case "description":
			value := attr.Value
			specification.Description = &value
		case "instructions":
			value := attr.Value
			specification.Instructions = &value
		default:
			return fmt.Errorf("%w: unexpected specification attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
	}
	if !seenName || !seenIFCVersion {
		return fmt.Errorf("%w: missing required specification attributes", ErrInvalidDocument)
	}

	seenApplicability := false
	seenRequirements := false
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
			if !isIDSElement(token.Name, token.Name.Local) {
				return fmt.Errorf("%w: unexpected specification child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
			switch token.Name.Local {
			case "applicability":
				if seenApplicability || seenRequirements {
					return fmt.Errorf("%w: applicability out of order or duplicated", ErrInvalidDocument)
				}
				seenApplicability = true
				if err := decoder.DecodeElement(&specification.Applicability, &token); err != nil {
					return err
				}
			case "requirements":
				if !seenApplicability || seenRequirements {
					return fmt.Errorf("%w: requirements out of order or duplicated", ErrInvalidDocument)
				}
				seenRequirements = true
				specification.Requirements = &Requirements{}
				if err := decoder.DecodeElement(specification.Requirements, &token); err != nil {
					return err
				}
			default:
				return fmt.Errorf("%w: unexpected specification child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
		case xml.EndElement:
			if token.Name == start.Name {
				if !seenApplicability {
					return fmt.Errorf("%w: specification.applicability is required", ErrInvalidDocument)
				}
				return nil
			}
		case xml.CharData:
			if !isWhitespace(token) {
				return fmt.Errorf("%w: unexpected text in specification", ErrInvalidDocument)
			}
		}
	}
}

func (facet *EntityFacet) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var out struct {
		name           Value
		predefinedType *Value
	}
	if err := unmarshalEntityLike(&out.name, &out.predefinedType, decoder, start); err != nil {
		return err
	}
	*facet = EntityFacet{Name: out.name, PredefinedType: out.predefinedType}
	return nil
}

func (facet *RequirementEntityFacet) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*facet = RequirementEntityFacet{}
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			return fmt.Errorf("%w: unexpected entity requirement attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
		switch attr.Name.Local {
		case "instructions":
			value := attr.Value
			facet.Instructions = &value
		default:
			return fmt.Errorf("%w: unexpected entity requirement attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
	}
	return unmarshalEntityLike(&facet.Name, &facet.PredefinedType, decoder, start)
}

func (facet *PartOfFacet) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*facet = PartOfFacet{}
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			return fmt.Errorf("%w: unexpected partOf attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
		switch attr.Name.Local {
		case "relation":
			value := attr.Value
			facet.Relation = &value
		default:
			return fmt.Errorf("%w: unexpected partOf attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
	}
	return unmarshalPartOfEntity(&facet.Entity, decoder, start)
}

func (facet *RequirementPartOfFacet) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*facet = RequirementPartOfFacet{}
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			return fmt.Errorf("%w: unexpected partOf requirement attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
		switch attr.Name.Local {
		case "relation":
			value := attr.Value
			facet.Relation = &value
		case "cardinality":
			value := Cardinality(attr.Value)
			facet.Cardinality = &value
		case "instructions":
			value := attr.Value
			facet.Instructions = &value
		default:
			return fmt.Errorf("%w: unexpected partOf requirement attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
	}
	return unmarshalPartOfEntity(&facet.Entity, decoder, start)
}

func (facet *ClassificationFacet) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*facet = ClassificationFacet{}
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			return fmt.Errorf("%w: unexpected classification attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
		switch attr.Name.Local {
		case "uri":
			value := attr.Value
			facet.URI = &value
		default:
			return fmt.Errorf("%w: unexpected classification attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
	}
	return unmarshalClassificationLike(&facet.Value, &facet.System, decoder, start)
}

func (facet *RequirementClassificationFacet) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*facet = RequirementClassificationFacet{}
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			return fmt.Errorf("%w: unexpected classification requirement attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
		switch attr.Name.Local {
		case "uri":
			value := attr.Value
			facet.URI = &value
		case "cardinality":
			value := Cardinality(attr.Value)
			facet.Cardinality = &value
		case "instructions":
			value := attr.Value
			facet.Instructions = &value
		default:
			return fmt.Errorf("%w: unexpected classification requirement attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
	}
	return unmarshalClassificationLike(&facet.Value, &facet.System, decoder, start)
}

func (facet *AttributeFacet) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*facet = AttributeFacet{}
	return unmarshalAttributeLike(&facet.Name, &facet.Value, decoder, start)
}

func (facet *RequirementAttributeFacet) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*facet = RequirementAttributeFacet{}
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			return fmt.Errorf("%w: unexpected attribute requirement attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
		switch attr.Name.Local {
		case "cardinality":
			value := Cardinality(attr.Value)
			facet.Cardinality = &value
		case "instructions":
			value := attr.Value
			facet.Instructions = &value
		default:
			return fmt.Errorf("%w: unexpected attribute requirement attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
	}
	return unmarshalAttributeLike(&facet.Name, &facet.Value, decoder, start)
}

func (facet *PropertyFacet) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*facet = PropertyFacet{}
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			return fmt.Errorf("%w: unexpected property attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
		switch attr.Name.Local {
		case "dataType":
			value := attr.Value
			facet.DataType = &value
		case "uri":
			value := attr.Value
			facet.URI = &value
		default:
			return fmt.Errorf("%w: unexpected property attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
	}
	return unmarshalPropertyLike(&facet.PropertySet, &facet.BaseName, &facet.Value, decoder, start)
}

func (facet *RequirementPropertyFacet) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*facet = RequirementPropertyFacet{}
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			return fmt.Errorf("%w: unexpected property requirement attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
		switch attr.Name.Local {
		case "dataType":
			value := attr.Value
			facet.DataType = &value
		case "uri":
			value := attr.Value
			facet.URI = &value
		case "cardinality":
			value := Cardinality(attr.Value)
			facet.Cardinality = &value
		case "instructions":
			value := attr.Value
			facet.Instructions = &value
		default:
			return fmt.Errorf("%w: unexpected property requirement attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
	}
	return unmarshalPropertyLike(&facet.PropertySet, &facet.BaseName, &facet.Value, decoder, start)
}

func (facet *MaterialFacet) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*facet = MaterialFacet{}
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			return fmt.Errorf("%w: unexpected material attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
		switch attr.Name.Local {
		case "uri":
			value := attr.Value
			facet.URI = &value
		default:
			return fmt.Errorf("%w: unexpected material attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
	}
	return unmarshalMaterialLike(&facet.Value, decoder, start)
}

func (facet *RequirementMaterialFacet) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	*facet = RequirementMaterialFacet{}
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			return fmt.Errorf("%w: unexpected material requirement attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
		switch attr.Name.Local {
		case "uri":
			value := attr.Value
			facet.URI = &value
		case "cardinality":
			value := Cardinality(attr.Value)
			facet.Cardinality = &value
		case "instructions":
			value := attr.Value
			facet.Instructions = &value
		default:
			return fmt.Errorf("%w: unexpected material requirement attribute %q", ErrInvalidDocument, qualifiedName(attr.Name))
		}
	}
	return unmarshalMaterialLike(&facet.Value, decoder, start)
}

func unmarshalEntityLike(name *Value, predefinedType **Value, decoder *xml.Decoder, start xml.StartElement) error {
	seenName := false
	seenPredefinedType := false

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
			if !isIDSElement(token.Name, token.Name.Local) {
				return fmt.Errorf("%w: unexpected entity child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
			switch token.Name.Local {
			case "name":
				if seenName || seenPredefinedType {
					return fmt.Errorf("%w: malformed entity sequence", ErrInvalidDocument)
				}
				seenName = true
				if err := decoder.DecodeElement(name, &token); err != nil {
					return err
				}
			case "predefinedType":
				if !seenName || seenPredefinedType {
					return fmt.Errorf("%w: malformed entity sequence", ErrInvalidDocument)
				}
				seenPredefinedType = true
				value := &Value{}
				if err := decoder.DecodeElement(value, &token); err != nil {
					return err
				}
				*predefinedType = value
			default:
				return fmt.Errorf("%w: unexpected entity child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
		case xml.EndElement:
			if token.Name == start.Name {
				if !seenName {
					return fmt.Errorf("%w: entity.name is required", ErrInvalidDocument)
				}
				return nil
			}
		case xml.CharData:
			if !isWhitespace(token) {
				return fmt.Errorf("%w: unexpected text in entity", ErrInvalidDocument)
			}
		}
	}
}

func unmarshalPartOfEntity(entity *EntityFacet, decoder *xml.Decoder, start xml.StartElement) error {
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
			if !isIDSElement(token.Name, "entity") || seenEntity {
				return fmt.Errorf("%w: malformed partOf sequence", ErrInvalidDocument)
			}
			seenEntity = true
			if err := decoder.DecodeElement(entity, &token); err != nil {
				return err
			}
		case xml.EndElement:
			if token.Name == start.Name {
				if !seenEntity {
					return fmt.Errorf("%w: partOf.entity is required", ErrInvalidDocument)
				}
				return nil
			}
		case xml.CharData:
			if !isWhitespace(token) {
				return fmt.Errorf("%w: unexpected text in partOf", ErrInvalidDocument)
			}
		}
	}
}

func unmarshalClassificationLike(value **Value, system *Value, decoder *xml.Decoder, start xml.StartElement) error {
	seenSystem := false
	seenValue := false

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
			if !isIDSElement(token.Name, token.Name.Local) {
				return fmt.Errorf("%w: unexpected classification child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
			switch token.Name.Local {
			case "value":
				if seenValue || seenSystem {
					return fmt.Errorf("%w: malformed classification sequence", ErrInvalidDocument)
				}
				seenValue = true
				child := &Value{}
				if err := decoder.DecodeElement(child, &token); err != nil {
					return err
				}
				*value = child
			case "system":
				if seenSystem {
					return fmt.Errorf("%w: malformed classification sequence", ErrInvalidDocument)
				}
				seenSystem = true
				if err := decoder.DecodeElement(system, &token); err != nil {
					return err
				}
			default:
				return fmt.Errorf("%w: unexpected classification child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
		case xml.EndElement:
			if token.Name == start.Name {
				if !seenSystem {
					return fmt.Errorf("%w: classification.system is required", ErrInvalidDocument)
				}
				return nil
			}
		case xml.CharData:
			if !isWhitespace(token) {
				return fmt.Errorf("%w: unexpected text in classification", ErrInvalidDocument)
			}
		}
	}
}

func unmarshalAttributeLike(name *Value, value **Value, decoder *xml.Decoder, start xml.StartElement) error {
	seenName := false
	seenValue := false

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
			if !isIDSElement(token.Name, token.Name.Local) {
				return fmt.Errorf("%w: unexpected attribute child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
			switch token.Name.Local {
			case "name":
				if seenName || seenValue {
					return fmt.Errorf("%w: malformed attribute sequence", ErrInvalidDocument)
				}
				seenName = true
				if err := decoder.DecodeElement(name, &token); err != nil {
					return err
				}
			case "value":
				if !seenName || seenValue {
					return fmt.Errorf("%w: malformed attribute sequence", ErrInvalidDocument)
				}
				seenValue = true
				child := &Value{}
				if err := decoder.DecodeElement(child, &token); err != nil {
					return err
				}
				*value = child
			default:
				return fmt.Errorf("%w: unexpected attribute child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
		case xml.EndElement:
			if token.Name == start.Name {
				if !seenName {
					return fmt.Errorf("%w: attribute.name is required", ErrInvalidDocument)
				}
				return nil
			}
		case xml.CharData:
			if !isWhitespace(token) {
				return fmt.Errorf("%w: unexpected text in attribute", ErrInvalidDocument)
			}
		}
	}
}

func unmarshalPropertyLike(propertySet *Value, baseName *Value, value **Value, decoder *xml.Decoder, start xml.StartElement) error {
	seenPropertySet := false
	seenBaseName := false
	seenValue := false

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
			if !isIDSElement(token.Name, token.Name.Local) {
				return fmt.Errorf("%w: unexpected property child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
			switch token.Name.Local {
			case "propertySet":
				if seenPropertySet || seenBaseName || seenValue {
					return fmt.Errorf("%w: malformed property sequence", ErrInvalidDocument)
				}
				seenPropertySet = true
				if err := decoder.DecodeElement(propertySet, &token); err != nil {
					return err
				}
			case "baseName":
				if !seenPropertySet || seenBaseName || seenValue {
					return fmt.Errorf("%w: malformed property sequence", ErrInvalidDocument)
				}
				seenBaseName = true
				if err := decoder.DecodeElement(baseName, &token); err != nil {
					return err
				}
			case "value":
				if !seenBaseName || seenValue {
					return fmt.Errorf("%w: malformed property sequence", ErrInvalidDocument)
				}
				seenValue = true
				child := &Value{}
				if err := decoder.DecodeElement(child, &token); err != nil {
					return err
				}
				*value = child
			default:
				return fmt.Errorf("%w: unexpected property child %q", ErrInvalidDocument, qualifiedName(token.Name))
			}
		case xml.EndElement:
			if token.Name == start.Name {
				if !seenPropertySet || !seenBaseName {
					return fmt.Errorf("%w: propertySet and baseName are required", ErrInvalidDocument)
				}
				return nil
			}
		case xml.CharData:
			if !isWhitespace(token) {
				return fmt.Errorf("%w: unexpected text in property", ErrInvalidDocument)
			}
		}
	}
}

func unmarshalMaterialLike(value **Value, decoder *xml.Decoder, start xml.StartElement) error {
	seenValue := false

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
			if !isIDSElement(token.Name, "value") || seenValue {
				return fmt.Errorf("%w: malformed material sequence", ErrInvalidDocument)
			}
			seenValue = true
			child := &Value{}
			if err := decoder.DecodeElement(child, &token); err != nil {
				return err
			}
			*value = child
		case xml.EndElement:
			if token.Name == start.Name {
				return nil
			}
		case xml.CharData:
			if !isWhitespace(token) {
				return fmt.Errorf("%w: unexpected text in material", ErrInvalidDocument)
			}
		}
	}
}

func infoElementOrder(name string) (int, bool) {
	switch name {
	case "title":
		return 0, true
	case "copyright":
		return 1, true
	case "version":
		return 2, true
	case "description":
		return 3, true
	case "author":
		return 4, true
	case "date":
		return 5, true
	case "purpose":
		return 6, true
	case "milestone":
		return 7, true
	default:
		return 0, false
	}
}
