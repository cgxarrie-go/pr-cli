package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

var lock = &sync.Mutex{}

var cfg *Config

// GetInstance returns singleton instance.
func GetInstance() *Config {
	if cfg == nil {
		lock.Lock()
		defer lock.Unlock()
		if cfg == nil {
			cfg = &Config{}
		}
	}

	return cfg
}

// Config main configuration for CLI
type Config struct {
	Azure AzureConfig
}

// AzureConfig configuration for Azure
type AzureConfig struct {
	CompanyName string
	PAT         string
	Projects    []AzureProjectConfig
}

// AzureProjectConfig configuration for Azure Projects
type AzureProjectConfig struct {
	ID            string
	RepositoryIDs []string
}

// Load read config file into config struct
func (c *Config) Load() (err error) {
	file, err := ioutil.ReadFile(".prcliconfig")
	if err != nil {
		return fmt.Errorf("Error loading config file: %s", err.Error())

	}

	err = json.Unmarshal(file, c)
	if err != nil {
		return fmt.Errorf("Error parsing config: %s", err.Error())
	}

	return nil
}
