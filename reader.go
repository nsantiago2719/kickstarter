package main

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type TomlConfig struct {
	Config map[string]Config `toml:"config"`
}

type Config struct {
	Source      string
	Destination string
	Description string
}

func readToml(f string) (*TomlConfig, error) {
	cfg := TomlConfig{}
	file, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
