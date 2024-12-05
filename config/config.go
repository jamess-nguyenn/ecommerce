package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type cacheConfig struct {
	Prefix string `json:"prefix"`
}

type jWTConfig struct {
	ExpirationHours int `json:"expiration_hours"`
}

type loggingConfig struct {
	OutputFile string `json:"output_file"`
}

var (
	Cache   *cacheConfig
	JWT     *jWTConfig
	Logging *loggingConfig
)

func loadConfig(filename string, config any) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open the config file %s: %w", filename, err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return fmt.Errorf("could not decode the config file %s: %w", filename, err)
	}

	return nil
}

func LoadConfigs() error {
	Cache = &cacheConfig{}
	JWT = &jWTConfig{}
	Logging = &loggingConfig{}

	filenameConfig := map[string]any{
		"cache":   Cache,
		"jwt":     JWT,
		"logging": Logging,
	}

	for filename, config := range filenameConfig {
		filenamePath := fmt.Sprintf("config/%s.json", filename)
		if err := loadConfig(filenamePath, config); err != nil {
			return err
		}
	}

	return nil
}
