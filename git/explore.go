package git

import (
	"github.com/GabrielBG99/commit-bot/conf"
	"github.com/urfave/cli/v2"
)

// CommitAndPush wrapper function
func CommitAndPush(c *cli.Context) error {
	config := &conf.Config{
		RootFolder:     c.String("base-dir"),
		IgnoreProjects: c.StringSlice("ignore"),
		LogFilePath:    c.String("log-file"),
	}
	projects := conf.GetProjects(config)
	for _, path := range projects {
		commitAndPush(path, config.LogFilePath)
	}
	return nil
}
