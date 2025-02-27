package config

import (
	"encoding/json"
	"os"
)

func Read() (*Config, error) {
	cfg := &Config{}

	configPath, err := getConfigFilePath()

	configFile, err := os.Open(configPath)
	if err != nil {
		return cfg, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.Current_user_name = username
	configPath, err := getConfigFilePath()

	json_content, err := json.Marshal(cfg)

	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, json_content, 0644)
	if err != nil {
		return err
	}

	return nil
}
