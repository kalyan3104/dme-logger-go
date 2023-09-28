package pipes

import "errors"

// ErrInvalidOperationGivenPartLoopState signals an error
var ErrInvalidOperationGivenPartLoopState = errors.New("invalid operation given state of loop")
