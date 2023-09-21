package origin

import (
	"fmt"

	"github.com/cgxarrie-go/prq/internal/azure/status"
	"github.com/cgxarrie-go/prq/internal/errors"
	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/utils"
	"github.com/muesli/termenv"
)

type service struct{}

// CreatePRsURL implements ports.OriginSvc.
func (s service) CreatePRsURL(o utils.Remote) (url string, err error) {
	if !o.IsGithub() {
		return url, errors.NewErrInvalidRepositoryType(o)
	}

	url = fmt.Sprintf("%s/pulls", s.baseUrl(o))

	return
}

// GetPRsURL implements ports.OriginSvc.
func (s service) GetPRsURL(o utils.Remote) (
	url string, err error) {

	if !o.IsGithub() {
		return url, errors.NewErrInvalidRepositoryType(o)
	}

	qs := "q=is%3Apr+is%3A"

	url = fmt.Sprintf("%s/pulls?%s%s", s.baseUrl(o), qs, status.Active.Name())

	return
}

// PRLink implements ports.OriginSvc.
func (s service) PRLink(o utils.Remote, id, text string) (
	url string, err error) {

	if !o.IsGithub() {
		return url, errors.NewErrInvalidRepositoryType(o)
	}

	ghOrigin := NewGithubOrigin(o)
	url = fmt.Sprintf("https://github.com/%s/%s/pull/%s", 
		ghOrigin.User(), ghOrigin.Repository(), id)
	return termenv.Hyperlink(url, text), nil
}

func NewService() ports.OriginSvc {
	return service{}
}

func (s service) baseUrl(o utils.Remote) string {
	ghOrigin := NewGithubOrigin(o)
	return fmt.Sprintf("https://api.github.com/repos/%s/%s",
	ghOrigin.User(), ghOrigin.Repository())
}
