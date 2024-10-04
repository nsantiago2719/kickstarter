package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	ks := &cli.App{
		Name:  "kickstarter",
		Usage: "run kickstarter kick to start configuring",
		Commands: []*cli.Command{
			{
				Name:    "kick",
				Aliases: []string{"k"},
				Usage:   "run and initialize your dotfiles based in your config",
				Action:  makeAction(handlerKick),
			},
		},
	}

	if err := ks.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
