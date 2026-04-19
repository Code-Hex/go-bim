package stbridge

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// Declaration stores the XML declaration when it exists.
type Declaration struct {
	Version    string
	Encoding   string
	Standalone string
}

// NodeType identifies the XML node kind.
type NodeType uint8

const (
	NodeTypeElement NodeType = iota
	NodeTypeText
	NodeTypeComment
	NodeTypeDirective
	NodeTypeProcInst
)

// Attribute preserves an XML attribute.
type Attribute struct {
	Name  xml.Name
	Value string
}

// Node preserves an XML node in document order.
type Node struct {
	Type     NodeType
	Name     xml.Name
	Attrs    []Attribute
	Content  string
	Children []*Node
}

// Attr reports the first attribute whose local name matches local.
//
// Use the generated Attr_* constants when you want editor completion for the
// published ST-Bridge attribute names.
func (n *Node) Attr(local AttributeName) (string, bool) {
	if n == nil {
		return "", false
	}

	for _, attr := range n.Attrs {
		if attr.Name.Local == local {
			return attr.Value, true
		}
	}

	return "", false
}

// ElementChildren returns the element children in order.
func (n *Node) ElementChildren() []*Node {
	if n == nil {
		return nil
	}

	children := make([]*Node, 0, len(n.Children))
	for _, child := range n.Children {
		if child.Type == NodeTypeElement {
			children = append(children, child)
		}
	}
	return children
}

// Child returns the first element child whose local name matches local.
//
// Use the generated Element_* constants when you want editor completion for
// the published ST-Bridge element names.
func (n *Node) Child(local ElementName) *Node {
	if n == nil {
		return nil
	}

	for _, child := range n.Children {
		if child.Type == NodeTypeElement && child.Name.Local == local {
			return child
		}
	}

	return nil
}

// ChildrenByName returns all element children whose local name matches local.
func (n *Node) ChildrenByName(local ElementName) []*Node {
	if n == nil {
		return nil
	}

	children := []*Node{}
	for _, child := range n.Children {
		if child.Type == NodeTypeElement && child.Name.Local == local {
			children = append(children, child)
		}
	}
	return children
}

// StringAttr returns the named attribute as a raw string.
func (n *Node) StringAttr(name AttributeName) (string, bool) {
	return n.Attr(name)
}

// Float64Attr returns the named attribute parsed as float64.
func (n *Node) Float64Attr(name AttributeName) (float64, bool, error) {
	return float64Attr(n, name)
}

// Int64Attr returns the named attribute parsed as int64.
func (n *Node) Int64Attr(name AttributeName) (int64, bool, error) {
	raw, ok := n.Attr(name)
	if !ok {
		return 0, false, nil
	}

	value, err := strconv.ParseInt(strings.TrimSpace(raw), 10, 64)
	if err != nil {
		return 0, true, invalidAttrValueError(name, raw, "int64", err)
	}
	return value, true, nil
}

// BoolAttr returns the named attribute parsed as a boolean.
func (n *Node) BoolAttr(name AttributeName) (bool, bool, error) {
	raw, ok := n.Attr(name)
	if !ok {
		return false, false, nil
	}

	value, err := strconv.ParseBool(strings.TrimSpace(raw))
	if err != nil {
		return false, true, invalidAttrValueError(name, raw, "bool", err)
	}
	return value, true, nil
}

// TextContent returns the concatenated text under the node.
func (n *Node) TextContent() string {
	if n == nil {
		return ""
	}

	var b strings.Builder
	n.appendText(&b)
	return b.String()
}

func (n *Node) appendText(b *strings.Builder) {
	if n.Type == NodeTypeText {
		b.WriteString(n.Content)
		return
	}
	for _, child := range n.Children {
		child.appendText(b)
	}
}

// Document is a parsed ST-Bridge XML document.
type Document struct {
	Declaration Declaration
	RawVersion  string
	Version     Version
	Namespace   string
	Leading     []*Node
	Root        *Node
	Trailing    []*Node
}

// Section returns the first top-level section with the given name.
//
// Use Section_* constants when you want completion for the well-known
// ST_BRIDGE sections.
func (d *Document) Section(name SectionName) *Node {
	if d == nil || d.Root == nil {
		return nil
	}
	return d.Root.Child(name)
}

// Common returns the StbCommon section under ST_BRIDGE.
func (d *Document) Common() *Node {
	return d.Section(Section_StbCommon)
}

// Model returns the StbModel section under ST_BRIDGE.
func (d *Document) Model() *Node {
	return d.Section(Section_StbModel)
}

// Extensions returns the optional StbExtensions section under ST_BRIDGE.
func (d *Document) Extensions() *Node {
	return d.Section(Section_StbExtensions)
}

// ExportInformation returns the optional StbExportInformation section under ST_BRIDGE.
//
// This section was added in ST-Bridge 2.1.0.
func (d *Document) ExportInformation() *Node {
	return d.Section(Section_StbExportInformation)
}

// CalData returns the optional StbCalData section under ST_BRIDGE.
//
// This section was added in ST-Bridge 2.0.2.
func (d *Document) CalData() *Node {
	return d.Section(Section_StbCalData)
}

// AnaModels returns the optional StbAnaModels section under ST_BRIDGE.
//
// This section was added in ST-Bridge 2.0.2.
func (d *Document) AnaModels() *Node {
	return d.Section(Section_StbAnaModels)
}

// Schema returns the versioned schema metadata for the document when available.
func (d *Document) Schema() (*Schema, error) {
	if d == nil {
		return nil, ErrNilDocument
	}
	return SchemaFor(d.Version)
}
