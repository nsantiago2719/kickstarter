package main

import (
	"os"
	"os/user"
	"path"

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
		usr, _ := user.Current()
		return path.Join(usr.HomeDir, cleanPath[1:])
	}
	return cleanPath
}
