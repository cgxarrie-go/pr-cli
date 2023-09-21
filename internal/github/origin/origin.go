package origin

import (
	"strings"

	"github.com/cgxarrie-go/prq/internal/utils"
)

type GithubOrigin struct {
	utils.Remote
	user       string
	repository string
}

func NewGithubOrigin(o utils.Remote) GithubOrigin {

	split := strings.Split(string(o), "/")

	repo := split[len(split)-1]
	repo = strings.Split(repo,".")[0]

	return GithubOrigin{
		Remote:     o,
		user:       split[len(split)-2],
		repository: repo,
	}
}

func (o GithubOrigin) User() string {
	return o.user
}

func (o GithubOrigin) Repository() string {
	return o.repository
}
