package ifc

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
	"unicode"
)

//go:embed reference/express/*.exp
var embeddedSchemas embed.FS

var (
	schemaFiles = map[Version]struct {
		Path string
		URL  string
	}{
		VersionIFC2X3TC1: {
			Path: "reference/express/IFC2X3_TC1.exp",
			URL:  "https://standards.buildingsmart.org/IFC/RELEASE/IFC2x3/TC1/EXPRESS/IFC2X3_TC1.exp",
		},
		VersionIFC4ADD2TC1: {
			Path: "reference/express/IFC4_ADD2_TC1.exp",
			URL:  "https://standards.buildingsmart.org/IFC/RELEASE/IFC4/ADD2_TC1/EXPRESS/IFC4.exp",
		},
		VersionIFC4X3ADD2: {
			Path: "reference/express/IFC4X3_ADD2.exp",
			URL:  "https://standards.buildingsmart.org/IFC/RELEASE/IFC4_3/HTML/IFC4X3_ADD2.exp",
		},
	}

	schemaCache sync.Map

	functionHeaderPattern = regexp.MustCompile(`(?s)^FUNCTION\s+([A-Za-z0-9_]+)\s*\((.*?)\)\s*:\s*(.*?)\s*;(.*)$`)
	ruleHeaderPattern     = regexp.MustCompile(`(?s)^RULE\s+([A-Za-z0-9_]+)\s+FOR\s*\((.*?)\)\s*;(.*)$`)
)

// ParseSchema reads an EXPRESS schema from r.
func ParseSchema(r io.Reader) (*Schema, error) {
	if r == nil {
		return nil, ErrEmptyInput
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("ifc: read schema: %w", err)
	}
	return ParseSchemaBytes(data)
}

// ParseSchemaBytes reads an EXPRESS schema from data.
func ParseSchemaBytes(data []byte) (*Schema, error) {
	if len(bytes.TrimSpace(data)) == 0 {
		return nil, ErrEmptyInput
	}
	return parseExpressSchema(string(data), VersionUnknown, "")
}

// ParseSchemaFile reads an EXPRESS schema from path.
func ParseSchemaFile(path string) (*Schema, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ifc: read schema %s: %w", path, err)
	}
	return ParseSchemaBytes(data)
}

// SchemaFor loads the embedded official EXPRESS schema for version.
func SchemaFor(version Version) (*Schema, error) {
	if !version.HasSchema() {
		return nil, fmt.Errorf("%w: %s", ErrSchemaUnavailable, version)
	}
	if cached, ok := schemaCache.Load(version); ok {
		return cloneSchema(cached.(*Schema)), nil
	}

	file, ok := schemaFiles[version]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrSchemaUnavailable, version)
	}

	data, err := embeddedSchemas.ReadFile(file.Path)
	if err != nil {
		return nil, fmt.Errorf("ifc: read embedded schema %s: %w", file.Path, err)
	}

	schema, err := parseExpressSchema(string(data), version, file.URL)
	if err != nil {
		return nil, err
	}

	actual, _ := schemaCache.LoadOrStore(version, schema)
	return cloneSchema(actual.(*Schema)), nil
}

type expressBlock struct {
	Kind string
	Raw  string
}

func parseExpressSchema(src string, version Version, sourceURL string) (*Schema, error) {
	clean := stripExpressComments(src)
	schemaName, blocks, err := splitExpressBlocks(clean)
	if err != nil {
		return nil, err
	}

	schema := &Schema{
		Version:   version,
		Name:      schemaName,
		SourceURL: sourceURL,
		Types:     map[string]*TypeDeclaration{},
		Entities:  map[string]*EntityDeclaration{},
		Functions: map[string]*FunctionDeclaration{},
		Rules:     map[string]*RuleDeclaration{},
	}

	for _, block := range blocks {
		switch block.Kind {
		case "TYPE":
			decl, err := parseTypeDeclaration(block.Raw)
			if err != nil {
				return nil, err
			}
			schema.Types[decl.Name] = decl
		case "ENTITY":
			decl, err := parseEntityDeclaration(block.Raw)
			if err != nil {
				return nil, err
			}
			schema.Entities[decl.Name] = decl
		case "FUNCTION":
			decl, err := parseFunctionDeclaration(block.Raw)
			if err != nil {
				return nil, err
			}
			schema.Functions[decl.Name] = decl
		case "RULE":
			decl, err := parseRuleDeclaration(block.Raw)
			if err != nil {
				return nil, err
			}
			schema.Rules[decl.Name] = decl
		}
	}

	return schema, nil
}

func stripExpressComments(src string) string {
	var out strings.Builder
	for i := 0; i < len(src); {
		if i+1 < len(src) && src[i] == '(' && src[i+1] == '*' {
			i += 2
			for i < len(src) {
				if i+1 < len(src) && src[i] == '*' && src[i+1] == ')' {
					i += 2
					break
				}
				if src[i] == '\n' {
					out.WriteByte('\n')
				}
				i++
			}
			continue
		}
		out.WriteByte(src[i])
		i++
	}
	return out.String()
}

func splitExpressBlocks(src string) (string, []expressBlock, error) {
	lines := strings.SplitAfter(src, "\n")
	var (
		schemaName string
		blocks     []expressBlock
		builder    strings.Builder
		kind       string
		endMarker  string
		inBlock    bool
	)

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if !inBlock {
			switch {
			case trimmed == "":
				continue
			case strings.HasPrefix(trimmed, "SCHEMA "):
				schemaName = parseSchemaName(trimmed)
			case strings.HasPrefix(trimmed, "TYPE "):
				inBlock = true
				kind = "TYPE"
				endMarker = "END_TYPE;"
				builder.Reset()
				builder.WriteString(line)
			case strings.HasPrefix(trimmed, "ENTITY "):
				inBlock = true
				kind = "ENTITY"
				endMarker = "END_ENTITY;"
				builder.Reset()
				builder.WriteString(line)
			case strings.HasPrefix(trimmed, "FUNCTION "):
				inBlock = true
				kind = "FUNCTION"
				endMarker = "END_FUNCTION;"
				builder.Reset()
				builder.WriteString(line)
			case strings.HasPrefix(trimmed, "RULE "):
				inBlock = true
				kind = "RULE"
				endMarker = "END_RULE;"
				builder.Reset()
				builder.WriteString(line)
			case trimmed == "END_SCHEMA;":
				return schemaName, blocks, nil
			}
			continue
		}

		builder.WriteString(line)
		if trimmed == endMarker {
			blocks = append(blocks, expressBlock{Kind: kind, Raw: strings.TrimSpace(builder.String())})
			inBlock = false
		}
	}

	if inBlock {
		return "", nil, fmt.Errorf("%w: unterminated EXPRESS %s block", ErrInvalidExchangeStructure, kind)
	}
	if schemaName == "" {
		return "", nil, fmt.Errorf("%w: missing SCHEMA declaration", ErrInvalidExchangeStructure)
	}
	return schemaName, blocks, nil
}

func parseSchemaName(line string) string {
	body := strings.TrimSpace(strings.TrimPrefix(line, "SCHEMA"))
	body = strings.TrimSpace(strings.TrimSuffix(body, ";"))
	return body
}

func parseTypeDeclaration(raw string) (*TypeDeclaration, error) {
	body := strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(raw), "END_TYPE;"))
	lines := splitNonEmptyLines(body)
	if len(lines) == 0 {
		return nil, fmt.Errorf("%w: empty TYPE block", ErrInvalidExchangeStructure)
	}

	first := strings.TrimSpace(lines[0])
	name, rest, ok := cutTypeHeader(first)
	if !ok {
		return nil, fmt.Errorf("%w: malformed TYPE header %q", ErrInvalidExchangeStructure, first)
	}

	definitionParts := []string{rest}
	i := 1
	for ; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if strings.EqualFold(line, "WHERE") {
			i++
			break
		}
		definitionParts = append(definitionParts, line)
	}

	definition := strings.TrimSpace(trimTrailingSemicolon(strings.Join(definitionParts, "\n")))
	where := parseNamedRules(strings.Join(lines[i:], "\n"))
	typeExpr, err := parseTypeExpr(definition)
	if err != nil {
		return nil, fmt.Errorf("ifc: parse TYPE %s: %w", name, err)
	}

	return &TypeDeclaration{
		Name:       name,
		Definition: typeExpr,
		Where:      where,
		Raw:        strings.TrimSpace(raw),
	}, nil
}

func parseEntityDeclaration(raw string) (*EntityDeclaration, error) {
	body := strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(raw), "END_ENTITY;"))
	lines := splitNonEmptyLines(body)
	if len(lines) == 0 {
		return nil, fmt.Errorf("%w: empty ENTITY block", ErrInvalidExchangeStructure)
	}

	name, initial, ok := cutEntityHeader(strings.TrimSpace(lines[0]))
	if !ok {
		return nil, fmt.Errorf("%w: malformed ENTITY header %q", ErrInvalidExchangeStructure, lines[0])
	}

	headerParts := []string{initial}
	i := 1
	if !strings.Contains(initial, ";") {
		for ; i < len(lines); i++ {
			line := strings.TrimSpace(lines[i])
			headerParts = append(headerParts, line)
			if strings.Contains(line, ";") {
				i++
				break
			}
		}
	}

	headerText := strings.TrimSpace(trimTrailingSemicolon(strings.Join(headerParts, " ")))
	entity := &EntityDeclaration{
		Name:                name,
		Abstract:            strings.Contains(strings.ToUpper(headerText), "ABSTRACT"),
		SupertypeExpression: extractClause(headerText, "SUPERTYPE OF", "SUBTYPE OF"),
		Supertypes:          parseSubtypeList(extractClause(headerText, "SUBTYPE OF", "")),
		Raw:                 strings.TrimSpace(raw),
	}

	section := "ATTRIBUTES"
	for ; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		switch strings.ToUpper(line) {
		case "DERIVE":
			section = "DERIVE"
			continue
		case "INVERSE":
			section = "INVERSE"
			continue
		case "UNIQUE":
			section = "UNIQUE"
			continue
		case "WHERE":
			section = "WHERE"
			continue
		}

		statement := line
		for !strings.HasSuffix(strings.TrimSpace(statement), ";") && i+1 < len(lines) {
			i++
			statement += "\n" + strings.TrimSpace(lines[i])
		}

		switch section {
		case "ATTRIBUTES":
			attribute, err := parseExplicitAttribute(statement)
			if err != nil {
				return nil, err
			}
			entity.Attributes = append(entity.Attributes, attribute)
		case "DERIVE":
			attribute, err := parseDerivedAttribute(statement)
			if err != nil {
				return nil, err
			}
			entity.Derived = append(entity.Derived, attribute)
		case "INVERSE":
			attribute, err := parseInverseAttribute(statement)
			if err != nil {
				return nil, err
			}
			entity.Inverse = append(entity.Inverse, attribute)
		case "UNIQUE":
			rules := parseNamedRules(statement)
			entity.Unique = append(entity.Unique, rules...)
		case "WHERE":
			rules := parseNamedRules(statement)
			entity.Where = append(entity.Where, rules...)
		}
	}

	return entity, nil
}

func parseFunctionDeclaration(raw string) (*FunctionDeclaration, error) {
	match := functionHeaderPattern.FindStringSubmatch(strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(raw), "END_FUNCTION;")))
	if match == nil {
		return nil, fmt.Errorf("%w: malformed FUNCTION block", ErrInvalidExchangeStructure)
	}

	parameters, err := parseFormalParameters(match[2])
	if err != nil {
		return nil, err
	}
	returnType, err := parseTypeExpr(strings.TrimSpace(match[3]))
	if err != nil {
		return nil, fmt.Errorf("ifc: parse FUNCTION %s return type: %w", match[1], err)
	}

	return &FunctionDeclaration{
		Name:       match[1],
		Parameters: parameters,
		ReturnType: returnType,
		Body:       strings.TrimSpace(match[4]),
		Raw:        strings.TrimSpace(raw),
	}, nil
}

func parseRuleDeclaration(raw string) (*RuleDeclaration, error) {
	match := ruleHeaderPattern.FindStringSubmatch(strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(raw), "END_RULE;")))
	if match == nil {
		return nil, fmt.Errorf("%w: malformed RULE block", ErrInvalidExchangeStructure)
	}

	bodyText, whereText, hasWhere := splitSection(strings.TrimSpace(match[3]), "WHERE")

	rule := &RuleDeclaration{
		Name:      match[1],
		AppliesTo: splitTopLevelList(match[2], ','),
		Body:      strings.TrimSpace(bodyText),
		Raw:       strings.TrimSpace(raw),
	}
	if hasWhere {
		rule.Where = parseNamedRules(whereText)
	}
	return rule, nil
}

func cutTypeHeader(line string) (string, string, bool) {
	if !strings.HasPrefix(line, "TYPE ") {
		return "", "", false
	}
	body := strings.TrimSpace(strings.TrimPrefix(line, "TYPE"))
	name, rest, ok := strings.Cut(body, "=")
	if !ok {
		return "", "", false
	}
	return strings.TrimSpace(name), strings.TrimSpace(rest), true
}

func cutEntityHeader(line string) (string, string, bool) {
	if !strings.HasPrefix(line, "ENTITY ") {
		return "", "", false
	}
	body := strings.TrimSpace(strings.TrimPrefix(line, "ENTITY"))
	name := nextWord(body)
	if name == "" {
		return "", "", false
	}
	return name, strings.TrimSpace(body[len(name):]), true
}

func parseExplicitAttribute(statement string) (ExplicitAttribute, error) {
	statement = strings.TrimSpace(trimTrailingSemicolon(statement))
	name, rest, ok := strings.Cut(statement, ":")
	if !ok {
		return ExplicitAttribute{}, fmt.Errorf("%w: malformed attribute %q", ErrInvalidExchangeStructure, statement)
	}

	rest = strings.TrimSpace(rest)
	optional := false
	if strings.HasPrefix(strings.ToUpper(rest), "OPTIONAL ") {
		optional = true
		rest = strings.TrimSpace(rest[len("OPTIONAL "):])
	}

	typeExpr, err := parseTypeExpr(rest)
	if err != nil {
		return ExplicitAttribute{}, err
	}

	return ExplicitAttribute{
		Name:     strings.TrimSpace(name),
		Type:     typeExpr,
		Optional: optional,
	}, nil
}

func parseDerivedAttribute(statement string) (DerivedAttribute, error) {
	statement = strings.TrimSpace(trimTrailingSemicolon(statement))
	name, rest, ok := strings.Cut(statement, ":")
	if !ok {
		return DerivedAttribute{}, fmt.Errorf("%w: malformed derived attribute %q", ErrInvalidExchangeStructure, statement)
	}
	typeText, expr, ok := strings.Cut(rest, ":=")
	if !ok {
		return DerivedAttribute{}, fmt.Errorf("%w: malformed derived attribute expression %q", ErrInvalidExchangeStructure, statement)
	}
	typeExpr, err := parseTypeExpr(strings.TrimSpace(typeText))
	if err != nil {
		return DerivedAttribute{}, err
	}
	return DerivedAttribute{
		Name:       strings.TrimSpace(name),
		Type:       typeExpr,
		Expression: strings.TrimSpace(expr),
	}, nil
}

func parseInverseAttribute(statement string) (InverseAttribute, error) {
	statement = strings.TrimSpace(trimTrailingSemicolon(statement))
	name, rest, ok := strings.Cut(statement, ":")
	if !ok {
		return InverseAttribute{}, fmt.Errorf("%w: malformed inverse attribute %q", ErrInvalidExchangeStructure, statement)
	}
	index := strings.LastIndex(strings.ToUpper(rest), " FOR ")
	if index == -1 {
		return InverseAttribute{}, fmt.Errorf("%w: malformed inverse attribute FOR clause %q", ErrInvalidExchangeStructure, statement)
	}
	typeExpr, err := parseTypeExpr(strings.TrimSpace(rest[:index]))
	if err != nil {
		return InverseAttribute{}, err
	}
	return InverseAttribute{
		Name: strings.TrimSpace(name),
		Type: typeExpr,
		For:  strings.TrimSpace(rest[index+len(" FOR "):]),
	}, nil
}

func parseNamedRules(text string) []NamedRule {
	statements := collectStatements(text)
	out := make([]NamedRule, 0, len(statements))
	for _, statement := range statements {
		statement = strings.TrimSpace(trimTrailingSemicolon(statement))
		name, expr, ok := strings.Cut(statement, ":")
		if !ok {
			continue
		}
		out = append(out, NamedRule{
			Name:       strings.TrimSpace(name),
			Expression: strings.TrimSpace(expr),
		})
	}
	return out
}

func parseFormalParameters(text string) ([]FormalParameter, error) {
	groups := splitTopLevelList(text, ';')
	parameters := []FormalParameter{}
	for _, group := range groups {
		group = strings.TrimSpace(group)
		if group == "" {
			continue
		}

		byReference := false
		upper := strings.ToUpper(group)
		if strings.HasPrefix(upper, "VAR ") {
			byReference = true
			group = strings.TrimSpace(group[len("VAR "):])
		}

		namesPart, typePart, ok := strings.Cut(group, ":")
		if !ok {
			return nil, fmt.Errorf("%w: malformed function parameter group %q", ErrInvalidExchangeStructure, group)
		}
		typeExpr, err := parseTypeExpr(strings.TrimSpace(typePart))
		if err != nil {
			return nil, err
		}
		for _, name := range splitTopLevelList(namesPart, ',') {
			name = strings.TrimSpace(name)
			if name == "" {
				continue
			}
			parameters = append(parameters, FormalParameter{
				Name:        name,
				Type:        typeExpr,
				ByReference: byReference,
			})
		}
	}
	return parameters, nil
}

func parseTypeExpr(text string) (TypeExpr, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return TypeExpr{}, nil
	}
	parser := newTypeExprParser(text)
	expr, err := parser.parse()
	if err != nil {
		return TypeExpr{}, fmt.Errorf("ifc: parse type expression %q: %w", text, err)
	}
	if parser.pos != len(parser.tokens) {
		return TypeExpr{}, fmt.Errorf("ifc: parse type expression %q: trailing tokens %v", text, parser.tokens[parser.pos:])
	}
	expr.Raw = text
	return expr, nil
}

type typeExprParser struct {
	tokens []string
	pos    int
}

func newTypeExprParser(text string) *typeExprParser {
	return &typeExprParser{tokens: tokenizeTypeExpr(text)}
}

func (p *typeExprParser) parse() (TypeExpr, error) {
	extensible := p.acceptWord("EXTENSIBLE")
	head, ok := p.nextWord()
	if !ok {
		return TypeExpr{}, fmt.Errorf("%w: missing type head", ErrInvalidExchangeStructure)
	}

	upperHead := strings.ToUpper(head)
	switch upperHead {
	case "ENUMERATION":
		if !p.acceptWord("OF") {
			return TypeExpr{}, fmt.Errorf("%w: expected OF after ENUMERATION", ErrInvalidExchangeStructure)
		}
		items, err := p.parseNameList()
		if err != nil {
			return TypeExpr{}, err
		}
		return TypeExpr{
			Kind:       TypeKindEnumeration,
			Name:       "ENUMERATION",
			Extensible: extensible,
			Items:      items,
		}, nil
	case "SELECT":
		items, err := p.parseNameList()
		if err != nil {
			return TypeExpr{}, err
		}
		return TypeExpr{
			Kind:       TypeKindSelect,
			Name:       "SELECT",
			Extensible: extensible,
			Items:      items,
		}, nil
	case "ARRAY", "LIST", "SET", "BAG":
		var lower, upper string
		if p.accept("[") {
			value, ok := p.nextToken()
			if !ok || !p.accept(":") {
				return TypeExpr{}, fmt.Errorf("%w: malformed aggregate bounds", ErrInvalidExchangeStructure)
			}
			maximum, ok := p.nextToken()
			if !ok || !p.accept("]") {
				return TypeExpr{}, fmt.Errorf("%w: malformed aggregate bounds", ErrInvalidExchangeStructure)
			}
			lower = value
			upper = maximum
		}
		if !p.acceptWord("OF") {
			return TypeExpr{}, fmt.Errorf("%w: malformed aggregate type", ErrInvalidExchangeStructure)
		}
		elementOptional := p.acceptWord("OPTIONAL")
		elementUnique := p.acceptWord("UNIQUE")
		element, err := p.parse()
		if err != nil {
			return TypeExpr{}, err
		}
		return TypeExpr{
			Kind:     TypeKindAggregate,
			Name:     upperHead,
			Lower:    lower,
			Upper:    upper,
			Optional: elementOptional,
			Unique:   elementUnique,
			Element:  &element,
		}, nil
	case "GENERIC":
		expr := TypeExpr{
			Kind: TypeKindGeneric,
			Name: "GENERIC",
		}
		if p.accept(":") {
			value, ok := p.nextWord()
			if !ok {
				return TypeExpr{}, fmt.Errorf("%w: malformed generic type parameter", ErrInvalidExchangeStructure)
			}
			expr.GenericParameter = value
		}
		return expr, nil
	default:
		expr := TypeExpr{Name: head}
		if isBuiltinType(upperHead) {
			expr.Kind = TypeKindBuiltin
			expr.Name = upperHead
		} else {
			expr.Kind = TypeKindNamed
		}

		if p.accept("(") {
			lower, ok := p.nextToken()
			if !ok {
				return TypeExpr{}, fmt.Errorf("%w: missing width lower bound", ErrInvalidExchangeStructure)
			}
			upper := lower
			if p.accept(":") {
				value, ok := p.nextToken()
				if !ok {
					return TypeExpr{}, fmt.Errorf("%w: missing width upper bound", ErrInvalidExchangeStructure)
				}
				upper = value
			}
			if !p.accept(")") {
				return TypeExpr{}, fmt.Errorf("%w: missing closing ')' in width", ErrInvalidExchangeStructure)
			}
			expr.WidthLower = lower
			expr.WidthUpper = upper
		}

		expr.Fixed = p.acceptWord("FIXED")
		return expr, nil
	}
}

func (p *typeExprParser) parseNameList() ([]string, error) {
	if !p.accept("(") {
		return nil, fmt.Errorf("%w: expected '('", ErrInvalidExchangeStructure)
	}
	items := []string{}
	for {
		token, ok := p.nextToken()
		if !ok {
			return nil, fmt.Errorf("%w: unterminated name list", ErrInvalidExchangeStructure)
		}
		items = append(items, strings.TrimSpace(token))
		if p.accept(")") {
			return items, nil
		}
		if !p.accept(",") {
			return nil, fmt.Errorf("%w: expected ',' in name list", ErrInvalidExchangeStructure)
		}
	}
}

func (p *typeExprParser) accept(token string) bool {
	if p.pos >= len(p.tokens) || p.tokens[p.pos] != token {
		return false
	}
	p.pos++
	return true
}

func (p *typeExprParser) acceptWord(word string) bool {
	if p.pos >= len(p.tokens) || !strings.EqualFold(p.tokens[p.pos], word) {
		return false
	}
	p.pos++
	return true
}

func (p *typeExprParser) nextToken() (string, bool) {
	if p.pos >= len(p.tokens) {
		return "", false
	}
	token := p.tokens[p.pos]
	p.pos++
	return token, true
}

func (p *typeExprParser) nextWord() (string, bool) {
	if p.pos >= len(p.tokens) {
		return "", false
	}
	token := p.tokens[p.pos]
	if token == "(" || token == ")" || token == "[" || token == "]" || token == ":" || token == "," {
		return "", false
	}
	p.pos++
	return token, true
}

func tokenizeTypeExpr(text string) []string {
	tokens := []string{}
	for i := 0; i < len(text); {
		r := rune(text[i])
		if unicode.IsSpace(r) {
			i++
			continue
		}
		switch text[i] {
		case '(', ')', '[', ']', ':', ',':
			tokens = append(tokens, text[i:i+1])
			i++
			continue
		}

		start := i
		for i < len(text) {
			switch text[i] {
			case '(', ')', '[', ']', ':', ',', ' ', '\t', '\r', '\n':
				goto emit
			}
			i++
		}
	emit:
		tokens = append(tokens, text[start:i])
	}
	return tokens
}

func isBuiltinType(name string) bool {
	switch name {
	case "REAL", "NUMBER", "INTEGER", "BOOLEAN", "LOGICAL", "STRING", "BINARY":
		return true
	default:
		return false
	}
}

func splitNonEmptyLines(text string) []string {
	lines := strings.Split(text, "\n")
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		out = append(out, line)
	}
	return out
}

func trimTrailingSemicolon(text string) string {
	return strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(text), ";"))
}

func collectStatements(text string) []string {
	lines := splitNonEmptyLines(text)
	if len(lines) == 0 {
		return nil
	}

	statements := []string{}
	var current strings.Builder
	for _, line := range lines {
		if current.Len() > 0 {
			current.WriteByte('\n')
		}
		current.WriteString(strings.TrimSpace(line))
		if strings.HasSuffix(strings.TrimSpace(line), ";") {
			statements = append(statements, current.String())
			current.Reset()
		}
	}
	if current.Len() > 0 {
		statements = append(statements, current.String())
	}
	return statements
}

func splitSection(text string, keyword string) (string, string, bool) {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if strings.EqualFold(strings.TrimSpace(line), keyword) {
			return strings.Join(lines[:i], "\n"), strings.Join(lines[i+1:], "\n"), true
		}
	}
	return text, "", false
}

func extractClause(text string, startKeyword string, endKeyword string) string {
	upper := strings.ToUpper(text)
	start := strings.Index(upper, startKeyword)
	if start == -1 {
		return ""
	}
	rest := text[start+len(startKeyword):]
	if endKeyword != "" {
		upperRest := strings.ToUpper(rest)
		if end := strings.Index(upperRest, endKeyword); end != -1 {
			rest = rest[:end]
		}
	}
	return strings.TrimSpace(rest)
}

func parseSubtypeList(clause string) []string {
	clause = strings.TrimSpace(clause)
	if clause == "" {
		return nil
	}
	if strings.HasPrefix(clause, "(") && strings.HasSuffix(clause, ")") {
		clause = clause[1 : len(clause)-1]
	}
	return splitTopLevelList(clause, ',')
}

func splitTopLevelList(text string, delimiter rune) []string {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil
	}

	var (
		items        []string
		current      strings.Builder
		parenDepth   int
		bracketDepth int
	)

	for _, r := range text {
		switch r {
		case '(':
			parenDepth++
		case ')':
			parenDepth--
		case '[':
			bracketDepth++
		case ']':
			bracketDepth--
		}

		if r == delimiter && parenDepth == 0 && bracketDepth == 0 {
			items = append(items, strings.TrimSpace(current.String()))
			current.Reset()
			continue
		}
		current.WriteRune(r)
	}
	if current.Len() > 0 {
		items = append(items, strings.TrimSpace(current.String()))
	}
	return items
}

func nextWord(text string) string {
	text = strings.TrimSpace(text)
	for i, r := range text {
		if unicode.IsSpace(r) || r == ';' {
			return text[:i]
		}
	}
	return text
}
