package azure

import (
	"fmt"
	"strings"

	"github.com/cgxarrie-go/prq/utils"
)

func getRepoParams() (organization string, projectName string, repoName string,
	err error) {

	origin, err := utils.GitCurrentOriginURL()
	if err != nil {
		return "", "", "",
			fmt.Errorf("getting origin url for current location: %w", err)
	}

	split := strings.Split(origin, "/")
	repoName = split[len(split)-1]
	projectName = split[len(split)-3]
	organization = split[len(split)-4]

	return
}
