package ports

import (
	"github.com/cgxarrie-go/prq/internal/models"
	"github.com/cgxarrie-go/prq/internal/utils"
)

type RemoteService interface {
	GetPRsURL(origin utils.Remote) (url string, err error)
	CreatePRsURL(origin utils.Remote) (url string, err error)
	PRLink(origin utils.Remote, id, text string) (url string, err error)
	NewBranch(name string) models.Branch
	DefaultTargetBranch() models.Branch
}
