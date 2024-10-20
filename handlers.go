package main

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
)

func makeAction(f cli.ActionFunc) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		if err := f(ctx); err != nil {
			fmt.Printf("error: %v", err)
		}
		return nil
	}
}

func handlerKick(c *cli.Context) error {
	if c.NArg() > 0 {
		return errors.New("Command kick is not expecting arguments")
	}

	cfg, err := readToml(c.String("config"))
	if err != nil {
		return err
	}

	// Iterate thru all config declared
	for k, v := range cfg.Config {
		fmt.Println("Configuring ", k)
		if err := copyConfig(v.Source, v.Destination); err != nil {
			return err
		}
	}
	return err
}
