package src

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AllowedIps      []string `json:"allowed_ips"`
	CheckAllowedIps bool     `json:"check_allowed_ips"`
	Port            int      `json:"port"`
	Token           string   `json:"token"`
}

const configFile = "config.json"

func LoadConfig() (*Config, error) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("Config file not found, creating default config...")
		defaultConfig := Config{
			AllowedIps:      []string{"0.0.0.0", "127.0.0.1", "127.0.0.2"},
			CheckAllowedIps: true,
			Port:            25555,
			Token:           "default",
		}

		file, err := os.Create(configFile)

		if err != nil {
			return nil, err
		}

		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(defaultConfig); err != nil {
			return nil, err
		}

		return &defaultConfig, nil
	}

	file, err := os.Open(configFile)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
