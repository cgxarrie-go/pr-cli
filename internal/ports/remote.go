package ports

import (
	"github.com/cgxarrie-go/prq/internal/models"
)

type Remote interface {
	GetPRsURL() string
	CreatePRsURL() string
	PRLinkURL(id string) string
	NewBranch(name string) models.Branch
	DefaultTargetBranch() models.Branch
	Repository() string
}
