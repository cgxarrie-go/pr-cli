package origin

import (
	"strings"

	"github.com/cgxarrie-go/prq/utils"
)

type GithubOrigin struct {
	utils.Origin
	user string
}

func NewGithubOrigin(o utils.Origin) GithubOrigin {

	split := strings.Split(string(o), "/")

	return GithubOrigin{
		Origin:       o,
		user: split[len(split)-2],
	}
}


func (o GithubOrigin)User() string {
	return o.user	
}