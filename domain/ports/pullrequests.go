package ports

import (
	"github.com/cgxarrie-go/prq/domain/models"
)

// PRReader Contract for all services reading Pull requests from providers
type PRReader interface {
	GetPRs(req interface{}) (prs []models.PullRequest, err error)
}

type PRCreator interface {
	Create(req interface{}) (id string, err error)
}
