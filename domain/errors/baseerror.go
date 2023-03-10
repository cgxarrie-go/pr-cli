package errors

import "fmt"

// BaseError bas error for all errors to display in the same way
type BaseError struct {
	message string
}

// ErrConfigFileNotFound error for config file not found
type ErrConfigFileNotFound struct {
	BaseError
}

// Error returns the ErrConfigFileNotFound error message
func (e BaseError) Error() string {
	return e.message
}

// Print prints the formatted error message
func Print(e error) {
	fmt.Printf("ERROR!! : %s\n", e.Error())
}
