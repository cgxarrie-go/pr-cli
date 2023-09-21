package origin

import (
	"strings"

	"github.com/cgxarrie-go/prq/internal/utils"
)

type AzureOrigin struct {
	utils.Remote
	organization string
	project      string
	repository   string
}

func NewAzureOrigin(o utils.Remote) AzureOrigin {

	split := strings.Split(string(o), "/")

	return AzureOrigin{
		Remote:       o,
		organization: split[len(split)-4],
		project:      split[len(split)-3],
		repository:   split[len(split)-1],
	}
}

func (o AzureOrigin) Organization() string {
	return o.organization
}

func (o AzureOrigin) Project() string {
	return o.project
}

func (o AzureOrigin) Repository() string {
	return o.repository
}
