package ports

import (
	"github.com/cgxarrie-go/prq/domain/models"
)

// PRReader Contract for all services reading Pull requests from providers
type PRReader interface {
	GetPRs(req ListPRRequest) (prs []models.PullRequest, err error)
}

type PRCreator interface {
	Create(req CreatePRRequest) (pr models.CreatedPullRequest, err error)
}
