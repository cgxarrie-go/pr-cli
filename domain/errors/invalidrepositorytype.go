package errors

import (
	"fmt"

	"github.com/cgxarrie-go/prq/utils"
)

type ErrInvalidRepositoryType struct {
	BaseError
}

func NewErrInvalidRepositoryType(origin utils.Origin) error {
	return ErrInvalidRepositoryType{
		BaseError: BaseError{
			message: fmt.Sprintf("invalid repository type %s", origin),
		},
	}
}