package origin

import (
	"strings"

	"github.com/cgxarrie-go/prq/utils"
)

type AzureOrigin struct {
	utils.Origin
	organization string
	project string
	repository string
}

func NewAzureOrigin(o utils.Origin) AzureOrigin {

	split := strings.Split(string(o), "/")

	return AzureOrigin{
		Origin:       o,
		organization: split[len(split)-4],
		project:      split[len(split)-3],
		repository:   split[len(split)-1],
	}
}


func (o AzureOrigin)Organizaion() string {
	return o.organization
}

func (o AzureOrigin)Project() string {
	return o.project
}

func (o AzureOrigin)Repository() string {
	return o.repository
}