package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "csv",
				Aliases: []string{"c"},
				Usage:   "Load csv from `FILE`",
				Value:   "problems.csv",
			},

			&cli.BoolFlag{
				Name:    "shuffle",
				Aliases: []string{"s"},
				Usage:   "Shuffle questions from file",
				Value:   false,
			},
		},
		Action: Quiz,
	}

	app.Run(os.Args)
}
