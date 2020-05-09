package main

import (
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "csv",
				Aliases: []string{"c"},
				Usage:   "load csv from `FILE`",
				Value:   "problems.csv",
			},

			&cli.BoolFlag{
				Name:    "shuffle",
				Aliases: []string{"s"},
				Usage:   "shuffle questions from file",
				Value:   false,
			},
			&cli.IntFlag{
				Name:        "limit",
				Aliases:     []string{"l"},
				Usage:       "limit `DURATION` per quiz question, in seconds",
				DefaultText: "no limit",
			},
		},
		Action: func(c *cli.Context) error {
			return Quiz(ParseQuizArgs(c))
		},
	}

	app.Run(os.Args)
}

// ParseQuizArgs receives a CLI context and returns a struct our Quiz understands.
func ParseQuizArgs(c *cli.Context) *QuizArgs {
	args := QuizArgs{
		csvPath:   c.String("csv"),
		shuffle:   c.Bool("shuffle"),
		timeLimit: time.Duration(c.Int("limit")) * time.Second,
		timeout:   c.IsSet("limit"),
	}
	return &args
}
