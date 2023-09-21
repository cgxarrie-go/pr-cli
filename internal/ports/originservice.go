package ports

import (
	"github.com/cgxarrie-go/prq/internal/utils"
)

type OriginSvc interface {
	GetPRsURL(origin utils.Remote) (url string, err error)
	CreatePRsURL(origin utils.Remote) (url string, err error)
	PRLink(origin utils.Remote, id, text string) (url string, err error)
}