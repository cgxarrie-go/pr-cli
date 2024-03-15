package ports

import (
	"github.com/cgxarrie-go/prq/internal/models"
)

// PRReader Contract for all services reading Pull requests from providers
type PRReader interface {
	GetPRs(req ListPRRequest) (prs []models.PullRequest, err error)
}

type PRCreator interface {
	Run(req CreatePRSvcRequest) (pr CreatePRSvcResponse, err error)
}
