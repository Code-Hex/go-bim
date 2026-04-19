package stbridge

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html/charset"
)

var (
	// ErrEmptyXML reports that no XML payload was available after sanitization.
	ErrEmptyXML = errors.New("stbridge: empty XML document")
	// ErrInvalidRoot reports that the root element is not ST_BRIDGE.
	ErrInvalidRoot = errors.New("stbridge: root element must be ST_BRIDGE")
	// ErrMissingVersion reports that the ST_BRIDGE root omits the version attribute.
	ErrMissingVersion = errors.New("stbridge: ST_BRIDGE version attribute is required")
	// ErrInvalidNamespace reports that a 2.x ST-Bridge document uses the wrong XML namespace.
	ErrInvalidNamespace = errors.New("stbridge: invalid ST-Bridge XML namespace")
	// ErrNilDocument reports that a nil document receiver was used.
	ErrNilDocument = errors.New("stbridge: nil document")
	// ErrSchemaUnavailable reports that no embedded schema metadata exists for a version.
	ErrSchemaUnavailable = errors.New("stbridge: schema metadata is unavailable for this version")
)

var declarationAttributePattern = regexp.MustCompile(`([A-Za-z_:][A-Za-z0-9_.:-]*)\s*=\s*"([^"]*)"|([A-Za-z_:][A-Za-z0-9_.:-]*)\s*=\s*'([^']*)'`)

// Parse reads a ST-Bridge document from r.
func Parse(r io.Reader) (*Document, error) {
	if r == nil {
		return nil, ErrEmptyXML
	}

	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("stbridge: read XML: %w", err)
	}

	parsed, err := parseXML(data)
	if err != nil {
		return nil, err
	}

	if parsed.root.Name.Local != "ST_BRIDGE" {
		return nil, ErrInvalidRoot
	}

	rawVersion, ok := parsed.root.Attr("version")
	if !ok {
		return nil, ErrMissingVersion
	}
	version := NormalizeVersion(rawVersion)
	if version.HasSchema() && parsed.root.Name.Space != "https://www.building-smart.or.jp/dl" {
		return nil, fmt.Errorf("%w: %q", ErrInvalidNamespace, parsed.root.Name.Space)
	}

	return &Document{
		Declaration: parsed.declaration,
		RawVersion:  rawVersion,
		Version:     version,
		Namespace:   parsed.root.Name.Space,
		Leading:     parsed.leading,
		Root:        parsed.root,
		Trailing:    parsed.trailing,
	}, nil
}

// ParseBytes reads a ST-Bridge document from data.
func ParseBytes(data []byte) (*Document, error) {
	return Parse(bytes.NewReader(data))
}

// ParseFile reads a ST-Bridge document from path.
func ParseFile(path string) (*Document, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("stbridge: read %s: %w", path, err)
	}
	return ParseBytes(data)
}

type parsedXML struct {
	declaration Declaration
	leading     []*Node
	root        *Node
	trailing    []*Node
}

func parseXML(data []byte) (*parsedXML, error) {
	clean := sanitizeXML(data)
	if len(clean) == 0 {
		return nil, ErrEmptyXML
	}

	decoder := xml.NewDecoder(bytes.NewReader(clean))
	decoder.CharsetReader = charset.NewReaderLabel

	result := &parsedXML{}
	stack := []*Node{}

	for {
		token, err := decoder.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("stbridge: decode XML: %w", err)
		}

		switch token := token.(type) {
		case xml.ProcInst:
			if token.Target == "xml" {
				result.declaration = parseDeclaration(token.Inst)
				continue
			}
			appendNode(&stack, &Node{
				Type:    NodeTypeProcInst,
				Name:    xml.Name{Local: token.Target},
				Content: string(token.Inst),
			}, result)
		case xml.Directive:
			appendNode(&stack, &Node{
				Type:    NodeTypeDirective,
				Content: string(token),
			}, result)
		case xml.Comment:
			appendNode(&stack, &Node{
				Type:    NodeTypeComment,
				Content: string(token),
			}, result)
		case xml.CharData:
			if len(stack) == 0 && strings.TrimSpace(string(token)) != "" {
				return nil, fmt.Errorf("stbridge: unexpected top-level character data")
			}
			appendNode(&stack, &Node{
				Type:    NodeTypeText,
				Content: string(token),
			}, result)
		case xml.StartElement:
			if len(stack) == 0 && result.root != nil {
				return nil, fmt.Errorf("stbridge: multiple root elements are not allowed")
			}
			node := &Node{
				Type:  NodeTypeElement,
				Name:  token.Name,
				Attrs: make([]Attribute, 0, len(token.Attr)),
			}
			for _, attr := range token.Attr {
				node.Attrs = append(node.Attrs, Attribute{Name: attr.Name, Value: attr.Value})
			}
			appendNode(&stack, node, result)
			stack = append(stack, node)
		case xml.EndElement:
			if len(stack) == 0 {
				return nil, fmt.Errorf("stbridge: unexpected end element %q", token.Name.Local)
			}
			stack = stack[:len(stack)-1]
		}
	}

	if result.root == nil {
		return nil, ErrEmptyXML
	}

	return result, nil
}

func appendNode(stack *[]*Node, node *Node, result *parsedXML) {
	if len(*stack) == 0 {
		if node.Type == NodeTypeElement && result.root == nil {
			result.root = node
			return
		}
		if result.root == nil {
			result.leading = append(result.leading, node)
		} else {
			result.trailing = append(result.trailing, node)
		}
		return
	}
	parent := (*stack)[len(*stack)-1]
	parent.Children = append(parent.Children, node)
}

func sanitizeXML(data []byte) []byte {
	if len(bytes.TrimSpace(data)) == 0 {
		return nil
	}

	index := bytes.IndexByte(data, '<')
	if index == -1 {
		return nil
	}

	return data[index:]
}

func parseDeclaration(inst []byte) Declaration {
	decl := Declaration{}
	matches := declarationAttributePattern.FindAllSubmatch(inst, -1)
	for _, match := range matches {
		key := string(match[1])
		value := string(match[2])
		if key == "" {
			key = string(match[3])
			value = string(match[4])
		}
		switch key {
		case "version":
			decl.Version = value
		case "encoding":
			decl.Encoding = value
		case "standalone":
			decl.Standalone = value
		}
	}
	return decl
}
