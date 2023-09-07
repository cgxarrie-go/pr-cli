package azure

import (
	"fmt"

	"github.com/cgxarrie-go/prq/config"
	"github.com/muesli/termenv"
)

func loadConfig() (azcfg config.AzureConfig, err error) {
	cfg := config.GetInstance()
	cfg.Load()
	if err != nil {
		return azcfg, err
	}

	return cfg.Azure, nil
}

func getPRLink(text, organization, project, repository, id string) string {
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_git/%s/pullrequest/%s",
		organization, project, repository, id)
	return termenv.Hyperlink(url, text)
}
