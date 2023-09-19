package github

import (
	"github.com/cgxarrie-go/prq/internal/config"
)

func loadConfig() (ghcfg config.GithubConfig, err error) {
	cfg := config.GetInstance()
	cfg.Load()
	if err != nil {
		return ghcfg, err
	}

	return cfg.Github, nil
}