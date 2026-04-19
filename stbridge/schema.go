package stbridge

import "slices"

// Schema is the extracted metadata for a published ST-Bridge XSD.
//
// Schema values returned from SchemaFor are detached clones, so callers can
// inspect or even modify them without corrupting the shared cache.
type Schema struct {
	Version         Version
	TargetNamespace string
	elements        map[ElementName]*ElementSpec
	simpleTypes     map[SimpleTypeName]TypeSpec
}

// Element returns the named element metadata.
//
// Use the generated Element_* constants when you want completion for the
// schema-backed ST-Bridge element names.
func (s *Schema) Element(name ElementName) (*ElementSpec, bool) {
	if s == nil {
		return nil, false
	}
	spec, ok := s.elements[name]
	return spec, ok
}

// Elements returns a cloned element map keyed by schema element name.
func (s *Schema) Elements() map[ElementName]*ElementSpec {
	if s == nil {
		return nil
	}
	return cloneElementMap(s.elements)
}

// ElementNames returns the published element names in sorted order.
func (s *Schema) ElementNames() []ElementName {
	if s == nil {
		return nil
	}

	names := make([]ElementName, 0, len(s.elements))
	for name := range s.elements {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}

// SimpleType returns the named top-level simple type metadata.
func (s *Schema) SimpleType(name SimpleTypeName) (TypeSpec, bool) {
	if s == nil {
		return TypeSpec{}, false
	}
	spec, ok := s.simpleTypes[name]
	return cloneTypeSpec(spec), ok
}

// SimpleTypes returns a cloned simple-type map keyed by type name.
func (s *Schema) SimpleTypes() map[SimpleTypeName]TypeSpec {
	if s == nil {
		return nil
	}
	return cloneTypeMap(s.simpleTypes)
}

// SimpleTypeNames returns the published simple-type names in sorted order.
func (s *Schema) SimpleTypeNames() []SimpleTypeName {
	if s == nil {
		return nil
	}

	names := make([]SimpleTypeName, 0, len(s.simpleTypes))
	for name := range s.simpleTypes {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}

// ElementSpec describes a schema element.
type ElementSpec struct {
	Name                ElementName
	ValueType           TypeSpec
	Attributes          []AttributeSpec
	IdentityConstraints []IdentityConstraintSpec
	ContentModel        *ParticleSpec
}

// Attribute returns the named attribute metadata.
func (e *ElementSpec) Attribute(name AttributeName) (AttributeSpec, bool) {
	if e == nil {
		return AttributeSpec{}, false
	}
	for _, attribute := range e.Attributes {
		if attribute.Name == name {
			return cloneAttributeSpec(attribute), true
		}
	}
	return AttributeSpec{}, false
}

// AttributeNames returns the attribute names in schema order.
func (e *ElementSpec) AttributeNames() []AttributeName {
	if e == nil {
		return nil
	}

	names := make([]AttributeName, 0, len(e.Attributes))
	for _, attribute := range e.Attributes {
		names = append(names, attribute.Name)
	}
	return names
}

// Child returns the first matching element reference found in the content model.
func (e *ElementSpec) Child(name ElementName) (ChildSpec, bool) {
	if e == nil {
		return ChildSpec{}, false
	}
	return findChildSpec(e.ContentModel, name)
}

// ChildNames returns the child element names in content-model order.
func (e *ElementSpec) ChildNames() []ElementName {
	if e == nil {
		return nil
	}

	names := []ElementName{}
	seen := map[ElementName]struct{}{}
	collectChildNames(e.ContentModel, seen, &names)
	return names
}

// AttributeSpec describes a schema attribute.
type AttributeSpec struct {
	Name     AttributeName
	Type     TypeSpec
	Required bool
	Fixed    string
	Default  string
}

// ChildSpec describes a referenced child element.
type ChildSpec struct {
	Name      ElementName
	MinOccurs int
	MaxOccurs int
}

// IdentityConstraintKind identifies an xs:key/xs:unique/xs:keyref declaration.
type IdentityConstraintKind string

const (
	IdentityConstraintUnique IdentityConstraintKind = "unique"
	IdentityConstraintKey    IdentityConstraintKind = "key"
	IdentityConstraintKeyRef IdentityConstraintKind = "keyref"
)

// IdentityConstraintSpec describes an XSD identity constraint attached to an element.
type IdentityConstraintSpec struct {
	Kind     IdentityConstraintKind
	Name     string
	Selector string
	Fields   []string
	Refer    string
}

// ParticleKind identifies a schema content-model node.
type ParticleKind string

const (
	ParticleElement  ParticleKind = "element"
	ParticleSequence ParticleKind = "sequence"
	ParticleChoice   ParticleKind = "choice"
	ParticleAll      ParticleKind = "all"
)

// ParticleSpec preserves the shape of xs:sequence / xs:choice / xs:all trees.
type ParticleSpec struct {
	Kind      ParticleKind
	Name      ElementName
	MinOccurs int
	MaxOccurs int
	Children  []*ParticleSpec
}

// TypeSpec describes a builtin or simpleType-derived value domain.
type TypeSpec struct {
	Name         SimpleTypeName
	Base         string
	ListItemType string
	Enumerations []string
	Patterns     []string
	MinLength    string
	MaxLength    string
	Length       string
	MinInclusive string
	MaxInclusive string
	MinExclusive string
	MaxExclusive string
}

// IsList reports whether the type is defined as xs:list.
func (t TypeSpec) IsList() bool {
	return t.ListItemType != ""
}

func findChildSpec(particle *ParticleSpec, name ElementName) (ChildSpec, bool) {
	if particle == nil {
		return ChildSpec{}, false
	}
	if particle.Kind == ParticleElement && particle.Name == name {
		return ChildSpec{
			Name:      particle.Name,
			MinOccurs: particle.MinOccurs,
			MaxOccurs: particle.MaxOccurs,
		}, true
	}
	for _, child := range particle.Children {
		spec, ok := findChildSpec(child, name)
		if ok {
			return spec, true
		}
	}
	return ChildSpec{}, false
}

func collectChildNames(particle *ParticleSpec, seen map[ElementName]struct{}, names *[]ElementName) {
	if particle == nil {
		return
	}
	if particle.Kind == ParticleElement {
		if _, ok := seen[particle.Name]; !ok {
			seen[particle.Name] = struct{}{}
			*names = append(*names, particle.Name)
		}
		return
	}
	for _, child := range particle.Children {
		collectChildNames(child, seen, names)
	}
}

func cloneSchema(schema *Schema) *Schema {
	if schema == nil {
		return nil
	}
	return &Schema{
		Version:         schema.Version,
		TargetNamespace: schema.TargetNamespace,
		elements:        cloneElementMap(schema.elements),
		simpleTypes:     cloneTypeMap(schema.simpleTypes),
	}
}

func cloneElementMap(source map[ElementName]*ElementSpec) map[ElementName]*ElementSpec {
	if source == nil {
		return nil
	}
	cloned := make(map[ElementName]*ElementSpec, len(source))
	for name, element := range source {
		cloned[name] = cloneElementSpec(element)
	}
	return cloned
}

func cloneTypeMap(source map[SimpleTypeName]TypeSpec) map[SimpleTypeName]TypeSpec {
	if source == nil {
		return nil
	}
	cloned := make(map[SimpleTypeName]TypeSpec, len(source))
	for name, spec := range source {
		cloned[name] = cloneTypeSpec(spec)
	}
	return cloned
}

func cloneElementSpec(spec *ElementSpec) *ElementSpec {
	if spec == nil {
		return nil
	}
	cloned := &ElementSpec{
		Name:                spec.Name,
		ValueType:           cloneTypeSpec(spec.ValueType),
		Attributes:          make([]AttributeSpec, 0, len(spec.Attributes)),
		IdentityConstraints: append([]IdentityConstraintSpec(nil), spec.IdentityConstraints...),
		ContentModel:        cloneParticleSpec(spec.ContentModel),
	}
	for _, attribute := range spec.Attributes {
		cloned.Attributes = append(cloned.Attributes, cloneAttributeSpec(attribute))
	}
	for i := range cloned.IdentityConstraints {
		cloned.IdentityConstraints[i].Fields = append([]string(nil), cloned.IdentityConstraints[i].Fields...)
	}
	return cloned
}

func cloneAttributeSpec(spec AttributeSpec) AttributeSpec {
	spec.Type = cloneTypeSpec(spec.Type)
	return spec
}

func cloneParticleSpec(spec *ParticleSpec) *ParticleSpec {
	if spec == nil {
		return nil
	}
	cloned := &ParticleSpec{
		Kind:      spec.Kind,
		Name:      spec.Name,
		MinOccurs: spec.MinOccurs,
		MaxOccurs: spec.MaxOccurs,
		Children:  make([]*ParticleSpec, 0, len(spec.Children)),
	}
	for _, child := range spec.Children {
		cloned.Children = append(cloned.Children, cloneParticleSpec(child))
	}
	return cloned
}

func cloneTypeSpec(spec TypeSpec) TypeSpec {
	spec.Enumerations = append([]string(nil), spec.Enumerations...)
	spec.Patterns = append([]string(nil), spec.Patterns...)
	return spec
}
