package ifc

import (
	"fmt"
	"strconv"
	"strings"
)

// ParameterKind identifies the lexical kind of a STEP parameter value.
type ParameterKind uint8

const (
	ParameterKindOmitted ParameterKind = iota
	ParameterKindDerived
	ParameterKindString
	ParameterKindBinary
	ParameterKindInteger
	ParameterKindReal
	ParameterKindEnumeration
	ParameterKindReference
	ParameterKindAggregate
	ParameterKindNamed
)

// Parameter preserves a parsed STEP parameter.
type Parameter interface {
	fmt.Stringer
	Kind() ParameterKind
}

// OmittedParameter preserves the "$" marker.
type OmittedParameter struct{}

func (OmittedParameter) Kind() ParameterKind { return ParameterKindOmitted }
func (OmittedParameter) String() string      { return "$" }

// DerivedParameter preserves the "*" marker.
type DerivedParameter struct{}

func (DerivedParameter) Kind() ParameterKind { return ParameterKindDerived }
func (DerivedParameter) String() string      { return "*" }

// StringParameter preserves a STEP string.
type StringParameter struct {
	Raw     string
	Decoded string
}

func (StringParameter) Kind() ParameterKind { return ParameterKindString }
func (p StringParameter) String() string    { return "'" + p.Raw + "'" }

// BinaryParameter preserves a STEP binary literal.
type BinaryParameter struct {
	Digits string
}

func (BinaryParameter) Kind() ParameterKind { return ParameterKindBinary }
func (p BinaryParameter) String() string    { return `"` + p.Digits + `"` }

// IntegerParameter preserves an integer literal.
type IntegerParameter struct {
	Raw   string
	Value int64
}

func (IntegerParameter) Kind() ParameterKind { return ParameterKindInteger }
func (p IntegerParameter) String() string {
	if p.Raw != "" {
		return p.Raw
	}
	return strconv.FormatInt(p.Value, 10)
}

// RealParameter preserves a real literal.
type RealParameter struct {
	Raw   string
	Value float64
}

func (RealParameter) Kind() ParameterKind { return ParameterKindReal }
func (p RealParameter) String() string {
	if p.Raw != "" {
		return p.Raw
	}
	return strconv.FormatFloat(p.Value, 'g', -1, 64)
}

// EnumerationParameter preserves an enumeration literal such as ".ARCHITECT.".
type EnumerationParameter struct {
	Value string
}

func (EnumerationParameter) Kind() ParameterKind { return ParameterKindEnumeration }
func (p EnumerationParameter) String() string    { return "." + p.Value + "." }

// ReferenceParameter preserves an instance reference such as "#42".
type ReferenceParameter struct {
	ID int64
}

func (ReferenceParameter) Kind() ParameterKind { return ParameterKindReference }
func (p ReferenceParameter) String() string    { return "#" + strconv.FormatInt(p.ID, 10) }

// AggregateParameter preserves a tuple/list/set literal.
type AggregateParameter struct {
	Elements []Parameter
}

func (AggregateParameter) Kind() ParameterKind { return ParameterKindAggregate }
func (p AggregateParameter) String() string {
	parts := make([]string, 0, len(p.Elements))
	for _, element := range p.Elements {
		parts = append(parts, element.String())
	}
	return "(" + strings.Join(parts, ",") + ")"
}

// NamedParameter preserves a named STEP construct such as IFCTEXT('x')
// or an inline entity value.
type NamedParameter struct {
	Name      string
	Arguments []Parameter
}

func (NamedParameter) Kind() ParameterKind { return ParameterKindNamed }
func (p NamedParameter) String() string {
	parts := make([]string, 0, len(p.Arguments))
	for _, argument := range p.Arguments {
		parts = append(parts, argument.String())
	}
	return p.Name + "(" + strings.Join(parts, ",") + ")"
}
