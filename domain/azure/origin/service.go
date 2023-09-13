package origin

import (
	"fmt"
	"strings"

	"github.com/cgxarrie-go/prq/domain/errors"
	"github.com/cgxarrie-go/prq/domain/ports"
	"github.com/cgxarrie-go/prq/utils"
	"github.com/muesli/termenv"
)

type service struct{}

// Organizaion implements ports.OriginSvc.
func (s service) Organizaion(origin utils.Origin) (organization string, 
	err error) {
	if !origin.IsAzure() {
		return organization, errors.NewErrInvalidRepositoryType(origin)
	}

	org, _, _ := s.getRepoParamsFromOrigin(origin)
	return org, nil
}

// CreatePRsURL implements ports.OriginSvc.
func (s service) CreatePRsURL(origin utils.Origin) (url string, err error) {
	if !origin.IsAzure() {
		return url, errors.NewErrInvalidRepositoryType(origin)
	}

	org, proj, repo := s.getRepoParamsFromOrigin(origin)
	base := s.baseUrl(org, proj)
	url = fmt.Sprintf("%s/repositories/%s/pullrequests?api-version=7.0"+
		"&supportsIterations=true", base, repo)

	return
}

// GetPRsURL implements ports.OriginSvc.
func (s service) GetPRsURL(origin utils.Origin, status ports.PRStatus) (
	url string, err error) {

	if !origin.IsAzure() {
		return url, errors.NewErrInvalidRepositoryType(origin)
	}

	org, proj, repo := s.getRepoParamsFromOrigin(origin)
	base := s.baseUrl(org, proj)
	url = fmt.Sprintf("%s/repositories/%s/pullrequests?api-version=7.0"+
		"&searchCriteria.status=%d", base, repo, status)

	return
}

// PRLink implements ports.OriginSvc.
func (s service) PRLink(origin utils.Origin, id, text string) (
	url string, err error) {

	if !origin.IsAzure() {
		return url, errors.NewErrInvalidRepositoryType(origin)
	}

	org, proj, repo := s.getRepoParamsFromOrigin(origin)
	base := s.baseUrl(org, proj)
	url = fmt.Sprintf("%s/%s/pullrequest/%s", base, repo, id)
	return termenv.Hyperlink(url, text), nil
}

func NewService() ports.OriginSvc {
	return service{}
}

func (s service) baseUrl(org, proj string) string {
	return fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git", org, proj)
}

func (s service) getRepoParamsFromOrigin(origin utils.Origin) (
	organization string, projectName string, repoName string) {

	split := strings.Split(string(origin), "/")
	repoName = split[len(split)-1]
	projectName = split[len(split)-3]
	organization = split[len(split)-4]

	return
}
