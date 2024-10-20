package main

import (
	"os/user"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadPath(t *testing.T) {
	usr, _ := user.Current()
	homePath := readPath("~/view")
	expectedPath := filepath.Join(usr.HomeDir, "view")

	assert.Equal(t, homePath, expectedPath)
}
