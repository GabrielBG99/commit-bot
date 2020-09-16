package conf

import (
	"log"
	"os/user"
	"path"
)

// GetConfigFilePath returns the current user's commit-bot config file
func GetConfigFilePath() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	path := path.Join(user.HomeDir, ".config", "commit_bot", "config.yml")
	return path
}
