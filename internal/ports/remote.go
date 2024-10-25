package ports

import (
	"github.com/cgxarrie-go/prq/internal/models"
	"github.com/cgxarrie-go/prq/internal/remotetype"
)

type Remote interface {
	GetPRsURL() string
	CreatePRsURL() string
	PRLinkURL(id string) string
	CodeURL() string
	NewBranch(name string) models.Branch
	DefaultTargetBranch() models.Branch
	Repository() string
	Type() remotetype.RemoteType
	Path() string
}
