package errors

import "fmt"

// ErrInvalidRequestType occurs the the type request sent to a method
// is not of the expected type
type ErrInvalidRequestType struct {
	message string
}

// NewErrInvalidRequestType returns a new ErrInvalidRequestType error
func NewErrInvalidRequestType(
	expectedObject interface{},
	receivedObject interface{}) ErrInvalidRequestType {
	return ErrInvalidRequestType{
		message: fmt.Sprintf("Invalid request type. Expected %T but got %T",
			expectedObject, receivedObject),
	}
}

// Error returns the ErrInvalidRequestType error message
func (e ErrInvalidRequestType) Error() string {
	return e.message
}

// Is returns if error is of ErrInvalidRequestType type
func (e ErrInvalidRequestType) Is(err error) bool {
	_, ok := err.(ErrInvalidRequestType)
	return ok
}
