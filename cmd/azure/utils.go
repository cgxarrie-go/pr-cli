package azure

import (
	"github.com/cgxarrie-go/prq/internal/config"
)

func loadConfig() (azcfg config.AzureConfig) {
	cfg := config.GetInstance()
	cfg.Load()
	return cfg.Azure
}
