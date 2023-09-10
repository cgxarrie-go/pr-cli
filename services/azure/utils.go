package azure

import (
	"fmt"
	"strings"

	"github.com/cgxarrie-go/prq/utils"
)

func getRepoParamsFromOrigin(origin utils.Origin) (
	organization string, projectName string, repoName string, err error) {

	split := strings.Split(string(origin), "/")
	repoName = split[len(split)-1]
	projectName = split[len(split)-3]
	organization = split[len(split)-4]

	return
}

func getRepoParams() (organization string, projectName string, repoName string,
	err error) {

	origin, err := utils.GitCurrentOrigin()
	if err != nil {
		return "", "", "",
			fmt.Errorf("getting origin url for current location: %w", err)
	}

	return getRepoParamsFromOrigin(origin)
}
