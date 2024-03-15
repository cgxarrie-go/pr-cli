package errors

import (
	"fmt"
)

type ErrInvalidRepositoryType struct {
	BaseError
}

func NewUnknownRepositoryType(r string) error {
	return ErrInvalidRepositoryType{
		BaseError: BaseError{
			message: fmt.Sprintf("remote type not supported: %s", r),
		},
	}
}
