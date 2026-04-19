package ids

import "encoding/xml"

const (
	idsNamespace               = "http://standards.buildingsmart.org/IDS"
	xmlSchemaNamespace         = "http://www.w3.org/2001/XMLSchema"
	xmlSchemaInstanceNamespace = "http://www.w3.org/2001/XMLSchema-instance"
)

// Document is a parsed IDS XML document.
//
// Version is the IDS schema version inferred from xsi:schemaLocation or forced
// via WithVersion. It is distinct from Info.Version, which is user-authored
// metadata inside the IDS file.
type Document struct {
	XMLName         xml.Name
	RootAttributes  []xml.Attr
	SchemaLocations []SchemaLocation
	Info            Info
	Specifications  []Specification
	Version         Version
}

// SchemaLocation records an xsi:schemaLocation pair.
type SchemaLocation struct {
	Namespace string
	Location  string
}

// SchemaLocation returns the schema location pair for namespace, if present.
func (d *Document) SchemaLocation(namespace string) (SchemaLocation, bool) {
	if d == nil {
		return SchemaLocation{}, false
	}
	for _, location := range d.SchemaLocations {
		if location.Namespace == namespace {
			return location, true
		}
	}
	return SchemaLocation{}, false
}

// Info contains IDS metadata.
type Info struct {
	Title       string  `xml:"title"`
	Copyright   *string `xml:"copyright"`
	Version     *string `xml:"version"`
	Description *string `xml:"description"`
	Author      *string `xml:"author"`
	Date        *string `xml:"date"`
	Purpose     *string `xml:"purpose"`
	Milestone   *string `xml:"milestone"`
}

// Specification describes one IDS specification entry.
type Specification struct {
	Name          string        `xml:"name,attr"`
	IFCVersions   IFCVersions   `xml:"ifcVersion,attr"`
	Identifier    *string       `xml:"identifier,attr"`
	Description   *string       `xml:"description,attr"`
	Instructions  *string       `xml:"instructions,attr"`
	Applicability Applicability `xml:"applicability"`
	Requirements  *Requirements `xml:"requirements"`
}

// IFCVersion identifies an IFC schema token used in IDS.
type IFCVersion string

const (
	IFCVersionIFC2X3     IFCVersion = "IFC2X3"
	IFCVersionIFC4       IFCVersion = "IFC4"
	IFCVersionIFC4X3     IFCVersion = "IFC4X3"
	IFCVersionIFC4X3ADD2 IFCVersion = "IFC4X3_ADD2"
)

// IFCVersions is the whitespace-separated ifcVersion attribute.
type IFCVersions []IFCVersion

// Strings returns the raw IFC version tokens.
func (versions IFCVersions) Strings() []string {
	out := make([]string, len(versions))
	for i, version := range versions {
		out[i] = string(version)
	}
	return out
}

// Cardinality identifies requirement facet optionality.
type Cardinality string

const (
	CardinalityRequired   Cardinality = "required"
	CardinalityProhibited Cardinality = "prohibited"
	CardinalityOptional   Cardinality = "optional"
)

// FacetKind identifies an IDS facet type.
type FacetKind string

const (
	FacetEntity         FacetKind = "entity"
	FacetPartOf         FacetKind = "partOf"
	FacetClassification FacetKind = "classification"
	FacetAttribute      FacetKind = "attribute"
	FacetProperty       FacetKind = "property"
	FacetMaterial       FacetKind = "material"
)

// Applicability contains applicability facets and minOccurs / maxOccurs metadata.
type Applicability struct {
	MinOccurs *string
	MaxOccurs *string
	Facets    []ApplicabilityFacet
}

// EffectiveMinOccurs reports the effective minOccurs, including the XML Schema default.
func (applicability Applicability) EffectiveMinOccurs() string {
	if applicability.MinOccurs == nil {
		return "1"
	}
	return *applicability.MinOccurs
}

// EffectiveMaxOccurs reports the effective maxOccurs, including the XML Schema default.
func (applicability Applicability) EffectiveMaxOccurs() string {
	if applicability.MaxOccurs == nil {
		return "1"
	}
	return *applicability.MaxOccurs
}

// Requirements contains ordered requirement facets.
type Requirements struct {
	Description *string
	Facets      []RequirementFacet
}

// ApplicabilityFacet is one applicability facet.
type ApplicabilityFacet struct {
	kind           FacetKind
	entity         *EntityFacet
	partOf         *PartOfFacet
	classification *ClassificationFacet
	attribute      *AttributeFacet
	property       *PropertyFacet
	material       *MaterialFacet
}

// Kind reports which facet is populated.
func (facet ApplicabilityFacet) Kind() FacetKind {
	return facet.kind
}

// Entity returns the entity facet, when applicable.
func (facet ApplicabilityFacet) Entity() *EntityFacet {
	return facet.entity
}

// PartOf returns the partOf facet, when applicable.
func (facet ApplicabilityFacet) PartOf() *PartOfFacet {
	return facet.partOf
}

// Classification returns the classification facet, when applicable.
func (facet ApplicabilityFacet) Classification() *ClassificationFacet {
	return facet.classification
}

// Attribute returns the attribute facet, when applicable.
func (facet ApplicabilityFacet) Attribute() *AttributeFacet {
	return facet.attribute
}

// Property returns the property facet, when applicable.
func (facet ApplicabilityFacet) Property() *PropertyFacet {
	return facet.property
}

// Material returns the material facet, when applicable.
func (facet ApplicabilityFacet) Material() *MaterialFacet {
	return facet.material
}

// RequirementFacet is one requirement facet.
type RequirementFacet struct {
	kind           FacetKind
	entity         *RequirementEntityFacet
	partOf         *RequirementPartOfFacet
	classification *RequirementClassificationFacet
	attribute      *RequirementAttributeFacet
	property       *RequirementPropertyFacet
	material       *RequirementMaterialFacet
}

// Kind reports which facet is populated.
func (facet RequirementFacet) Kind() FacetKind {
	return facet.kind
}

// Entity returns the entity facet, when applicable.
func (facet RequirementFacet) Entity() *RequirementEntityFacet {
	return facet.entity
}

// PartOf returns the partOf facet, when applicable.
func (facet RequirementFacet) PartOf() *RequirementPartOfFacet {
	return facet.partOf
}

// Classification returns the classification facet, when applicable.
func (facet RequirementFacet) Classification() *RequirementClassificationFacet {
	return facet.classification
}

// Attribute returns the attribute facet, when applicable.
func (facet RequirementFacet) Attribute() *RequirementAttributeFacet {
	return facet.attribute
}

// Property returns the property facet, when applicable.
func (facet RequirementFacet) Property() *RequirementPropertyFacet {
	return facet.property
}

// Material returns the material facet, when applicable.
func (facet RequirementFacet) Material() *RequirementMaterialFacet {
	return facet.material
}

// EntityFacet is the shared entity facet shape.
type EntityFacet struct {
	Name           Value  `xml:"name"`
	PredefinedType *Value `xml:"predefinedType"`
}

// RequirementEntityFacet is an entity requirement facet.
type RequirementEntityFacet struct {
	Name           Value   `xml:"name"`
	PredefinedType *Value  `xml:"predefinedType"`
	Instructions   *string `xml:"instructions,attr"`
}

// PartOfFacet is the shared partOf facet shape.
type PartOfFacet struct {
	Relation *string     `xml:"relation,attr"`
	Entity   EntityFacet `xml:"entity"`
}

// RequirementPartOfFacet is a partOf requirement facet.
type RequirementPartOfFacet struct {
	Relation     *string      `xml:"relation,attr"`
	Entity       EntityFacet  `xml:"entity"`
	Cardinality  *Cardinality `xml:"cardinality,attr"`
	Instructions *string      `xml:"instructions,attr"`
}

// ClassificationFacet is the shared classification facet shape.
type ClassificationFacet struct {
	Value  *Value  `xml:"value"`
	System Value   `xml:"system"`
	URI    *string `xml:"uri,attr"`
}

// RequirementClassificationFacet is a classification requirement facet.
type RequirementClassificationFacet struct {
	Value        *Value       `xml:"value"`
	System       Value        `xml:"system"`
	URI          *string      `xml:"uri,attr"`
	Cardinality  *Cardinality `xml:"cardinality,attr"`
	Instructions *string      `xml:"instructions,attr"`
}

// AttributeFacet is the shared attribute facet shape.
type AttributeFacet struct {
	Name  Value  `xml:"name"`
	Value *Value `xml:"value"`
}

// RequirementAttributeFacet is an attribute requirement facet.
type RequirementAttributeFacet struct {
	Name         Value        `xml:"name"`
	Value        *Value       `xml:"value"`
	Cardinality  *Cardinality `xml:"cardinality,attr"`
	Instructions *string      `xml:"instructions,attr"`
}

// PropertyFacet is the shared property facet shape.
type PropertyFacet struct {
	PropertySet Value   `xml:"propertySet"`
	BaseName    Value   `xml:"baseName"`
	Value       *Value  `xml:"value"`
	DataType    *string `xml:"dataType,attr"`
	URI         *string `xml:"uri,attr"`
}

// RequirementPropertyFacet is a property requirement facet.
type RequirementPropertyFacet struct {
	PropertySet  Value        `xml:"propertySet"`
	BaseName     Value        `xml:"baseName"`
	Value        *Value       `xml:"value"`
	DataType     *string      `xml:"dataType,attr"`
	URI          *string      `xml:"uri,attr"`
	Cardinality  *Cardinality `xml:"cardinality,attr"`
	Instructions *string      `xml:"instructions,attr"`
}

// MaterialFacet is the shared material facet shape.
type MaterialFacet struct {
	Value *Value  `xml:"value"`
	URI   *string `xml:"uri,attr"`
}

// RequirementMaterialFacet is a material requirement facet.
type RequirementMaterialFacet struct {
	Value        *Value       `xml:"value"`
	URI          *string      `xml:"uri,attr"`
	Cardinality  *Cardinality `xml:"cardinality,attr"`
	Instructions *string      `xml:"instructions,attr"`
}

// Value is an idsValue, which is either a simpleValue or an xs:restriction.
type Value struct {
	kind        ValueKind
	simpleValue *string
	restriction *Restriction
}

// ValueKind identifies which idsValue branch was used.
type ValueKind string

const (
	ValueKindUnknown     ValueKind = ""
	ValueKindSimple      ValueKind = "simpleValue"
	ValueKindRestriction ValueKind = "restriction"
)

// Kind reports whether the value came from simpleValue or xs:restriction.
func (value Value) Kind() ValueKind {
	return value.kind
}

// SimpleValue returns the simpleValue payload, when present.
func (value Value) SimpleValue() (string, bool) {
	if value.kind != ValueKindSimple || value.simpleValue == nil {
		return "", false
	}
	return *value.simpleValue, true
}

// Restriction returns the xs:restriction payload, when present.
func (value Value) Restriction() (*Restriction, bool) {
	if value.kind != ValueKindRestriction || value.restriction == nil {
		return nil, false
	}
	return value.restriction, true
}

// Restriction is an xs:restriction used inside idsValue.
type Restriction struct {
	Base     string
	BaseName xml.Name
	Facets   []RestrictionFacet
}

// RestrictionFacetType identifies an xs:restriction child element.
type RestrictionFacetType string

const (
	RestrictionEnumeration  RestrictionFacetType = "enumeration"
	RestrictionPattern      RestrictionFacetType = "pattern"
	RestrictionMinInclusive RestrictionFacetType = "minInclusive"
	RestrictionMaxInclusive RestrictionFacetType = "maxInclusive"
	RestrictionMinExclusive RestrictionFacetType = "minExclusive"
	RestrictionMaxExclusive RestrictionFacetType = "maxExclusive"
	RestrictionLength       RestrictionFacetType = "length"
	RestrictionMinLength    RestrictionFacetType = "minLength"
	RestrictionMaxLength    RestrictionFacetType = "maxLength"
)

// RestrictionFacet is one xs:restriction child entry.
type RestrictionFacet struct {
	Name       xml.Name
	Type       RestrictionFacetType
	Value      string
	Attributes []xml.Attr
}

// EffectiveCardinality reports the implicit or explicit cardinality.
func (facet RequirementPartOfFacet) EffectiveCardinality() Cardinality {
	if facet.Cardinality == nil {
		return CardinalityRequired
	}
	return *facet.Cardinality
}

// EffectiveCardinality reports the implicit or explicit cardinality.
func (facet RequirementClassificationFacet) EffectiveCardinality() Cardinality {
	if facet.Cardinality == nil {
		return CardinalityRequired
	}
	return *facet.Cardinality
}

// EffectiveCardinality reports the implicit or explicit cardinality.
func (facet RequirementAttributeFacet) EffectiveCardinality() Cardinality {
	if facet.Cardinality == nil {
		return CardinalityRequired
	}
	return *facet.Cardinality
}

// EffectiveCardinality reports the implicit or explicit cardinality.
func (facet RequirementPropertyFacet) EffectiveCardinality() Cardinality {
	if facet.Cardinality == nil {
		return CardinalityRequired
	}
	return *facet.Cardinality
}

// EffectiveCardinality reports the implicit or explicit cardinality.
func (facet RequirementMaterialFacet) EffectiveCardinality() Cardinality {
	if facet.Cardinality == nil {
		return CardinalityRequired
	}
	return *facet.Cardinality
}
