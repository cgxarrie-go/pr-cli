package github

import (
	"github.com/cgxarrie-go/prq/internal/config"
)

func loadConfig() (ghcfg config.GithubConfig) {
	cfg := config.GetInstance()
	cfg.Load()
	return cfg.Github
}
