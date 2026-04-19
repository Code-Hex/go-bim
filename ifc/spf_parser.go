package ifc

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type spfTokenKind uint8

const (
	spfTokenEOF spfTokenKind = iota
	spfTokenWord
	spfTokenString
	spfTokenBinary
	spfTokenInteger
	spfTokenReal
	spfTokenReference
	spfTokenEnumeration
	spfTokenDollar
	spfTokenStar
	spfTokenLParen
	spfTokenRParen
	spfTokenComma
	spfTokenSemicolon
	spfTokenEquals
)

type spfToken struct {
	kind spfTokenKind
	raw  string
	text string
	pos  int
}

type spfLexer struct {
	data   []byte
	pos    int
	peeked *spfToken
}

func newSPFLexer(data []byte) *spfLexer {
	return &spfLexer{data: data}
}

func (l *spfLexer) next() (spfToken, error) {
	if l.peeked != nil {
		token := *l.peeked
		l.peeked = nil
		return token, nil
	}
	return l.scan()
}

func (l *spfLexer) peek() (spfToken, error) {
	if l.peeked != nil {
		return *l.peeked, nil
	}
	token, err := l.scan()
	if err != nil {
		return spfToken{}, err
	}
	l.peeked = &token
	return token, nil
}

func (l *spfLexer) scan() (spfToken, error) {
	if err := l.skipWhitespaceAndComments(); err != nil {
		return spfToken{}, err
	}
	if l.pos >= len(l.data) {
		return spfToken{kind: spfTokenEOF, pos: l.pos}, nil
	}

	start := l.pos
	remaining := l.data[l.pos:]

	switch {
	case bytes.HasPrefix(remaining, []byte("ISO-10303-21")):
		l.pos += len("ISO-10303-21")
		return spfToken{kind: spfTokenWord, raw: "ISO-10303-21", text: "ISO-10303-21", pos: start}, nil
	case bytes.HasPrefix(remaining, []byte("END-ISO-10303-21")):
		l.pos += len("END-ISO-10303-21")
		return spfToken{kind: spfTokenWord, raw: "END-ISO-10303-21", text: "END-ISO-10303-21", pos: start}, nil
	}

	switch l.data[l.pos] {
	case '\'':
		return l.scanString()
	case '"':
		return l.scanBinary()
	case '#':
		return l.scanReference()
	case '.':
		return l.scanEnumeration()
	case '$':
		l.pos++
		return spfToken{kind: spfTokenDollar, raw: "$", text: "$", pos: start}, nil
	case '*':
		l.pos++
		return spfToken{kind: spfTokenStar, raw: "*", text: "*", pos: start}, nil
	case '(':
		l.pos++
		return spfToken{kind: spfTokenLParen, raw: "(", text: "(", pos: start}, nil
	case ')':
		l.pos++
		return spfToken{kind: spfTokenRParen, raw: ")", text: ")", pos: start}, nil
	case ',':
		l.pos++
		return spfToken{kind: spfTokenComma, raw: ",", text: ",", pos: start}, nil
	case ';':
		l.pos++
		return spfToken{kind: spfTokenSemicolon, raw: ";", text: ";", pos: start}, nil
	case '=':
		l.pos++
		return spfToken{kind: spfTokenEquals, raw: "=", text: "=", pos: start}, nil
	}

	if isSPFWordStart(rune(l.data[l.pos])) {
		return l.scanWord()
	}
	if isSPFNumberStart(l.data, l.pos) {
		return l.scanNumber()
	}

	return spfToken{}, fmt.Errorf("%w at byte %d: %q", ErrUnexpectedToken, start, string(l.data[l.pos]))
}

func (l *spfLexer) skipWhitespaceAndComments() error {
	for l.pos < len(l.data) {
		switch l.data[l.pos] {
		case ' ', '\t', '\r', '\n':
			l.pos++
		case '/':
			if l.pos+1 >= len(l.data) || l.data[l.pos+1] != '*' {
				return nil
			}
			end := bytes.Index(l.data[l.pos+2:], []byte("*/"))
			if end == -1 {
				return fmt.Errorf("%w: unterminated comment", ErrInvalidExchangeStructure)
			}
			l.pos += end + 4
		default:
			return nil
		}
	}
	return nil
}

func (l *spfLexer) scanWord() (spfToken, error) {
	start := l.pos
	for l.pos < len(l.data) && isSPFWordPart(rune(l.data[l.pos])) {
		l.pos++
	}
	raw := string(l.data[start:l.pos])
	return spfToken{
		kind: spfTokenWord,
		raw:  raw,
		text: strings.ToUpper(raw),
		pos:  start,
	}, nil
}

func (l *spfLexer) scanString() (spfToken, error) {
	start := l.pos
	l.pos++

	var raw strings.Builder
	for l.pos < len(l.data) {
		switch l.data[l.pos] {
		case '\'':
			if l.pos+1 < len(l.data) && l.data[l.pos+1] == '\'' {
				raw.WriteString("''")
				l.pos += 2
				continue
			}
			l.pos++
			return spfToken{
				kind: spfTokenString,
				raw:  raw.String(),
				text: decodeSTEPString(raw.String()),
				pos:  start,
			}, nil
		default:
			raw.WriteByte(l.data[l.pos])
			l.pos++
		}
	}

	return spfToken{}, fmt.Errorf("%w: unterminated string literal", ErrInvalidExchangeStructure)
}

func (l *spfLexer) scanBinary() (spfToken, error) {
	start := l.pos
	l.pos++

	begin := l.pos
	for l.pos < len(l.data) && l.data[l.pos] != '"' {
		l.pos++
	}
	if l.pos >= len(l.data) {
		return spfToken{}, fmt.Errorf("%w: unterminated binary literal", ErrInvalidExchangeStructure)
	}
	raw := string(l.data[begin:l.pos])
	l.pos++
	return spfToken{kind: spfTokenBinary, raw: raw, text: raw, pos: start}, nil
}

func (l *spfLexer) scanReference() (spfToken, error) {
	start := l.pos
	l.pos++
	begin := l.pos
	for l.pos < len(l.data) && unicode.IsDigit(rune(l.data[l.pos])) {
		l.pos++
	}
	if begin == l.pos {
		return spfToken{}, fmt.Errorf("%w at byte %d: expected digits after '#'", ErrUnexpectedToken, start)
	}
	raw := string(l.data[begin:l.pos])
	return spfToken{kind: spfTokenReference, raw: raw, text: raw, pos: start}, nil
}

func (l *spfLexer) scanEnumeration() (spfToken, error) {
	start := l.pos
	l.pos++
	begin := l.pos
	for l.pos < len(l.data) && isSPFWordPart(rune(l.data[l.pos])) {
		l.pos++
	}
	if begin == l.pos || l.pos >= len(l.data) || l.data[l.pos] != '.' {
		return spfToken{}, fmt.Errorf("%w at byte %d: invalid enumeration literal", ErrUnexpectedToken, start)
	}
	raw := string(l.data[begin:l.pos])
	l.pos++
	return spfToken{kind: spfTokenEnumeration, raw: raw, text: strings.ToUpper(raw), pos: start}, nil
}

func (l *spfLexer) scanNumber() (spfToken, error) {
	start := l.pos
	if l.data[l.pos] == '+' || l.data[l.pos] == '-' {
		l.pos++
	}

	digits := 0
	for l.pos < len(l.data) && unicode.IsDigit(rune(l.data[l.pos])) {
		l.pos++
		digits++
	}

	isReal := false
	if l.pos < len(l.data) && l.data[l.pos] == '.' {
		isReal = true
		l.pos++
		for l.pos < len(l.data) && unicode.IsDigit(rune(l.data[l.pos])) {
			l.pos++
			digits++
		}
	}

	if digits == 0 {
		return spfToken{}, fmt.Errorf("%w at byte %d: invalid numeric literal", ErrUnexpectedToken, start)
	}

	if l.pos < len(l.data) && (l.data[l.pos] == 'E' || l.data[l.pos] == 'e') {
		isReal = true
		l.pos++
		if l.pos < len(l.data) && (l.data[l.pos] == '+' || l.data[l.pos] == '-') {
			l.pos++
		}
		expDigits := 0
		for l.pos < len(l.data) && unicode.IsDigit(rune(l.data[l.pos])) {
			l.pos++
			expDigits++
		}
		if expDigits == 0 {
			return spfToken{}, fmt.Errorf("%w at byte %d: invalid exponent", ErrUnexpectedToken, start)
		}
	}

	raw := string(l.data[start:l.pos])
	kind := spfTokenInteger
	if isReal {
		kind = spfTokenReal
	}
	return spfToken{kind: kind, raw: raw, text: raw, pos: start}, nil
}

func isSPFWordStart(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isSPFWordPart(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

func isSPFNumberStart(data []byte, pos int) bool {
	if pos >= len(data) {
		return false
	}
	switch data[pos] {
	case '+', '-':
		return pos+1 < len(data) && unicode.IsDigit(rune(data[pos+1]))
	default:
		return unicode.IsDigit(rune(data[pos]))
	}
}

type spfParser struct {
	lexer *spfLexer
	token spfToken
}

func newSPFParser(data []byte) (*spfParser, error) {
	parser := &spfParser{lexer: newSPFLexer(data)}
	if err := parser.next(); err != nil {
		return nil, err
	}
	return parser, nil
}

func parseSPF(data []byte) (*Document, error) {
	parser, err := newSPFParser(data)
	if err != nil {
		return nil, err
	}

	if err := parser.expectWord("ISO-10303-21"); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidExchangeStructure, err)
	}
	if err := parser.expect(spfTokenSemicolon); err != nil {
		return nil, err
	}

	doc := &Document{
		Instances: make(map[int64]*Instance),
	}

	for {
		switch {
		case parser.isWord("HEADER"):
			header, err := parser.parseHeader()
			if err != nil {
				return nil, err
			}
			doc.Header = *header
		case parser.isWord("DATA"):
			section, err := parser.parseData()
			if err != nil {
				return nil, err
			}
			doc.Data = append(doc.Data, section)
			for _, instance := range section.Instances {
				if _, exists := doc.Instances[instance.ID]; exists {
					return nil, fmt.Errorf("%w: #%d", ErrDuplicateInstanceID, instance.ID)
				}
				doc.Instances[instance.ID] = instance
				doc.Order = append(doc.Order, instance.ID)
			}
		case parser.isWord("END-ISO-10303-21"):
			if err := parser.next(); err != nil {
				return nil, err
			}
			if err := parser.expect(spfTokenSemicolon); err != nil {
				return nil, err
			}
			if parser.token.kind != spfTokenEOF {
				return nil, fmt.Errorf("%w: trailing tokens after END-ISO-10303-21", ErrInvalidExchangeStructure)
			}
			return doc, nil
		default:
			return nil, fmt.Errorf("%w at byte %d: unexpected section %q", ErrInvalidExchangeStructure, parser.token.pos, parser.token.raw)
		}
	}
}

func (p *spfParser) parseHeader() (*Header, error) {
	if err := p.expectWord("HEADER"); err != nil {
		return nil, err
	}
	if err := p.expect(spfTokenSemicolon); err != nil {
		return nil, err
	}

	header := &Header{}
	for !p.isWord("ENDSEC") {
		record, err := p.parseHeaderRecord()
		if err != nil {
			return nil, err
		}
		header.Records = append(header.Records, *record)
	}
	if err := p.expectWord("ENDSEC"); err != nil {
		return nil, err
	}
	if err := p.expect(spfTokenSemicolon); err != nil {
		return nil, err
	}
	return header, nil
}

func (p *spfParser) parseHeaderRecord() (*HeaderRecord, error) {
	keyword, err := p.expectWordToken()
	if err != nil {
		return nil, err
	}
	if err := p.expect(spfTokenLParen); err != nil {
		return nil, err
	}
	arguments, err := p.parseParameterList()
	if err != nil {
		return nil, err
	}
	if err := p.expect(spfTokenRParen); err != nil {
		return nil, err
	}
	if err := p.expect(spfTokenSemicolon); err != nil {
		return nil, err
	}
	return &HeaderRecord{Keyword: keyword.text, Arguments: arguments}, nil
}

func (p *spfParser) parseData() (*DataSection, error) {
	if err := p.expectWord("DATA"); err != nil {
		return nil, err
	}
	if p.token.kind == spfTokenLParen {
		if err := p.skipParenthesized(); err != nil {
			return nil, err
		}
	}
	if err := p.expect(spfTokenSemicolon); err != nil {
		return nil, err
	}

	section := &DataSection{}
	for !p.isWord("ENDSEC") {
		instance, err := p.parseInstance()
		if err != nil {
			return nil, err
		}
		section.Instances = append(section.Instances, instance)
	}
	if err := p.expectWord("ENDSEC"); err != nil {
		return nil, err
	}
	if err := p.expect(spfTokenSemicolon); err != nil {
		return nil, err
	}
	return section, nil
}

func (p *spfParser) parseInstance() (*Instance, error) {
	if p.token.kind != spfTokenReference {
		return nil, fmt.Errorf("%w at byte %d: expected reference, got %q", ErrUnexpectedToken, p.token.pos, p.token.raw)
	}
	reference := p.token
	if err := p.next(); err != nil {
		return nil, err
	}
	if err := p.expect(spfTokenEquals); err != nil {
		return nil, err
	}
	entity, err := p.expectWordToken()
	if err != nil {
		return nil, err
	}
	if err := p.expect(spfTokenLParen); err != nil {
		return nil, err
	}
	arguments, err := p.parseParameterList()
	if err != nil {
		return nil, err
	}
	if err := p.expect(spfTokenRParen); err != nil {
		return nil, err
	}
	if err := p.expect(spfTokenSemicolon); err != nil {
		return nil, err
	}

	id, err := strconv.ParseInt(reference.raw, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("ifc: parse instance id %q: %w", reference.raw, err)
	}

	return &Instance{
		ID:        id,
		Entity:    entity.text,
		Arguments: arguments,
	}, nil
}

func (p *spfParser) parseParameterList() ([]Parameter, error) {
	if p.token.kind == spfTokenRParen {
		return nil, nil
	}

	parameters := []Parameter{}
	for {
		parameter, err := p.parseParameter()
		if err != nil {
			return nil, err
		}
		parameters = append(parameters, parameter)

		if p.token.kind != spfTokenComma {
			break
		}
		if err := p.next(); err != nil {
			return nil, err
		}
	}
	return parameters, nil
}

func (p *spfParser) parseParameter() (Parameter, error) {
	switch p.token.kind {
	case spfTokenDollar:
		if err := p.next(); err != nil {
			return nil, err
		}
		return OmittedParameter{}, nil
	case spfTokenStar:
		if err := p.next(); err != nil {
			return nil, err
		}
		return DerivedParameter{}, nil
	case spfTokenString:
		token := p.token
		if err := p.next(); err != nil {
			return nil, err
		}
		return StringParameter{Raw: token.raw, Decoded: token.text}, nil
	case spfTokenBinary:
		token := p.token
		if err := p.next(); err != nil {
			return nil, err
		}
		return BinaryParameter{Digits: token.raw}, nil
	case spfTokenInteger:
		token := p.token
		if err := p.next(); err != nil {
			return nil, err
		}
		value, err := strconv.ParseInt(token.raw, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("ifc: parse integer %q: %w", token.raw, err)
		}
		return IntegerParameter{Raw: token.raw, Value: value}, nil
	case spfTokenReal:
		token := p.token
		if err := p.next(); err != nil {
			return nil, err
		}
		value, err := strconv.ParseFloat(token.raw, 64)
		if err != nil {
			return nil, fmt.Errorf("ifc: parse real %q: %w", token.raw, err)
		}
		return RealParameter{Raw: token.raw, Value: value}, nil
	case spfTokenReference:
		token := p.token
		if err := p.next(); err != nil {
			return nil, err
		}
		value, err := strconv.ParseInt(token.raw, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("ifc: parse reference %q: %w", token.raw, err)
		}
		return ReferenceParameter{ID: value}, nil
	case spfTokenEnumeration:
		token := p.token
		if err := p.next(); err != nil {
			return nil, err
		}
		return EnumerationParameter{Value: token.text}, nil
	case spfTokenLParen:
		if err := p.next(); err != nil {
			return nil, err
		}
		elements, err := p.parseParameterList()
		if err != nil {
			return nil, err
		}
		if err := p.expect(spfTokenRParen); err != nil {
			return nil, err
		}
		return AggregateParameter{Elements: elements}, nil
	case spfTokenWord:
		name := p.token.text
		if err := p.next(); err != nil {
			return nil, err
		}
		if err := p.expect(spfTokenLParen); err != nil {
			return nil, err
		}
		arguments, err := p.parseParameterList()
		if err != nil {
			return nil, err
		}
		if err := p.expect(spfTokenRParen); err != nil {
			return nil, err
		}
		return NamedParameter{Name: name, Arguments: arguments}, nil
	default:
		return nil, fmt.Errorf("%w at byte %d: invalid parameter token %q", ErrUnexpectedToken, p.token.pos, p.token.raw)
	}
}

func (p *spfParser) skipParenthesized() error {
	depth := 0
	for {
		switch p.token.kind {
		case spfTokenLParen:
			depth++
		case spfTokenRParen:
			depth--
			if depth == 0 {
				return p.next()
			}
		case spfTokenEOF:
			return fmt.Errorf("%w: unterminated parenthesized clause", ErrInvalidExchangeStructure)
		}
		if err := p.next(); err != nil {
			return err
		}
	}
}

func (p *spfParser) expect(kind spfTokenKind) error {
	if p.token.kind != kind {
		return fmt.Errorf("%w at byte %d: expected %s, got %q", ErrUnexpectedToken, p.token.pos, kind, p.token.raw)
	}
	return p.next()
}

func (p *spfParser) expectWord(expected string) error {
	if !p.isWord(expected) {
		return fmt.Errorf("%w at byte %d: expected %s, got %q", ErrUnexpectedToken, p.token.pos, expected, p.token.raw)
	}
	return p.next()
}

func (p *spfParser) expectWordToken() (spfToken, error) {
	if p.token.kind != spfTokenWord {
		return spfToken{}, fmt.Errorf("%w at byte %d: expected identifier, got %q", ErrUnexpectedToken, p.token.pos, p.token.raw)
	}
	token := p.token
	return token, p.next()
}

func (p *spfParser) isWord(expected string) bool {
	return p.token.kind == spfTokenWord && p.token.text == strings.ToUpper(expected)
}

func (p *spfParser) next() error {
	token, err := p.lexer.next()
	if err != nil {
		return err
	}
	p.token = token
	return nil
}

func decodeSTEPString(raw string) string {
	if raw == "" {
		return ""
	}

	var out strings.Builder
	for i := 0; i < len(raw); {
		switch {
		case i+1 < len(raw) && raw[i] == '\'' && raw[i+1] == '\'':
			out.WriteByte('\'')
			i += 2
		case raw[i] != '\\':
			r, size := utf8.DecodeRuneInString(raw[i:])
			out.WriteRune(r)
			i += size
		case strings.HasPrefix(raw[i:], `\S\`) && i+4 <= len(raw):
			out.WriteRune(rune(raw[i+3]) + 128)
			i += 4
		case strings.HasPrefix(raw[i:], `\X\`) && i+5 <= len(raw):
			value, ok := decodeHexRune(raw[i+3 : i+5])
			if ok {
				out.WriteRune(value)
				i += 5
				continue
			}
			out.WriteString(raw[i : i+5])
			i += 5
		case strings.HasPrefix(raw[i:], `\X2\`):
			segment, size, ok := decodeDelimitedHex(raw[i+4:], 4)
			if ok {
				out.WriteString(segment)
				i += 4 + size
				continue
			}
			out.WriteByte(raw[i])
			i++
		case strings.HasPrefix(raw[i:], `\X4\`):
			segment, size, ok := decodeDelimitedHex(raw[i+4:], 8)
			if ok {
				out.WriteString(segment)
				i += 4 + size
				continue
			}
			out.WriteByte(raw[i])
			i++
		case strings.HasPrefix(raw[i:], `\P`) && i+4 <= len(raw) && raw[i+3] == '\\':
			// IFC guidance explicitly mentions \PA\ page directives. The raw string is
			// preserved, so keeping the directive as a no-op here avoids losing data
			// while still decoding the following escape forms.
			i += 4
		default:
			out.WriteByte(raw[i])
			i++
		}
	}
	return out.String()
}

func decodeHexRune(input string) (rune, bool) {
	buf := make([]byte, hex.DecodedLen(len(input)))
	if _, err := hex.Decode(buf, []byte(input)); err != nil || len(buf) == 0 {
		return 0, false
	}
	return rune(buf[0]), true
}

func decodeDelimitedHex(input string, chunkSize int) (string, int, bool) {
	end := strings.Index(input, `\X0\`)
	if end == -1 {
		return "", 0, false
	}
	payload := input[:end]
	if len(payload)%chunkSize != 0 {
		return "", 0, false
	}

	var out strings.Builder
	for i := 0; i < len(payload); i += chunkSize {
		value, err := strconv.ParseUint(payload[i:i+chunkSize], 16, 32)
		if err != nil {
			return "", 0, false
		}
		out.WriteRune(rune(value))
	}
	return out.String(), end + len(`\X0\`), true
}

func (k spfTokenKind) String() string {
	switch k {
	case spfTokenEOF:
		return "EOF"
	case spfTokenWord:
		return "word"
	case spfTokenString:
		return "string"
	case spfTokenBinary:
		return "binary"
	case spfTokenInteger:
		return "integer"
	case spfTokenReal:
		return "real"
	case spfTokenReference:
		return "reference"
	case spfTokenEnumeration:
		return "enumeration"
	case spfTokenDollar:
		return "$"
	case spfTokenStar:
		return "*"
	case spfTokenLParen:
		return "("
	case spfTokenRParen:
		return ")"
	case spfTokenComma:
		return ","
	case spfTokenSemicolon:
		return ";"
	case spfTokenEquals:
		return "="
	default:
		return "unknown"
	}
}
