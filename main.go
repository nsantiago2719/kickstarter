package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var (
	usr, _    = user.Current()
	usrHome   = usr.HomeDir
)

func main() {
	kickFlags := []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Usage:   "Load configuration from `FILE`",
			Aliases: []string{"c"},
			Value:   "starter.toml",
		},
	}
	ks := &cli.App{
		Name:  "kickstarter",
		Usage: "run kickstarter kick to start configuring",
		Commands: []*cli.Command{
			{
				Name:    "kick",
				Flags:   kickFlags,
				Aliases: []string{"k"},
				Usage:   "run and initialize your dotfiles based in your config",
				Action:  makeAction(handlerKick),
			},
		},
	}

	if err := ks.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
	}
}
