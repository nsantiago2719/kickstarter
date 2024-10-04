package main

import (
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
	arg := c.Args().First()
	fmt.Println("kick config: ", arg)

	return nil
}
