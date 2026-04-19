package ifc

import "errors"

var (
	// ErrEmptyInput reports that the parser received no meaningful input.
	ErrEmptyInput = errors.New("ifc: empty input")
	// ErrInvalidExchangeStructure reports that the SPF framing is malformed.
	ErrInvalidExchangeStructure = errors.New("ifc: invalid STEP physical file structure")
	// ErrUnexpectedToken reports a syntax error while reading SPF or EXPRESS input.
	ErrUnexpectedToken = errors.New("ifc: unexpected token")
	// ErrDuplicateInstanceID reports duplicate #id assignments in a DATA section.
	ErrDuplicateInstanceID = errors.New("ifc: duplicate instance id")
	// ErrSchemaUnavailable reports that no embedded official schema exists for the requested version.
	ErrSchemaUnavailable = errors.New("ifc: schema metadata unavailable")
	// ErrNilDocument reports that a nil document receiver was used.
	ErrNilDocument = errors.New("ifc: nil document")
	// ErrNilSchema reports that a nil schema receiver was used.
	ErrNilSchema = errors.New("ifc: nil schema")
	// ErrNilInstance reports that a nil instance receiver or argument was used.
	ErrNilInstance = errors.New("ifc: nil instance")
	// ErrUnknownEntity reports an entity lookup failure in a schema.
	ErrUnknownEntity = errors.New("ifc: unknown entity")
	// ErrCyclicInheritance reports a schema inheritance loop.
	ErrCyclicInheritance = errors.New("ifc: cyclic inheritance")
	// ErrArgumentCountMismatch reports a mismatch between schema attributes and instance arguments.
	ErrArgumentCountMismatch = errors.New("ifc: explicit attribute / argument count mismatch")
)
