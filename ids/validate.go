package ids

import (
	"fmt"
	"net/url"
	"regexp"
	"time"
)

var (
	authorPattern    = regexp.MustCompile(`^[^@]+@[^\.]+\..+$`)
	upperCasePattern = regexp.MustCompile(`^[A-Z]+$`)
	allowedRelations = map[string]struct{}{
		"IFCRELAGGREGATES":                  {},
		"IFCRELASSIGNSTOGROUP":              {},
		"IFCRELCONTAINEDINSPATIALSTRUCTURE": {},
		"IFCRELNESTS":                       {},
		"IFCRELVOIDSELEMENT":                {},
		"IFCRELFILLSELEMENT":                {},
	}
	supportedRestrictionFacetTypes = map[RestrictionFacetType]struct{}{
		RestrictionEnumeration:  {},
		RestrictionPattern:      {},
		RestrictionMinInclusive: {},
		RestrictionMaxInclusive: {},
		RestrictionMinExclusive: {},
		RestrictionMaxExclusive: {},
		RestrictionLength:       {},
		RestrictionMinLength:    {},
		RestrictionMaxLength:    {},
	}
)

func validateDocument(document *Document) error {
	switch document.Version {
	case Version1_0_0:
		return validateDocumentWithRules(document, []IFCVersion{
			IFCVersionIFC2X3,
			IFCVersionIFC4,
			IFCVersionIFC4X3ADD2,
		})
	case Version0_9_7Candidate:
		return validateDocumentWithRules(document, []IFCVersion{
			IFCVersionIFC2X3,
			IFCVersionIFC4,
			IFCVersionIFC4X3,
		})
	default:
		return validateDocumentWithRules(document, []IFCVersion{
			IFCVersionIFC2X3,
			IFCVersionIFC4,
			IFCVersionIFC4X3,
			IFCVersionIFC4X3ADD2,
		})
	}
}

func validateDocumentWithRules(document *Document, allowedIFCVersions []IFCVersion) error {
	allowed := make(map[IFCVersion]struct{}, len(allowedIFCVersions))
	for _, version := range allowedIFCVersions {
		allowed[version] = struct{}{}
	}

	for _, specification := range document.Specifications {
		for _, version := range specification.IFCVersions {
			if _, ok := allowed[version]; !ok {
				return fmt.Errorf("%w: unsupported ifcVersion %q for IDS %s", ErrInvalidDocument, version, document.Version)
			}
		}

		minOccurs := specification.Applicability.EffectiveMinOccurs()
		maxOccurs := specification.Applicability.EffectiveMaxOccurs()
		switch {
		case minOccurs == "1" && maxOccurs == "unbounded":
		case minOccurs == "0" && maxOccurs == "unbounded":
		case minOccurs == "0" && maxOccurs == "0":
		default:
			return fmt.Errorf("%w: unsupported applicability occurs combination [%s:%s]", ErrInvalidDocument, minOccurs, maxOccurs)
		}

		if specification.Requirements == nil {
			goto nextSpec
		}
		for _, facet := range specification.Requirements.Facets {
			switch facet.Kind() {
			case FacetPartOf:
				if !isAllowedCardinality(facet.PartOf().EffectiveCardinality(), CardinalityRequired, CardinalityProhibited) {
					return fmt.Errorf("%w: unsupported partOf requirement cardinality %q", ErrInvalidDocument, facet.PartOf().EffectiveCardinality())
				}
			case FacetClassification:
				if !isAllowedCardinality(facet.Classification().EffectiveCardinality(), CardinalityRequired, CardinalityProhibited, CardinalityOptional) {
					return fmt.Errorf("%w: unsupported classification requirement cardinality %q", ErrInvalidDocument, facet.Classification().EffectiveCardinality())
				}
			case FacetAttribute:
				if !isAllowedCardinality(facet.Attribute().EffectiveCardinality(), CardinalityRequired, CardinalityProhibited, CardinalityOptional) {
					return fmt.Errorf("%w: unsupported attribute requirement cardinality %q", ErrInvalidDocument, facet.Attribute().EffectiveCardinality())
				}
			case FacetProperty:
				if !isAllowedCardinality(facet.Property().EffectiveCardinality(), CardinalityRequired, CardinalityProhibited, CardinalityOptional) {
					return fmt.Errorf("%w: unsupported property requirement cardinality %q", ErrInvalidDocument, facet.Property().EffectiveCardinality())
				}
			case FacetMaterial:
				if !isAllowedCardinality(facet.Material().EffectiveCardinality(), CardinalityRequired, CardinalityProhibited, CardinalityOptional) {
					return fmt.Errorf("%w: unsupported material requirement cardinality %q", ErrInvalidDocument, facet.Material().EffectiveCardinality())
				}
			}
		}
	nextSpec:
		if err := validateSpecificationValues(specification, document.Version); err != nil {
			return err
		}
	}

	if err := validateInfo(document.Info); err != nil {
		return err
	}

	return nil
}

func isAllowedCardinality(cardinality Cardinality, allowed ...Cardinality) bool {
	for _, candidate := range allowed {
		if cardinality == candidate {
			return true
		}
	}
	return false
}

func validateInfo(info Info) error {
	if info.Author != nil && !authorPattern.MatchString(*info.Author) {
		return fmt.Errorf("%w: invalid author %q", ErrInvalidDocument, *info.Author)
	}
	if info.Date != nil && !isValidXSDate(*info.Date) {
		return fmt.Errorf("%w: invalid date %q", ErrInvalidDocument, *info.Date)
	}
	return nil
}

func validateSpecificationValues(specification Specification, version Version) error {
	for _, facet := range specification.Applicability.Facets {
		if err := validateApplicabilityFacet(facet, specification.IFCVersions, version); err != nil {
			return err
		}
	}
	if specification.Requirements == nil {
		return nil
	}
	for _, facet := range specification.Requirements.Facets {
		if err := validateRequirementFacet(facet, specification.IFCVersions); err != nil {
			return err
		}
	}
	return nil
}

func validateApplicabilityFacet(facet ApplicabilityFacet, versions IFCVersions, version Version) error {
	switch facet.Kind() {
	case FacetEntity:
		return validateEntityLikeValues(facet.Entity().Name, facet.Entity().PredefinedType)
	case FacetPartOf:
		if err := validateRelation(facet.PartOf().Relation); err != nil {
			return err
		}
		return validateEntityLikeValues(facet.PartOf().Entity.Name, facet.PartOf().Entity.PredefinedType)
	case FacetClassification:
		if version == Version1_0_0 && facet.Classification().URI != nil {
			return fmt.Errorf("%w: applicability classification uri is not allowed in IDS %s", ErrInvalidDocument, version)
		}
		if err := validateValue(nil, facet.Classification().Value, ""); err != nil {
			return err
		}
		return validateValue(&facet.Classification().System, nil, "")
	case FacetAttribute:
		if err := validateValue(&facet.Attribute().Name, nil, ""); err != nil {
			return err
		}
		return validateValue(nil, facet.Attribute().Value, "")
	case FacetProperty:
		property := facet.Property()
		if version == Version1_0_0 && property.URI != nil {
			return fmt.Errorf("%w: applicability property uri is not allowed in IDS %s", ErrInvalidDocument, version)
		}
		if property.Value != nil && property.DataType == nil {
			return fmt.Errorf("%w: applicability property value requires dataType", ErrInvalidDocument)
		}
		if err := validateValue(&property.PropertySet, nil, ""); err != nil {
			return err
		}
		if err := validateValue(&property.BaseName, nil, ""); err != nil {
			return err
		}
		return validatePropertyDataType(versions, property.DataType, property.Value)
	case FacetMaterial:
		material := facet.Material()
		if version == Version1_0_0 && material.URI != nil {
			return fmt.Errorf("%w: applicability material uri is not allowed in IDS %s", ErrInvalidDocument, version)
		}
		return validateValue(nil, material.Value, "")
	default:
		return nil
	}
}

func validateRequirementFacet(facet RequirementFacet, versions IFCVersions) error {
	switch facet.Kind() {
	case FacetEntity:
		return validateEntityLikeValues(facet.Entity().Name, facet.Entity().PredefinedType)
	case FacetPartOf:
		if err := validateRelation(facet.PartOf().Relation); err != nil {
			return err
		}
		return validateEntityLikeValues(facet.PartOf().Entity.Name, facet.PartOf().Entity.PredefinedType)
	case FacetClassification:
		classification := facet.Classification()
		if classification.EffectiveCardinality() == CardinalityOptional && classification.Value == nil {
			return fmt.Errorf("%w: optional classification requires value", ErrInvalidDocument)
		}
		if err := validateURI(classification.URI); err != nil {
			return err
		}
		if err := validateValue(nil, classification.Value, ""); err != nil {
			return err
		}
		return validateValue(&classification.System, nil, "")
	case FacetAttribute:
		attribute := facet.Attribute()
		if attribute.EffectiveCardinality() == CardinalityOptional && attribute.Value == nil {
			return fmt.Errorf("%w: optional attribute requires value", ErrInvalidDocument)
		}
		if err := validateValue(&attribute.Name, nil, ""); err != nil {
			return err
		}
		return validateValue(nil, attribute.Value, "")
	case FacetProperty:
		property := facet.Property()
		if property.Value != nil && property.DataType == nil {
			return fmt.Errorf("%w: property value requires dataType", ErrInvalidDocument)
		}
		switch property.EffectiveCardinality() {
		case CardinalityOptional:
			if property.DataType == nil {
				return fmt.Errorf("%w: optional property requires dataType", ErrInvalidDocument)
			}
		case CardinalityProhibited:
			if property.DataType != nil || property.Value != nil {
				return fmt.Errorf("%w: prohibited property cannot constrain dataType or value", ErrInvalidDocument)
			}
		}
		if err := validateValue(&property.PropertySet, nil, ""); err != nil {
			return err
		}
		if err := validateValue(&property.BaseName, nil, ""); err != nil {
			return err
		}
		if err := validateURI(property.URI); err != nil {
			return err
		}
		return validatePropertyDataType(versions, property.DataType, property.Value)
	case FacetMaterial:
		material := facet.Material()
		if material.EffectiveCardinality() == CardinalityOptional && material.Value == nil {
			return fmt.Errorf("%w: optional material requires value", ErrInvalidDocument)
		}
		if err := validateURI(material.URI); err != nil {
			return err
		}
		return validateValue(nil, material.Value, "")
	default:
		return nil
	}
}

func validateRelation(relation *string) error {
	if relation == nil {
		return nil
	}
	if _, ok := allowedRelations[*relation]; !ok {
		return fmt.Errorf("%w: unsupported relation %q", ErrInvalidDocument, *relation)
	}
	return nil
}

func validateDataType(versions IFCVersions, dataType *string) (dataTypeSpec, error) {
	if dataType == nil {
		return dataTypeSpec{}, nil
	}
	if !upperCasePattern.MatchString(*dataType) {
		return dataTypeSpec{}, fmt.Errorf("%w: invalid dataType %q", ErrInvalidDocument, *dataType)
	}
	spec, ok, err := lookupDataTypeSpec(*dataType)
	if err != nil {
		return dataTypeSpec{}, err
	}
	if !ok {
		return dataTypeSpec{}, fmt.Errorf("%w: unsupported dataType %q", ErrInvalidDocument, *dataType)
	}
	for _, version := range versions {
		if !spec.supports(version) {
			return dataTypeSpec{}, fmt.Errorf("%w: dataType %q is not supported for %s", ErrInvalidDocument, *dataType, version)
		}
	}
	return spec, nil
}

func validateURI(raw *string) error {
	if raw == nil {
		return nil
	}
	if _, err := url.Parse(*raw); err != nil {
		return fmt.Errorf("%w: invalid uri %q", ErrInvalidDocument, *raw)
	}
	return nil
}

func isValidXSDate(raw string) bool {
	layouts := []string{
		"2006-01-02",
		"2006-01-02Z07:00",
	}
	for _, layout := range layouts {
		if _, err := time.Parse(layout, raw); err == nil {
			return true
		}
	}
	return false
}

func validatePropertyDataType(versions IFCVersions, dataType *string, value *Value) error {
	spec, err := validateDataType(versions, dataType)
	if err != nil {
		return err
	}
	if dataType == nil || value == nil || spec.Base == "" {
		return validateValue(nil, value, "")
	}
	return validateValue(nil, value, spec.Base)
}

func validateEntityLikeValues(name Value, predefinedType *Value) error {
	if err := validateValue(&name, nil, ""); err != nil {
		return err
	}
	return validateValue(nil, predefinedType, "")
}

func validateValue(required *Value, optional *Value, expectedBase string) error {
	value := required
	if value == nil {
		value = optional
	}
	if value == nil {
		return nil
	}

	if simple, ok := value.SimpleValue(); ok {
		return validateLexicalValue("simpleValue", simple, expectedBase)
	}

	restriction, ok := value.Restriction()
	if !ok {
		return nil
	}
	if expectedBase != "" && restriction.Base != expectedBase {
		return fmt.Errorf("%w: restriction base %q does not match %s", ErrInvalidDocument, restriction.Base, expectedBase)
	}
	return validateRestriction(restriction)
}

func validateRestriction(restriction *Restriction) error {
	if restriction == nil {
		return nil
	}
	if _, ok := xmlBasePatterns[restriction.Base]; !ok {
		return fmt.Errorf("%w: unsupported restriction base %q", ErrInvalidDocument, restriction.Base)
	}
	for _, facet := range restriction.Facets {
		if _, ok := supportedRestrictionFacetTypes[facet.Type]; !ok {
			return fmt.Errorf("%w: unsupported restriction facet %q", ErrInvalidDocument, facet.Type)
		}
		switch facet.Type {
		case RestrictionEnumeration, RestrictionMinInclusive, RestrictionMaxInclusive, RestrictionMinExclusive, RestrictionMaxExclusive:
			if err := validateLexicalValue(fmt.Sprintf("restriction facet %s", facet.Type), facet.Value, restriction.Base); err != nil {
				return err
			}
		case RestrictionLength, RestrictionMinLength, RestrictionMaxLength:
			if err := validateLexicalValue(fmt.Sprintf("restriction facet %s", facet.Type), facet.Value, "xs:integer"); err != nil {
				return err
			}
		}
	}
	return nil
}

func validateLexicalValue(label, raw, base string) error {
	if base == "" {
		return nil
	}
	pattern, ok := xmlBasePatterns[base]
	if !ok {
		return fmt.Errorf("%w: unsupported restriction base %q", ErrInvalidDocument, base)
	}
	if !pattern.MatchString(raw) {
		return fmt.Errorf("%w: %s %q does not satisfy %s", ErrInvalidDocument, label, raw, base)
	}
	return nil
}
