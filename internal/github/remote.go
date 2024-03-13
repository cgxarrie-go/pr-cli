package github

import (
	"fmt"
	"strings"

	"github.com/cgxarrie-go/prq/internal/azure/status"
	"github.com/cgxarrie-go/prq/internal/config"
	"github.com/cgxarrie-go/prq/internal/models"
	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/remote"
	"github.com/cgxarrie-go/prq/internal/remotetype"
)

type githubRemote struct {
	remote.Remote
	user       string
	repository string
	baseUrl    string
}

func NewRemote(r string) ports.Remote {

	split := strings.Split(string(r), "/")

	repo := split[len(split)-1]
	repo = strings.Split(repo, ".")[0]
	usr := split[len(split)-2]

	defTgtBranch := models.NewBranch(
		config.GetInstance().Github.DefaultTargetBranch,
		fmt.Sprintf("refs/heads/%s", config.GetInstance().Github.DefaultTargetBranch))

	return githubRemote{
		Remote:     remote.NewRemote(r, remotetype.Github, defTgtBranch),
		user:       usr,
		repository: repo,
		baseUrl:    fmt.Sprintf("https://api.github.com/repos/%s/%s", usr, repo),
	}
}

func (r githubRemote) NewBranch(name string) models.Branch {
	prefix := "refs/heads/"
	bName := strings.TrimPrefix(name, prefix)
	fullName := fmt.Sprintf("%s%s", prefix, bName)
	return models.NewBranch(bName, fullName)
}

func (r githubRemote) GetPRsURL() string {
	qs := "q=is%3Apr+is%3A"
	return fmt.Sprintf("%s/pulls?%s%s", r.baseUrl, qs, status.Active.Name())
}

func (r githubRemote) CreatePRsURL() string {
	return fmt.Sprintf("%s/pulls", r.baseUrl)
}

func (r githubRemote) PRLinkURL(id string) string {

	return fmt.Sprintf("https://github.com/%s/%s/pull/%s",
		r.user, r.repository, id)

	// TODO: sample to get the link
	// return termenv.Hyperlink(url, text)
}
