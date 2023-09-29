package pipes

import (
	"errors"
	"fmt"
)

// ErrInvalidOperationGivenPartLoopState signals an error
var ErrInvalidOperationGivenPartLoopState = errors.New("invalid operation given state of loop")

// CreateErrUnmarshalLogLine creates an error
func CreateErrUnmarshalLogLine(marshalized []byte, originalErr error) error {
	return fmt.Errorf("unmarshal log line [%s]: %w", string(marshalized), originalErr)
}
