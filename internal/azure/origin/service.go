package origin

import (
	"fmt"
	"strings"

	"github.com/cgxarrie-go/prq/internal/azure/status"
	"github.com/cgxarrie-go/prq/internal/config"
	"github.com/cgxarrie-go/prq/internal/errors"
	"github.com/cgxarrie-go/prq/internal/models"
	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/utils"
	"github.com/muesli/termenv"
)

type service struct{}

// CreatePRsURL implements ports.OriginSvc.
func (s service) CreatePRsURL(o utils.Remote) (url string, err error) {
	if !o.IsAzure() {
		return url, errors.NewErrInvalidRepositoryType(o)
	}

	azOrigin := NewAzureOrigin(o)
	base := s.baseUrl(azOrigin)
	url = fmt.Sprintf("%s/repositories/%s/pullrequests?api-version=7.0"+
		"&supportsIterations=true", base, azOrigin.Repository())

	return
}

// GetPRsURL implements ports.OriginSvc.
func (s service) GetPRsURL(o utils.Remote) (
	url string, err error) {

	if !o.IsAzure() {
		return url, errors.NewErrInvalidRepositoryType(o)
	}

	azOrigin := NewAzureOrigin(o)
	base := s.baseUrl(azOrigin)
	url = fmt.Sprintf("%s/repositories/%s/pullrequests?api-version=7.0"+
		"&searchCriteria.status=%d", base, azOrigin.Repository(),
		status.Active)

	return
}

// PRLink implements ports.OriginSvc.
func (s service) PRLink(o utils.Remote, id, text string) (
	url string, err error) {

	if !o.IsAzure() {
		return url, errors.NewErrInvalidRepositoryType(o)
	}

	azOrigin := NewAzureOrigin(o)
	url = fmt.Sprintf("https://dev.azure.com/%s/%s/_git/%s/pullrequest/%s",
		azOrigin.Organization(), azOrigin.Project(), azOrigin.Repository(), id)
	return termenv.Hyperlink(url, text), nil
}

func NewService() ports.OriginSvc {
	return service{}
}

func (s service) baseUrl(o AzureOrigin) string {
	return fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git",
		o.Organization(),
		o.Project())
}

// NewBranch implements ports.OriginSvc.
func (s service) NewBranch(name string) models.Branch {
	prefix := "refs/heads/"
	bName := strings.TrimPrefix(name, prefix)
	fullName := fmt.Sprintf("%s%s", prefix, bName)
	return models.NewBranch(bName, fullName)
}

// DefaultTargetBranch implements ports.OriginSvc.
func (s service) DefaultTargetBranch() models.Branch {
	return s.NewBranch(config.GetInstance().Azure.DefaultTargetBranch)
}
