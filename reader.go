package main

import (
	"os"
	"path"

	"github.com/pelletier/go-toml/v2"
)

type TomlConfig struct {
	Config map[string]Config `toml:"configs"`
}

type Config struct {
	Source      string `toml:"source"`
	Destination string `toml:"destination"`
	Description string `toml:"description"`
}

func readToml(path string) (*TomlConfig, error) {
	cfg := TomlConfig{}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func readPath(p string) string {
	cleanPath := path.Clean(p)

	// check if tilde is the first character
	// substitute it with current users HomeDir
	if s := cleanPath[:1]; s == "~" {
		return path.Join(usrHome, cleanPath[1:])
	}
	return cleanPath
}
