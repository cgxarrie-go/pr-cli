package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var lock = &sync.Mutex{}

var cfg *Config

// GetInstance returns singleton instance.
func GetInstance() *Config {
	if cfg == nil {
		lock.Lock()
		defer lock.Unlock()
		cfg = &Config{}
	}

	return cfg
}

var (
	configFileName string = "prqcfg.json"
)

// Config main configuration for CLI
type Config struct {
	Azure AzureConfig
}

// AzureConfig configuration for Azure
type AzureConfig struct {
	PAT string
}

// NewConfig creates a new instance of Config
func NewConfig() Config {
	return Config{
		Azure: AzureConfig{},
	}
}

// AzureProjectConfig configuration for Azure Projects
type AzureProjectConfig struct {
	ID            string
	RepositoryIDs []string
}

// Load read config file into config struct
func (c *Config) Load() (err error) {
	fileName, err := c.fileName()
	if err != nil {
		return err
	}

	if _, err := os.Stat(fileName); err != nil {
		return fmt.Errorf("configuration file not found\n" +
			"Use prq config to add configuration")
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("loading config file: %w", err)

	}

	err = json.Unmarshal(file, c)
	if err != nil {
		return fmt.Errorf("error parsing config: %w", err)
	}

	return nil
}

// Save saves the config file
func (c *Config) Save() (err error) {
	fileName, err := c.fileName()
	if err != nil {
		return err
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(b))

	return err
}

func (c *Config) fileName() (folder string, err error) {
	ex, err := os.Executable()
	if err != nil {
		return folder, err
	}
	exPath := filepath.Dir(ex)
	fileName := fmt.Sprintf("%s/%s", exPath, configFileName)
	return fileName, nil
}
