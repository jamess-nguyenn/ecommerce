package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type CacheConfig struct {
	Prefix string `json:"prefix"`
}

type JWTConfig struct {
	ExpirationHours int `json:"expiration_hours"`
}

type LoggingConfig struct {
	OutputFile string `json:"output_file"`
}

var (
	Cache   *CacheConfig
	JWT     *JWTConfig
	Logging *LoggingConfig
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

func LoadConfigs() {
	Cache = &CacheConfig{}
	JWT = &JWTConfig{}
	Logging = &LoggingConfig{}

	filenameConfig := map[string]any{
		"cache":   Cache,
		"jwt":     JWT,
		"logging": Logging,
	}

	for filename, config := range filenameConfig {
		filenamePath := fmt.Sprintf("config/%s.json", filename)
		_ = loadConfig(filenamePath, config)
	}
}
