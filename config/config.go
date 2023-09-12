package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/cgxarrie-go/prq/cache/providers"
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
	DefaultProvider providers.Provider
	Azure           AzureConfig
}

// AzureConfig configuration for Azure
type AzureConfig struct {
	Organization string
	PAT          string
	Projects     []AzureProjectConfig
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
		return fmt.Errorf("Error loading config file: %s", err.Error())

	}

	err = json.Unmarshal(file, c)
	if err != nil {
		return fmt.Errorf("Error parsing config: %s", err.Error())
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
