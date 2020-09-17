package main

import (
	"log"
	"os"
	"os/user"
	"path"

	"github.com/GabrielBG99/commit-bot/git"

	"github.com/urfave/cli/v2"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	app := &cli.App{
		Name:  "commit-bot",
		Usage: "Commit and Push projects automatically",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "base-dir",
				Aliases: []string{"b"},
				Value:   path.Join(user.HomeDir, "Documents"),
				Usage:   "Set the base dir to search recussively for git projects",
			},
			&cli.StringSliceFlag{
				Name:    "ignore",
				Aliases: []string{"i"},
				Usage:   "Projects to ignore while searching for git repositories",
			},
			&cli.StringFlag{
				Name:  "log-file",
				Value: path.Join(user.HomeDir, ".commit_bot.log"),
				Usage: "Path to log file",
			},
		},
		Action: cli.ActionFunc(git.CommitAndPush),
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
