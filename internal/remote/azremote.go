package remote

import (
	"fmt"
	"strings"

	"github.com/cgxarrie-go/prq/internal/azure/status"
	"github.com/cgxarrie-go/prq/internal/config"
	"github.com/cgxarrie-go/prq/internal/models"
	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/remotetype"
)

type azureRemote struct {
	Remote
	organization string
	project      string
	repository   string
	baseUrl      string
}

func newAzureRemote(r string) ports.Remote {

	split := strings.Split(string(r), "/")
	org := split[len(split)-4]
	prj := split[len(split)-3]
	repo := split[len(split)-1]

	defTgtBranch := models.NewBranch(
		config.GetInstance().Azure.DefaultTargetBranch,
		fmt.Sprintf("refs/heads/%s", config.GetInstance().Azure.
			DefaultTargetBranch))

	return azureRemote{
		Remote:       newRemote(r, remotetype.Azure, defTgtBranch),
		organization: org,
		project:      prj,
		repository:   repo,
		baseUrl: fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git",
			org, prj),
	}
}

func (r azureRemote) NewBranch(name string) models.Branch {
	prefix := "refs/heads/"
	bName := strings.TrimPrefix(name, prefix)
	fullName := fmt.Sprintf("%s%s", prefix, bName)
	return models.NewBranch(bName, fullName)
}

func (r azureRemote) GetPRsURL() string {
	return fmt.Sprintf("%s/repositories/%s/pullrequests?api-version=7.0"+
		"&searchCriteria.status=%d", r.baseUrl, r.repository,
		status.Active)
}

func (r azureRemote) CreatePRsURL() string {
	return fmt.Sprintf("%s/repositories/%s/pullrequests?api-version=7.0"+
		"&supportsIterations=true", r.baseUrl, r.repository)
}

func (r azureRemote) Repository() string {
	return r.repository
}

func (r azureRemote) PRLinkURL(id string) string {
	return fmt.Sprintf("https://dev.azure.com/%s/%s/_git/%s/pullrequest/%s",
		r.organization, r.project, r.repository, id)
}
