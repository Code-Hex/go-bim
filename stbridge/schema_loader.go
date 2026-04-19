package stbridge

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

//go:embed reference/xsd/*.xsd
var schemaFiles embed.FS

var (
	schemaByVersion = map[Version]string{
		Version200: "reference/xsd/ST-Bridge_v200.xsd",
		Version201: "reference/xsd/ST-Bridge_v201.xsd",
		Version202: "reference/xsd/ST-Bridge_v202.xsd",
		Version210: "reference/xsd/ST-Bridge_v210.xsd",
	}
	schemaCache sync.Map
)

// SchemaFor loads the embedded schema metadata for version.
func SchemaFor(version Version) (*Schema, error) {
	if !version.HasSchema() {
		return nil, fmt.Errorf("%w: %s", ErrSchemaUnavailable, version)
	}

	if cached, ok := schemaCache.Load(version); ok {
		return cloneSchema(cached.(*Schema)), nil
	}

	filename, ok := schemaByVersion[version]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrSchemaUnavailable, version)
	}

	data, err := schemaFiles.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("stbridge: read embedded schema %s: %w", filename, err)
	}

	parsed, err := parseXML(data)
	if err != nil {
		return nil, fmt.Errorf("stbridge: parse schema %s: %w", filename, err)
	}

	schema, err := buildSchema(version, parsed.root)
	if err != nil {
		return nil, fmt.Errorf("stbridge: build schema %s: %w", filename, err)
	}

	actual, _ := schemaCache.LoadOrStore(version, schema)
	return cloneSchema(actual.(*Schema)), nil
}

type schemaBuilder struct {
	elements        map[string]*Node
	attributeGroups map[string]*Node
	simpleTypes     map[string]*Node
	simpleTypeCache map[string]TypeSpec
}

func buildSchema(version Version, root *Node) (*Schema, error) {
	if root == nil {
		return nil, fmt.Errorf("nil XSD root")
	}
	if root.Name.Local != "schema" {
		return nil, fmt.Errorf("root element %q is not an XSD schema", root.Name.Local)
	}

	builder := &schemaBuilder{
		elements:        map[string]*Node{},
		attributeGroups: map[string]*Node{},
		simpleTypes:     map[string]*Node{},
		simpleTypeCache: map[string]TypeSpec{},
	}

	for _, child := range root.ElementChildren() {
		switch child.Name.Local {
		case "element":
			if name, ok := child.Attr("name"); ok {
				builder.elements[name] = child
			}
		case "attributeGroup":
			if name, ok := child.Attr("name"); ok {
				builder.attributeGroups[name] = child
			}
		case "simpleType":
			if name, ok := child.Attr("name"); ok {
				builder.simpleTypes[name] = child
			}
		}
	}

	schema := &Schema{
		Version:         version,
		TargetNamespace: attrValue(root, "targetNamespace"),
		elements:        make(map[ElementName]*ElementSpec, len(builder.elements)),
		simpleTypes:     make(map[SimpleTypeName]TypeSpec, len(builder.simpleTypes)),
	}

	for name, node := range builder.simpleTypes {
		schema.simpleTypes[name] = builder.buildSimpleType(name, node)
	}

	for name, node := range builder.elements {
		schema.elements[name] = builder.buildElement(name, node)
	}

	return schema, nil
}

func (b *schemaBuilder) buildElement(name string, node *Node) *ElementSpec {
	spec := &ElementSpec{
		Name:       name,
		Attributes: []AttributeSpec{},
	}

	if typeName, ok := node.Attr("type"); ok {
		spec.ValueType = b.resolveNamedType(localName(typeName))
	}

	if complexType := node.Child("complexType"); complexType != nil {
		b.collectComplexType(spec, complexType)
	}

	if simpleType := node.Child("simpleType"); simpleType != nil {
		spec.ValueType = b.buildAnonymousType(simpleType)
	}

	for _, child := range node.ElementChildren() {
		switch child.Name.Local {
		case "unique", "key", "keyref":
			spec.IdentityConstraints = append(spec.IdentityConstraints, identityConstraintFromNode(child))
		}
	}

	return spec
}

func (b *schemaBuilder) collectComplexType(spec *ElementSpec, node *Node) {
	particles := []*ParticleSpec{}

	for _, child := range node.ElementChildren() {
		switch child.Name.Local {
		case "sequence", "choice", "all", "element":
			particle := b.buildParticle(child)
			if particle != nil {
				particles = append(particles, particle)
			}
		case "attribute":
			addAttributeSpec(spec, b.attributeSpecFromNode(child))
		case "attributeGroup":
			b.collectAttributeGroup(spec, child)
		case "simpleContent":
			b.collectSimpleContent(spec, child)
		}
	}

	spec.ContentModel = wrapParticles(particles)
}

func (b *schemaBuilder) collectSimpleContent(spec *ElementSpec, node *Node) {
	for _, child := range node.ElementChildren() {
		if child.Name.Local != "extension" {
			continue
		}

		if base, ok := child.Attr("base"); ok {
			spec.ValueType = b.resolveNamedType(localName(base))
		}

		for _, nested := range child.ElementChildren() {
			switch nested.Name.Local {
			case "attribute":
				addAttributeSpec(spec, b.attributeSpecFromNode(nested))
			case "attributeGroup":
				b.collectAttributeGroup(spec, nested)
			}
		}
	}
}

func (b *schemaBuilder) collectAttributeGroup(spec *ElementSpec, node *Node) {
	target := node
	if ref, ok := node.Attr("ref"); ok {
		group, ok := b.attributeGroups[localName(ref)]
		if !ok {
			return
		}
		target = group
	}

	for _, child := range target.ElementChildren() {
		switch child.Name.Local {
		case "attribute":
			addAttributeSpec(spec, b.attributeSpecFromNode(child))
		case "attributeGroup":
			b.collectAttributeGroup(spec, child)
		}
	}
}

func (b *schemaBuilder) buildParticle(node *Node) *ParticleSpec {
	switch node.Name.Local {
	case "element":
		name, ok := node.Attr("ref")
		if !ok {
			name, ok = node.Attr("name")
		}
		if !ok {
			return nil
		}
		return &ParticleSpec{
			Kind:      ParticleElement,
			Name:      localName(name),
			MinOccurs: parseOccurs(attrValue(node, "minOccurs"), 1),
			MaxOccurs: parseOccurs(attrValue(node, "maxOccurs"), 1),
		}
	case "sequence", "choice", "all":
		children := []*ParticleSpec{}
		for _, child := range node.ElementChildren() {
			particle := b.buildParticle(child)
			if particle != nil {
				children = append(children, particle)
			}
		}
		return &ParticleSpec{
			Kind:      ParticleKind(node.Name.Local),
			MinOccurs: parseOccurs(attrValue(node, "minOccurs"), 1),
			MaxOccurs: parseOccurs(attrValue(node, "maxOccurs"), 1),
			Children:  children,
		}
	default:
		return nil
	}
}

func (b *schemaBuilder) attributeSpecFromNode(node *Node) AttributeSpec {
	spec := AttributeSpec{
		Name:     attrValue(node, "name"),
		Required: strings.EqualFold(attrValue(node, "use"), "required"),
		Fixed:    attrValue(node, "fixed"),
		Default:  attrValue(node, "default"),
	}

	if typeName, ok := node.Attr("type"); ok {
		spec.Type = b.resolveNamedType(localName(typeName))
	}
	if spec.Type.Name == "" {
		if simpleType := node.Child("simpleType"); simpleType != nil {
			spec.Type = b.buildAnonymousType(simpleType)
		}
	}

	return spec
}

func (b *schemaBuilder) buildSimpleType(name string, node *Node) TypeSpec {
	if spec, ok := b.simpleTypeCache[name]; ok {
		return cloneTypeSpec(spec)
	}

	spec := TypeSpec{Name: name}
	b.simpleTypeCache[name] = spec
	spec = b.buildTypeSpec(spec, node)
	b.simpleTypeCache[name] = spec
	return cloneTypeSpec(spec)
}

func (b *schemaBuilder) buildAnonymousType(node *Node) TypeSpec {
	return b.buildTypeSpec(TypeSpec{}, node)
}

func (b *schemaBuilder) buildTypeSpec(initial TypeSpec, node *Node) TypeSpec {
	if restriction := node.Child("restriction"); restriction != nil {
		return b.buildRestrictedType(initial, restriction)
	}
	if list := node.Child("list"); list != nil {
		return b.buildListType(initial, list)
	}
	return initial
}

func (b *schemaBuilder) buildRestrictedType(initial TypeSpec, restriction *Node) TypeSpec {
	spec := cloneTypeSpec(initial)
	baseName := localName(attrValue(restriction, "base"))
	if baseName != "" {
		baseSpec := b.resolveNamedType(baseName)
		if baseSpec.Name != "" {
			spec = mergeTypeSpec(baseSpec, spec.Name)
		}
		if spec.Base == "" {
			spec.Base = baseName
		}
		if spec.Name == "" && initial.Name != "" {
			spec.Name = initial.Name
		}
	}

	for _, child := range restriction.ElementChildren() {
		value := attrValue(child, "value")
		switch child.Name.Local {
		case "enumeration":
			spec.Enumerations = append(spec.Enumerations, value)
		case "pattern":
			spec.Patterns = append(spec.Patterns, value)
		case "minLength":
			spec.MinLength = value
		case "maxLength":
			spec.MaxLength = value
		case "length":
			spec.Length = value
		case "minInclusive":
			spec.MinInclusive = value
		case "maxInclusive":
			spec.MaxInclusive = value
		case "minExclusive":
			spec.MinExclusive = value
		case "maxExclusive":
			spec.MaxExclusive = value
		}
	}

	return spec
}

func (b *schemaBuilder) buildListType(initial TypeSpec, list *Node) TypeSpec {
	spec := cloneTypeSpec(initial)
	spec.Base = "list"

	if itemType, ok := list.Attr("itemType"); ok {
		spec.ListItemType = localName(itemType)
	}

	if nested := list.Child("simpleType"); nested != nil {
		nestedSpec := b.buildAnonymousType(nested)
		spec.ListItemType = nestedSpec.Name
		if spec.ListItemType == "" {
			spec.ListItemType = nestedSpec.Base
		}
	}

	return spec
}

func (b *schemaBuilder) resolveNamedType(name string) TypeSpec {
	if name == "" {
		return TypeSpec{}
	}
	if node, ok := b.simpleTypes[name]; ok {
		return b.buildSimpleType(name, node)
	}
	return TypeSpec{
		Name: name,
		Base: name,
	}
}

func wrapParticles(particles []*ParticleSpec) *ParticleSpec {
	switch len(particles) {
	case 0:
		return nil
	case 1:
		return particles[0]
	default:
		return &ParticleSpec{
			Kind:      ParticleSequence,
			MinOccurs: 1,
			MaxOccurs: 1,
			Children:  particles,
		}
	}
}

func addAttributeSpec(spec *ElementSpec, attribute AttributeSpec) {
	if attribute.Name == "" {
		return
	}
	for _, existing := range spec.Attributes {
		if existing.Name == attribute.Name {
			return
		}
	}
	spec.Attributes = append(spec.Attributes, attribute)
}

func mergeTypeSpec(base TypeSpec, name string) TypeSpec {
	merged := cloneTypeSpec(base)
	if name != "" {
		merged.Name = name
	}
	return merged
}

func identityConstraintFromNode(node *Node) IdentityConstraintSpec {
	spec := IdentityConstraintSpec{
		Kind:  IdentityConstraintKind(node.Name.Local),
		Name:  attrValue(node, "name"),
		Refer: localName(attrValue(node, "refer")),
	}

	if selector := node.Child("selector"); selector != nil {
		spec.Selector = attrValue(selector, "xpath")
	}

	for _, child := range node.ChildrenByName("field") {
		spec.Fields = append(spec.Fields, attrValue(child, "xpath"))
	}

	return spec
}

func localName(name string) string {
	if index := strings.IndexByte(name, ':'); index >= 0 {
		return name[index+1:]
	}
	return name
}

func attrValue(node *Node, name string) string {
	value, _ := node.Attr(name)
	return value
}

func parseOccurs(value string, fallback int) int {
	if value == "" {
		return fallback
	}
	if strings.EqualFold(value, "unbounded") {
		return -1
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}
