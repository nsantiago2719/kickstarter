package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestHandlerKick(t *testing.T) {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Usage:   "Load configuration from `FILE`",
			Aliases: []string{"c"},
			Value:   "starter.toml",
		},
	}
	app := &cli.App{
		Name:  "kickstarter",
		Usage: "run kickstarter kick to start configuring",
		Commands: []*cli.Command{
			{
				Name:    "kick",
				Flags:   flags,
				Aliases: []string{"k"},
				Usage:   "run and initialize your dotfiles based in your config",
				Action:  makeAction(handlerKick),
			},
		},
	}

	os.Args = []string{"kickstarter", "kick"}

	_ = app.Run(os.Args)

	filePaths := [4]string{
		"/tmp/starter.toml",
		"/tmp/config/sub-1/sub-1-1a",
		"/tmp/config/sub-2/sub-2-1/sub-2-1a",
		"/tmp/config/test",
	}

	for _, v := range filePaths {
		_, err := os.Stat(v)
		fmt.Println("Checking file: ", v)
		fileExists := !os.IsNotExist(err)
		assert.Equal(t, true, fileExists)
	}
}
