package main

import (
	"os/user"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadToml(t *testing.T) {
	val, _ := readToml("./starter.toml")
	assert.IsType(t, &TomlConfig{}, val)
}

func TestReadPath(t *testing.T) {
	usr, _ := user.Current()
	homePath := readPath("~/view")
	expectedPath := path.Join(usr.HomeDir, "view")

	assert.Equal(t, homePath, expectedPath)
}
