package ports

import (
	"github.com/cgxarrie-go/prq/internal/remote"
)

type ListPRRequest interface {
	Origins() remote.Remotes
}
