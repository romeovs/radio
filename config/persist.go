package config

import (
	"encoding/json"
	"os"
)

// Open the config file and return the contents.
func Open(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return &Config{}, nil
	}

	cfg := new(Config)
	err = json.NewDecoder(file).Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// Save the config to the config file.
func Save(path string, config *Config) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	return json.NewEncoder(file).Encode(config)
}
