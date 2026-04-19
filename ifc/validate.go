package ifc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	findingCodeSchemaVersionMismatch = "schema.version.mismatch"
	findingCodeEntityUnknown         = "schema.entity.unknown"
	findingCodeEntityAbstract        = "schema.entity.abstract"
	findingCodeAttributeCount        = "schema.attribute.count"
	findingCodeAttributeRequiredOmit = "schema.attribute.required_omitted"
	findingCodeReferenceUnresolved   = "schema.reference.unresolved"
	findingCodeTypeMismatch          = "schema.type.mismatch"
	findingCodeEnumInvalid           = "schema.enum.invalid"
	findingCodeSelectInvalid         = "schema.select.invalid"
	findingCodeAggregateCardinality  = "schema.aggregate.cardinality"
	findingCodeAggregateUnique       = "schema.aggregate.unique"
	findingCodeInverseCardinality    = "schema.inverse.cardinality"
)

type matchState uint8

const (
	matchStateNotApplicable matchState = iota
	matchStateInvalid
	matchStateValid
)

// Severity identifies the validation severity level.
type Severity uint8

const (
	// SeverityError reports a schema validation error.
	SeverityError Severity = iota + 1
)

// Finding reports one schema validation error.
type Finding struct {
	Severity   Severity
	Code       string
	Message    string
	InstanceID int64
	Entity     string
	Attribute  string
}

// Report preserves the validation output for one IFC document.
type Report struct {
	Version  Version
	Findings []Finding
}

// Valid reports whether the validation found no errors.
func (r *Report) Valid() bool {
	return r != nil && len(r.Findings) == 0
}

type validateConfig struct {
	version     Version
	maxFindings int
}

// ValidateOption customizes Validate and Validator.
type ValidateOption func(*validateConfig) error

// WithValidationVersion forces validation to use version instead of doc.Version.
func WithValidationVersion(version Version) ValidateOption {
	return func(cfg *validateConfig) error {
		if version != VersionUnknown && !version.HasSchema() {
			return fmt.Errorf("%w: %s", ErrSchemaUnavailable, version)
		}
		cfg.version = version
		return nil
	}
}

// WithMaxFindings caps the number of returned findings. Zero means unlimited.
func WithMaxFindings(n int) ValidateOption {
	return func(cfg *validateConfig) error {
		if n < 0 {
			return fmt.Errorf("ifc: max findings must be >= 0")
		}
		cfg.maxFindings = n
		return nil
	}
}

// Validator validates IFC documents against the embedded official schema metadata.
type Validator struct {
	cfg    validateConfig
	schema *Schema
}

// NewValidator creates a schema validator.
func NewValidator(opts ...ValidateOption) (*Validator, error) {
	cfg := validateConfig{}
	for _, opt := range opts {
		if err := opt(&cfg); err != nil {
			return nil, err
		}
	}

	var schema *Schema
	if cfg.version != VersionUnknown {
		var err error
		schema, err = SchemaFor(cfg.version)
		if err != nil {
			return nil, err
		}
	}

	return &Validator{
		cfg:    cfg,
		schema: schema,
	}, nil
}

// Validate validates doc against the embedded official schema metadata.
func Validate(doc *Document, opts ...ValidateOption) (*Report, error) {
	validator, err := NewValidator(opts...)
	if err != nil {
		return nil, err
	}
	return validator.Validate(doc)
}

// Validate validates doc against the embedded official schema metadata.
func (v *Validator) Validate(doc *Document) (*Report, error) {
	if doc == nil {
		return nil, ErrNilDocument
	}

	version := v.cfg.version
	if version == VersionUnknown {
		version = doc.Version
	}
	if version == VersionUnknown {
		return nil, ErrValidationVersionRequired
	}
	if !version.HasSchema() {
		return nil, fmt.Errorf("%w: %s", ErrSchemaUnavailable, version)
	}

	schema := v.schema
	if schema == nil || schema.Version != version {
		var err error
		schema, err = SchemaFor(version)
		if err != nil {
			return nil, err
		}
	}

	report := &Report{Version: version}
	state := newValidationState(schema, report, v.cfg.maxFindings)

	if len(doc.RawSchemaIdentifiers) > 0 && !schemaIdentifiersMatch(version, doc.RawSchemaIdentifiers) {
		state.addDocumentFinding(
			findingCodeSchemaVersionMismatch,
			fmt.Sprintf(
				"FILE_SCHEMA %q does not match effective validation schema %s",
				doc.RawSchemaIdentifiers,
				version.SchemaName(),
			),
		)
	}

	for _, id := range doc.Order {
		if state.stopped() {
			break
		}
		instance, ok := doc.Instance(id)
		if !ok {
			continue
		}
		state.validateInstance(doc, instance)
	}

	for _, id := range doc.Order {
		if state.stopped() {
			break
		}
		instance, ok := doc.Instance(id)
		if !ok {
			continue
		}
		state.validateInverse(instance)
	}

	return report, nil
}

type validationState struct {
	schema        *Schema
	report        *Report
	maxFindings   int
	explicitCache map[string][]ExplicitAttribute
	inverseCache  map[string][]InverseAttribute
	derivedCache  map[string]map[string]struct{}
	subtypeCache  map[string]bool
	referenceHits map[referenceIndexKey]map[string]int
}

type referenceIndexKey struct {
	targetID  int64
	attribute string
}

type matchOptions struct {
	collectRefs bool
}

type matchResult struct {
	state   matchState
	refs    []int64
	code    string
	message string
}

func newValidationState(schema *Schema, report *Report, maxFindings int) *validationState {
	return &validationState{
		schema:        schema,
		report:        report,
		maxFindings:   maxFindings,
		explicitCache: map[string][]ExplicitAttribute{},
		inverseCache:  map[string][]InverseAttribute{},
		derivedCache:  map[string]map[string]struct{}{},
		subtypeCache:  map[string]bool{},
		referenceHits: map[referenceIndexKey]map[string]int{},
	}
}

func (s *validationState) stopped() bool {
	return s.maxFindings > 0 && len(s.report.Findings) >= s.maxFindings
}

func (s *validationState) addDocumentFinding(code string, message string) {
	s.addFinding(Finding{
		Severity: SeverityError,
		Code:     code,
		Message:  message,
	})
}

func (s *validationState) addInstanceFinding(code string, message string, instance *Instance, attribute string) {
	finding := Finding{
		Severity: SeverityError,
		Code:     code,
		Message:  message,
	}
	if instance != nil {
		finding.InstanceID = instance.ID
		finding.Entity = instance.Entity
	}
	if attribute != "" {
		finding.Attribute = attribute
	}
	s.addFinding(finding)
}

func (s *validationState) addFinding(f Finding) {
	if s.stopped() {
		return
	}
	s.report.Findings = append(s.report.Findings, f)
}

func (s *validationState) validateInstance(doc *Document, instance *Instance) {
	entity, ok := s.schema.Entity(instance.Entity)
	if !ok {
		s.addInstanceFinding(
			findingCodeEntityUnknown,
			fmt.Sprintf("entity %s is not declared in %s", instance.Entity, s.schema.Name),
			instance,
			"",
		)
		return
	}

	collectRefs := true
	if entity.Abstract {
		s.addInstanceFinding(
			findingCodeEntityAbstract,
			fmt.Sprintf("entity %s is abstract and cannot be instantiated", entity.Name),
			instance,
			"",
		)
		collectRefs = false
	}

	attributes, err := s.allExplicitAttributes(instance.Entity)
	if err != nil {
		s.addInstanceFinding(
			findingCodeEntityUnknown,
			err.Error(),
			instance,
			"",
		)
		return
	}

	if len(attributes) != len(instance.Arguments) {
		s.addInstanceFinding(
			findingCodeAttributeCount,
			fmt.Sprintf(
				"entity %s expects %d explicit attributes, got %d arguments",
				instance.Entity,
				len(attributes),
				len(instance.Arguments),
			),
			instance,
			"",
		)
		return
	}

	for i, attribute := range attributes {
		if s.stopped() {
			return
		}

		argument := instance.Arguments[i]
		if _, ok := argument.(DerivedParameter); ok {
			if s.isDerivedAttribute(instance.Entity, attribute.Name) {
				continue
			}
			s.addInstanceFinding(
				findingCodeTypeMismatch,
				fmt.Sprintf("attribute %s does not allow derived marker '*'", attribute.Name),
				instance,
				attribute.Name,
			)
			continue
		}

		if _, ok := argument.(OmittedParameter); ok {
			if attribute.Optional {
				continue
			}
			s.addInstanceFinding(
				findingCodeAttributeRequiredOmit,
				fmt.Sprintf("attribute %s is required", attribute.Name),
				instance,
				attribute.Name,
			)
			continue
		}

		result := s.matchParameter(doc, argument, attribute.Type, matchOptions{collectRefs: collectRefs})
		if result.state == matchStateValid {
			if collectRefs && len(result.refs) > 0 {
				s.indexReferences(instance, attribute.Name, result.refs)
			}
			continue
		}

		if result.state == matchStateNotApplicable {
			result = s.defaultNotApplicableFailure(argument, attribute.Type)
		}
		s.addInstanceFinding(result.code, result.message, instance, attribute.Name)
	}
}

func (s *validationState) validateInverse(instance *Instance) {
	inverseAttributes, err := s.allInverseAttributes(instance.Entity)
	if err != nil {
		return
	}

	for _, inverse := range inverseAttributes {
		if s.stopped() {
			return
		}

		lower, upper, sourceType := inverseExpectation(inverse.Type)
		count := s.countInverseReferences(instance.ID, inverse.For, sourceType)
		if count >= lower && (upper < 0 || count <= upper) {
			continue
		}

		s.addInstanceFinding(
			findingCodeInverseCardinality,
			fmt.Sprintf(
				"inverse %s expects %s references via %s, got %d",
				inverse.Name,
				describeCardinality(lower, upper),
				inverse.For,
				count,
			),
			instance,
			inverse.Name,
		)
	}
}

func (s *validationState) matchParameter(doc *Document, parameter Parameter, expected TypeExpr, opts matchOptions) matchResult {
	switch expected.Kind {
	case TypeKindNamed:
		return s.matchNamed(doc, parameter, expected.Name, opts)
	case TypeKindBuiltin:
		return s.matchBuiltin(parameter, expected)
	case TypeKindEnumeration:
		return s.matchEnumeration(parameter, expected)
	case TypeKindAggregate:
		return s.matchAggregate(doc, parameter, expected, opts)
	case TypeKindSelect:
		return s.matchSelect(doc, parameter, expected, opts)
	case TypeKindGeneric:
		return validMatch(nil)
	default:
		return invalidMatch(
			findingCodeTypeMismatch,
			fmt.Sprintf("unsupported expected type %s", describeType(expected)),
		)
	}
}

func (s *validationState) matchNamed(doc *Document, parameter Parameter, name string, opts matchOptions) matchResult {
	if entity, ok := s.schema.Entity(name); ok {
		return s.matchEntity(doc, parameter, entity.Name, opts)
	}

	typeDecl, ok := s.schema.Type(name)
	if !ok {
		return invalidMatch(
			findingCodeTypeMismatch,
			fmt.Sprintf("unknown schema type %s", name),
		)
	}

	if named, ok := parameter.(NamedParameter); ok {
		if strings.EqualFold(named.Name, typeDecl.Name) {
			if len(named.Arguments) != 1 {
				return invalidMatch(
					findingCodeTypeMismatch,
					fmt.Sprintf("typed value %s expects exactly one argument", typeDecl.Name),
				)
			}
			result := s.matchParameter(doc, named.Arguments[0], typeDecl.Definition, opts)
			if result.state == matchStateNotApplicable {
				return invalidMatch(
					findingCodeTypeMismatch,
					fmt.Sprintf("typed value %s does not match %s", named.Name, describeType(typeDecl.Definition)),
				)
			}
			return result
		}

		if typeDecl.Definition.Kind != TypeKindSelect {
			return notApplicableMatch()
		}
	}

	return s.matchParameter(doc, parameter, typeDecl.Definition, opts)
}

func (s *validationState) matchEntity(doc *Document, parameter Parameter, expectedEntity string, opts matchOptions) matchResult {
	switch parameter := parameter.(type) {
	case ReferenceParameter:
		target, ok := doc.Instance(parameter.ID)
		if !ok {
			return invalidMatch(
				findingCodeReferenceUnresolved,
				fmt.Sprintf("reference #%d does not exist", parameter.ID),
			)
		}
		if !s.isSubtype(target.Entity, expectedEntity) {
			return notApplicableMatch()
		}
		if !opts.collectRefs {
			return validMatch(nil)
		}
		return validMatch([]int64{parameter.ID})

	case NamedParameter:
		if _, ok := s.schema.Type(parameter.Name); ok {
			return notApplicableMatch()
		}

		entity, ok := s.schema.Entity(parameter.Name)
		if !ok {
			return invalidMatch(
				findingCodeEntityUnknown,
				fmt.Sprintf("inline entity %s is not declared in %s", parameter.Name, s.schema.Name),
			)
		}
		if !s.isSubtype(entity.Name, expectedEntity) {
			return notApplicableMatch()
		}
		if entity.Abstract {
			return invalidMatch(
				findingCodeEntityAbstract,
				fmt.Sprintf("inline entity %s is abstract and cannot be instantiated", entity.Name),
			)
		}

		attributes, err := s.allExplicitAttributes(entity.Name)
		if err != nil {
			return invalidMatch(findingCodeEntityUnknown, err.Error())
		}
		if len(attributes) != len(parameter.Arguments) {
			return invalidMatch(
				findingCodeAttributeCount,
				fmt.Sprintf(
					"inline entity %s expects %d explicit attributes, got %d arguments",
					entity.Name,
					len(attributes),
					len(parameter.Arguments),
				),
			)
		}

		for i, attribute := range attributes {
			argument := parameter.Arguments[i]
			if _, ok := argument.(DerivedParameter); ok {
				if s.isDerivedAttribute(entity.Name, attribute.Name) {
					continue
				}
				return invalidMatch(
					findingCodeTypeMismatch,
					fmt.Sprintf("attribute %s does not allow derived marker '*'", attribute.Name),
				)
			}
			if _, ok := argument.(OmittedParameter); ok {
				if attribute.Optional {
					continue
				}
				return invalidMatch(
					findingCodeAttributeRequiredOmit,
					fmt.Sprintf("attribute %s is required", attribute.Name),
				)
			}

			result := s.matchParameter(doc, argument, attribute.Type, matchOptions{collectRefs: false})
			if result.state == matchStateValid {
				continue
			}
			if result.state == matchStateNotApplicable {
				return s.defaultNotApplicableFailure(argument, attribute.Type)
			}
			return result
		}

		return validMatch(nil)

	default:
		return invalidMatch(
			findingCodeTypeMismatch,
			fmt.Sprintf("expected %s reference, got %T", expectedEntity, parameter),
		)
	}
}

func (s *validationState) matchBuiltin(parameter Parameter, expected TypeExpr) matchResult {
	switch expected.Name {
	case "INTEGER":
		if _, ok := parameter.(IntegerParameter); ok {
			return validMatch(nil)
		}
	case "REAL", "NUMBER":
		switch parameter.(type) {
		case IntegerParameter, RealParameter:
			return validMatch(nil)
		}
	case "STRING":
		value, ok := parameter.(StringParameter)
		if ok && widthMatches(utf8.RuneCountInString(value.Decoded), expected.WidthLower, expected.WidthUpper) {
			return validMatch(nil)
		}
		if ok {
			return invalidMatch(
				findingCodeTypeMismatch,
				fmt.Sprintf("string width for %s must be within [%s:%s]", describeType(expected), expected.WidthLower, expected.WidthUpper),
			)
		}
	case "BINARY":
		value, ok := parameter.(BinaryParameter)
		if ok && widthMatches(len(value.Digits), expected.WidthLower, expected.WidthUpper) {
			return validMatch(nil)
		}
		if ok {
			return invalidMatch(
				findingCodeTypeMismatch,
				fmt.Sprintf("binary width for %s must be within [%s:%s]", describeType(expected), expected.WidthLower, expected.WidthUpper),
			)
		}
	case "BOOLEAN":
		if enum, ok := parameter.(EnumerationParameter); ok {
			if enum.Value == "T" || enum.Value == "F" {
				return validMatch(nil)
			}
		}
	case "LOGICAL":
		if enum, ok := parameter.(EnumerationParameter); ok {
			if enum.Value == "T" || enum.Value == "F" || enum.Value == "U" {
				return validMatch(nil)
			}
		}
	}

	return invalidMatch(
		findingCodeTypeMismatch,
		fmt.Sprintf("value %s does not match %s", describeParameter(parameter), describeType(expected)),
	)
}

func (s *validationState) matchEnumeration(parameter Parameter, expected TypeExpr) matchResult {
	enum, ok := parameter.(EnumerationParameter)
	if !ok {
		return invalidMatch(
			findingCodeTypeMismatch,
			fmt.Sprintf("value %s does not match %s", describeParameter(parameter), describeType(expected)),
		)
	}

	value := strings.ToUpper(enum.Value)
	for _, item := range expected.Items {
		if strings.ToUpper(item) == value {
			return validMatch(nil)
		}
	}

	return invalidMatch(
		findingCodeEnumInvalid,
		fmt.Sprintf("enumeration %s is not valid for %s", enum.String(), describeType(expected)),
	)
}

func (s *validationState) matchAggregate(doc *Document, parameter Parameter, expected TypeExpr, opts matchOptions) matchResult {
	aggregate, ok := parameter.(AggregateParameter)
	if !ok {
		return invalidMatch(
			findingCodeTypeMismatch,
			fmt.Sprintf("value %s does not match %s", describeParameter(parameter), describeType(expected)),
		)
	}

	count := len(aggregate.Elements)
	if !withinCardinality(count, expected.Lower, expected.Upper, true) {
		return invalidMatch(
			findingCodeAggregateCardinality,
			fmt.Sprintf("aggregate %s expects %s elements, got %d", describeType(expected), describeBounds(expected.Lower, expected.Upper), count),
		)
	}

	if strings.EqualFold(expected.Name, "SET") || expected.Unique {
		seen := make(map[string]struct{}, len(aggregate.Elements))
		for _, element := range aggregate.Elements {
			key := element.String()
			if _, ok := seen[key]; ok {
				return invalidMatch(
					findingCodeAggregateUnique,
					fmt.Sprintf("aggregate %s requires unique values, duplicate %s found", describeType(expected), key),
				)
			}
			seen[key] = struct{}{}
		}
	}

	if expected.Element == nil {
		return invalidMatch(
			findingCodeTypeMismatch,
			fmt.Sprintf("aggregate %s has no element type", describeType(expected)),
		)
	}

	var refs []int64
	for _, element := range aggregate.Elements {
		if _, ok := element.(OmittedParameter); ok {
			if expected.Optional {
				continue
			}
			return invalidMatch(
				findingCodeTypeMismatch,
				fmt.Sprintf("aggregate %s does not allow omitted elements", describeType(expected)),
			)
		}
		if _, ok := element.(DerivedParameter); ok {
			return invalidMatch(
				findingCodeTypeMismatch,
				fmt.Sprintf("aggregate %s does not allow derived marker '*'", describeType(expected)),
			)
		}

		result := s.matchParameter(doc, element, *expected.Element, opts)
		if result.state == matchStateValid {
			refs = append(refs, result.refs...)
			continue
		}
		if result.state == matchStateNotApplicable {
			return s.defaultNotApplicableFailure(element, *expected.Element)
		}
		return result
	}

	return validMatch(refs)
}

func (s *validationState) matchSelect(doc *Document, parameter Parameter, expected TypeExpr, opts matchOptions) matchResult {
	switch parameter.(type) {
	case ReferenceParameter, NamedParameter:
	default:
		return invalidMatch(
			findingCodeSelectInvalid,
			fmt.Sprintf("value %s is not a valid typed SELECT value for %s", describeParameter(parameter), describeType(expected)),
		)
	}

	if reference, ok := parameter.(ReferenceParameter); ok {
		if _, exists := doc.Instance(reference.ID); !exists {
			return invalidMatch(
				findingCodeReferenceUnresolved,
				fmt.Sprintf("reference #%d does not exist", reference.ID),
			)
		}
	}

	var bestInvalid *matchResult
	for _, item := range expected.Items {
		result := s.matchNamed(
			doc,
			parameter,
			item,
			opts,
		)
		switch result.state {
		case matchStateValid:
			return result
		case matchStateInvalid:
			if bestInvalid == nil {
				candidate := result
				bestInvalid = &candidate
			}
		}
	}

	if bestInvalid != nil {
		return *bestInvalid
	}

	return invalidMatch(
		findingCodeSelectInvalid,
		fmt.Sprintf("value %s does not match any SELECT item in %s", describeParameter(parameter), describeType(expected)),
	)
}

func (s *validationState) defaultNotApplicableFailure(parameter Parameter, expected TypeExpr) matchResult {
	return invalidMatch(
		findingCodeTypeMismatch,
		fmt.Sprintf("value %s does not match %s", describeParameter(parameter), describeType(expected)),
	)
}

func (s *validationState) indexReferences(instance *Instance, attribute string, refs []int64) {
	if len(refs) == 0 {
		return
	}

	for _, targetID := range refs {
		key := referenceIndexKey{
			targetID:  targetID,
			attribute: attribute,
		}
		counts, ok := s.referenceHits[key]
		if !ok {
			counts = map[string]int{}
			s.referenceHits[key] = counts
		}
		counts[instance.Entity]++
	}
}

func (s *validationState) countInverseReferences(targetID int64, attribute string, sourceType TypeExpr) int {
	counts := s.referenceHits[referenceIndexKey{
		targetID:  targetID,
		attribute: attribute,
	}]
	if len(counts) == 0 {
		return 0
	}

	total := 0
	for sourceEntity, count := range counts {
		if s.entityMatchesType(sourceEntity, sourceType) {
			total += count
		}
	}
	return total
}

func (s *validationState) allExplicitAttributes(entityName string) ([]ExplicitAttribute, error) {
	key := strings.ToUpper(entityName)
	if cached, ok := s.explicitCache[key]; ok {
		return cached, nil
	}

	attributes, err := s.schema.AllExplicitAttributes(entityName)
	if err != nil {
		return nil, err
	}
	s.explicitCache[key] = attributes
	return attributes, nil
}

func (s *validationState) allInverseAttributes(entityName string) ([]InverseAttribute, error) {
	key := strings.ToUpper(entityName)
	if cached, ok := s.inverseCache[key]; ok {
		return cached, nil
	}

	entity, ok := s.schema.Entity(entityName)
	if !ok {
		return nil, fmt.Errorf("%w %s", ErrUnknownEntity, entityName)
	}

	var attributes []InverseAttribute
	for _, supertype := range entity.Supertypes {
		parentAttributes, err := s.allInverseAttributes(supertype)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, parentAttributes...)
	}
	attributes = append(attributes, entity.Inverse...)
	s.inverseCache[key] = attributes
	return attributes, nil
}

func (s *validationState) isDerivedAttribute(entityName string, attributeName string) bool {
	attributes, err := s.derivedAttributes(entityName)
	if err != nil {
		return false
	}
	_, ok := attributes[normalizeAttributeName(attributeName)]
	return ok
}

func (s *validationState) derivedAttributes(entityName string) (map[string]struct{}, error) {
	key := strings.ToUpper(entityName)
	if cached, ok := s.derivedCache[key]; ok {
		return cached, nil
	}

	entity, ok := s.schema.Entity(entityName)
	if !ok {
		return nil, fmt.Errorf("%w %s", ErrUnknownEntity, entityName)
	}

	derived := map[string]struct{}{}
	for _, supertype := range entity.Supertypes {
		parentDerived, err := s.derivedAttributes(supertype)
		if err != nil {
			return nil, err
		}
		for name := range parentDerived {
			derived[name] = struct{}{}
		}
	}
	for _, attribute := range entity.Derived {
		derived[normalizeAttributeName(attribute.Name)] = struct{}{}
	}

	s.derivedCache[key] = derived
	return derived, nil
}

func (s *validationState) isSubtype(actual string, expected string) bool {
	if strings.EqualFold(actual, expected) {
		return true
	}

	key := strings.ToUpper(actual) + "\x00" + strings.ToUpper(expected)
	if cached, ok := s.subtypeCache[key]; ok {
		return cached
	}

	entity, ok := s.schema.Entity(actual)
	if !ok {
		s.subtypeCache[key] = false
		return false
	}

	for _, supertype := range entity.Supertypes {
		if s.isSubtype(supertype, expected) {
			s.subtypeCache[key] = true
			return true
		}
	}

	s.subtypeCache[key] = false
	return false
}

func (s *validationState) entityMatchesType(entityName string, expected TypeExpr) bool {
	switch expected.Kind {
	case TypeKindAggregate:
		if expected.Element == nil {
			return false
		}
		return s.entityMatchesType(entityName, *expected.Element)
	case TypeKindNamed:
		if entity, ok := s.schema.Entity(expected.Name); ok {
			return s.isSubtype(entityName, entity.Name)
		}
		typeDecl, ok := s.schema.Type(expected.Name)
		if !ok {
			return false
		}
		return s.entityMatchesType(entityName, typeDecl.Definition)
	case TypeKindSelect:
		for _, item := range expected.Items {
			if s.entityMatchesType(entityName, TypeExpr{Kind: TypeKindNamed, Name: item, Raw: item}) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func validMatch(refs []int64) matchResult {
	return matchResult{
		state: matchStateValid,
		refs:  refs,
	}
}

func invalidMatch(code string, message string) matchResult {
	return matchResult{
		state:   matchStateInvalid,
		code:    code,
		message: message,
	}
}

func notApplicableMatch() matchResult {
	return matchResult{state: matchStateNotApplicable}
}

func schemaIdentifiersMatch(version Version, rawIdentifiers []string) bool {
	expected := normalizeSchemaIdentifier(version.SchemaName())
	for _, raw := range rawIdentifiers {
		if normalizeSchemaIdentifier(raw) == expected {
			return true
		}
		if VersionFromSchemaIdentifier(raw) == version {
			return true
		}
	}
	return false
}

func normalizeAttributeName(name string) string {
	name = strings.TrimSpace(name)
	if index := strings.LastIndex(name, "."); index != -1 {
		name = name[index+1:]
	}
	if index := strings.LastIndex(name, `\`); index != -1 {
		name = name[index+1:]
	}
	return strings.ToUpper(strings.TrimSpace(name))
}

func widthMatches(length int, lower string, upper string) bool {
	if lower == upper && !strings.EqualFold(strings.TrimSpace(lower), "") {
		lower = ""
	}

	minimum, ok := parseOptionalBound(lower)
	if ok && length < minimum {
		return false
	}
	maximum, ok := parseOptionalBound(upper)
	if ok && length > maximum {
		return false
	}
	return true
}

func parseOptionalBound(raw string) (int, bool) {
	raw = strings.TrimSpace(raw)
	if raw == "" || raw == "?" {
		return 0, false
	}
	value, err := strconv.Atoi(raw)
	if err != nil {
		return 0, false
	}
	return value, true
}

func withinCardinality(count int, lower string, upper string, aggregateDefault bool) bool {
	minimum := 0
	if value, ok := parseOptionalBound(lower); ok {
		minimum = value
	}

	maximum := -1
	if value, ok := parseOptionalBound(upper); ok {
		maximum = value
	} else if !aggregateDefault {
		maximum = 1
	}

	if count < minimum {
		return false
	}
	if maximum >= 0 && count > maximum {
		return false
	}
	return true
}

func inverseExpectation(expr TypeExpr) (int, int, TypeExpr) {
	if expr.Kind == TypeKindAggregate && expr.Element != nil {
		lower := 0
		if value, ok := parseOptionalBound(expr.Lower); ok {
			lower = value
		}
		upper := -1
		if value, ok := parseOptionalBound(expr.Upper); ok {
			upper = value
		}
		return lower, upper, *expr.Element
	}
	return 1, 1, expr
}

func describeCardinality(lower int, upper int) string {
	if upper < 0 {
		return fmt.Sprintf("[%d:?]", lower)
	}
	if lower == upper {
		return fmt.Sprintf("[%d]", lower)
	}
	return fmt.Sprintf("[%d:%d]", lower, upper)
}

func describeBounds(lower string, upper string) string {
	if strings.TrimSpace(lower) == "" && strings.TrimSpace(upper) == "" {
		return "[0:?]"
	}
	if strings.TrimSpace(upper) == "" {
		upper = "?"
	}
	if strings.TrimSpace(lower) == "" {
		lower = "0"
	}
	return fmt.Sprintf("[%s:%s]", lower, upper)
}

func describeType(expr TypeExpr) string {
	if expr.Raw != "" {
		return expr.Raw
	}
	if expr.Name != "" {
		return expr.Name
	}
	return string(expr.Kind)
}

func describeParameter(parameter Parameter) string {
	if parameter == nil {
		return "<nil>"
	}
	return parameter.String()
}
