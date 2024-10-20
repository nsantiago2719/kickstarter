package main

import (
	"os"
	"testing"

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
}
