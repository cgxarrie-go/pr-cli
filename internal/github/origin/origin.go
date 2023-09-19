package origin

import (
	"strings"

	"github.com/cgxarrie-go/prq/internal/utils"
)

type GithubOrigin struct {
	utils.Origin
	user       string
	repository string
}

func NewGithubOrigin(o utils.Origin) GithubOrigin {

	split := strings.Split(string(o), "/")

	return GithubOrigin{
		Origin:     o,
		user:       split[len(split)-2],
		repository: split[len(split)-1],
	}
}

func (o GithubOrigin) User() string {
	return o.user
}

func (o GithubOrigin) Repository() string {
	return o.repository
}
