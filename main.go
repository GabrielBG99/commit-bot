package main

import (
	"log"

	"github.com/GabrielBG99/commit-bot/git"

	"github.com/GabrielBG99/commit-bot/conf"
)

func main() {
	config, err := conf.ReadConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	projects := conf.GetProjects(config)
	for _, path := range projects {
		git.CommitAndPush(path, config.LogFilePath)
	}
}
