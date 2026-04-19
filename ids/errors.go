package ids

import "errors"

var (
	// ErrEmptyInput reports that the parser received no meaningful input.
	ErrEmptyInput = errors.New("ids: empty input")
	// ErrInvalidDocument reports that the XML content cannot be interpreted as an IDS document.
	ErrInvalidDocument = errors.New("ids: invalid IDS document")
	// ErrInvalidOption reports that Parse received an invalid option value.
	ErrInvalidOption = errors.New("ids: invalid parse option")
	// ErrUnsupportedVersion reports that a caller forced an unsupported IDS version.
	ErrUnsupportedVersion = errors.New("ids: unsupported IDS version")
)
