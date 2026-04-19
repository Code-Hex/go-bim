package ifc

import (
	"slices"
	"strings"
)

// Schema preserves the declarations from an IFC EXPRESS schema.
type Schema struct {
	Version   Version
	Name      string
	SourceURL string

	Types     map[string]*TypeDeclaration
	Entities  map[string]*EntityDeclaration
	Functions map[string]*FunctionDeclaration
	Rules     map[string]*RuleDeclaration
}

// Type returns the named TYPE declaration.
func (s *Schema) Type(name string) (*TypeDeclaration, bool) {
	if s == nil {
		return nil, false
	}
	return lookupDeclaration(s.Types, name)
}

// Entity returns the named ENTITY declaration.
func (s *Schema) Entity(name string) (*EntityDeclaration, bool) {
	if s == nil {
		return nil, false
	}
	return lookupDeclaration(s.Entities, name)
}

// Function returns the named FUNCTION declaration.
func (s *Schema) Function(name string) (*FunctionDeclaration, bool) {
	if s == nil {
		return nil, false
	}
	return lookupDeclaration(s.Functions, name)
}

// Rule returns the named RULE declaration.
func (s *Schema) Rule(name string) (*RuleDeclaration, bool) {
	if s == nil {
		return nil, false
	}
	return lookupDeclaration(s.Rules, name)
}

// TypeNames returns the declared TYPE names in sorted order.
func (s *Schema) TypeNames() []string {
	return sortedKeys(s.Types)
}

// EntityNames returns the declared ENTITY names in sorted order.
func (s *Schema) EntityNames() []string {
	return sortedKeys(s.Entities)
}

// FunctionNames returns the declared FUNCTION names in sorted order.
func (s *Schema) FunctionNames() []string {
	return sortedKeys(s.Functions)
}

// RuleNames returns the declared RULE names in sorted order.
func (s *Schema) RuleNames() []string {
	return sortedKeys(s.Rules)
}

// NamedRule preserves a WHERE/UNIQUE rule label and expression.
type NamedRule struct {
	Name       string
	Expression string
}

// TypeDeclaration preserves a TYPE declaration.
type TypeDeclaration struct {
	Name       string
	Definition TypeExpr
	Where      []NamedRule
	Raw        string
}

// EntityDeclaration preserves an ENTITY declaration.
type EntityDeclaration struct {
	Name                string
	Abstract            bool
	SupertypeExpression string
	Supertypes          []string
	Attributes          []ExplicitAttribute
	Derived             []DerivedAttribute
	Inverse             []InverseAttribute
	Unique              []NamedRule
	Where               []NamedRule
	Raw                 string
}

// ExplicitAttribute preserves an explicit entity attribute.
type ExplicitAttribute struct {
	Name     string
	Type     TypeExpr
	Optional bool
}

// DerivedAttribute preserves a derived entity attribute.
type DerivedAttribute struct {
	Name       string
	Type       TypeExpr
	Expression string
}

// InverseAttribute preserves an inverse entity attribute.
type InverseAttribute struct {
	Name string
	Type TypeExpr
	For  string
}

// FunctionDeclaration preserves a FUNCTION declaration.
type FunctionDeclaration struct {
	Name       string
	Parameters []FormalParameter
	ReturnType TypeExpr
	Body       string
	Raw        string
}

// FormalParameter preserves a FUNCTION parameter.
type FormalParameter struct {
	Name        string
	Type        TypeExpr
	ByReference bool
}

// RuleDeclaration preserves a RULE declaration.
type RuleDeclaration struct {
	Name      string
	AppliesTo []string
	Body      string
	Where     []NamedRule
	Raw       string
}

// TypeKind identifies the flavor of an EXPRESS type expression.
type TypeKind string

const (
	TypeKindNamed       TypeKind = "named"
	TypeKindBuiltin     TypeKind = "builtin"
	TypeKindGeneric     TypeKind = "generic"
	TypeKindAggregate   TypeKind = "aggregate"
	TypeKindEnumeration TypeKind = "enumeration"
	TypeKindSelect      TypeKind = "select"
)

// TypeExpr preserves a parsed EXPRESS type expression.
type TypeExpr struct {
	Raw string

	Kind TypeKind
	Name string

	Extensible bool
	Fixed      bool
	Unique     bool
	Optional   bool

	GenericParameter string

	WidthLower string
	WidthUpper string

	Lower   string
	Upper   string
	Element *TypeExpr

	Items []string
}

// IsAggregate reports whether the type expression is an aggregate type.
func (t TypeExpr) IsAggregate() bool {
	return t.Kind == TypeKindAggregate
}

func sortedKeys[T any](m map[string]T) []string {
	if len(m) == 0 {
		return nil
	}
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}

func lookupDeclaration[T any](m map[string]*T, name string) (*T, bool) {
	if decl, ok := m[name]; ok {
		return decl, true
	}

	upper := strings.ToUpper(name)
	for key, decl := range m {
		if strings.ToUpper(key) == upper {
			return decl, true
		}
	}
	return nil, false
}

func cloneSchema(src *Schema) *Schema {
	if src == nil {
		return nil
	}

	dst := &Schema{
		Version:   src.Version,
		Name:      src.Name,
		SourceURL: src.SourceURL,
		Types:     make(map[string]*TypeDeclaration, len(src.Types)),
		Entities:  make(map[string]*EntityDeclaration, len(src.Entities)),
		Functions: make(map[string]*FunctionDeclaration, len(src.Functions)),
		Rules:     make(map[string]*RuleDeclaration, len(src.Rules)),
	}

	for key, value := range src.Types {
		dst.Types[key] = cloneTypeDeclaration(value)
	}
	for key, value := range src.Entities {
		dst.Entities[key] = cloneEntityDeclaration(value)
	}
	for key, value := range src.Functions {
		dst.Functions[key] = cloneFunctionDeclaration(value)
	}
	for key, value := range src.Rules {
		dst.Rules[key] = cloneRuleDeclaration(value)
	}

	return dst
}

func cloneTypeDeclaration(src *TypeDeclaration) *TypeDeclaration {
	if src == nil {
		return nil
	}
	return &TypeDeclaration{
		Name:       src.Name,
		Definition: cloneTypeExpr(src.Definition),
		Where:      slices.Clone(src.Where),
		Raw:        src.Raw,
	}
}

func cloneEntityDeclaration(src *EntityDeclaration) *EntityDeclaration {
	if src == nil {
		return nil
	}

	dst := &EntityDeclaration{
		Name:                src.Name,
		Abstract:            src.Abstract,
		SupertypeExpression: src.SupertypeExpression,
		Supertypes:          slices.Clone(src.Supertypes),
		Unique:              slices.Clone(src.Unique),
		Where:               slices.Clone(src.Where),
		Raw:                 src.Raw,
	}
	if len(src.Attributes) > 0 {
		dst.Attributes = make([]ExplicitAttribute, 0, len(src.Attributes))
		for _, attribute := range src.Attributes {
			dst.Attributes = append(dst.Attributes, ExplicitAttribute{
				Name:     attribute.Name,
				Type:     cloneTypeExpr(attribute.Type),
				Optional: attribute.Optional,
			})
		}
	}
	if len(src.Derived) > 0 {
		dst.Derived = make([]DerivedAttribute, 0, len(src.Derived))
		for _, attribute := range src.Derived {
			dst.Derived = append(dst.Derived, DerivedAttribute{
				Name:       attribute.Name,
				Type:       cloneTypeExpr(attribute.Type),
				Expression: attribute.Expression,
			})
		}
	}
	if len(src.Inverse) > 0 {
		dst.Inverse = make([]InverseAttribute, 0, len(src.Inverse))
		for _, attribute := range src.Inverse {
			dst.Inverse = append(dst.Inverse, InverseAttribute{
				Name: attribute.Name,
				Type: cloneTypeExpr(attribute.Type),
				For:  attribute.For,
			})
		}
	}
	return dst
}

func cloneFunctionDeclaration(src *FunctionDeclaration) *FunctionDeclaration {
	if src == nil {
		return nil
	}

	dst := &FunctionDeclaration{
		Name:       src.Name,
		ReturnType: cloneTypeExpr(src.ReturnType),
		Body:       src.Body,
		Raw:        src.Raw,
	}
	if len(src.Parameters) > 0 {
		dst.Parameters = make([]FormalParameter, 0, len(src.Parameters))
		for _, parameter := range src.Parameters {
			dst.Parameters = append(dst.Parameters, FormalParameter{
				Name:        parameter.Name,
				Type:        cloneTypeExpr(parameter.Type),
				ByReference: parameter.ByReference,
			})
		}
	}
	return dst
}

func cloneRuleDeclaration(src *RuleDeclaration) *RuleDeclaration {
	if src == nil {
		return nil
	}
	return &RuleDeclaration{
		Name:      src.Name,
		AppliesTo: slices.Clone(src.AppliesTo),
		Body:      src.Body,
		Where:     slices.Clone(src.Where),
		Raw:       src.Raw,
	}
}

func cloneTypeExpr(src TypeExpr) TypeExpr {
	dst := src
	dst.Items = slices.Clone(src.Items)
	if src.Element != nil {
		element := cloneTypeExpr(*src.Element)
		dst.Element = &element
	}
	return dst
}
