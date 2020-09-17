package conf

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Config default config structure
type Config struct {
	RootFolder     string
	IgnoreProjects []string
	LogFilePath    string
}

// GetProjects get git projects from config
func GetProjects(config *Config) []string {
	var projects []string

	filepath.Walk(config.RootFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		if strings.HasSuffix(path, ".git") {
			project := filepath.Dir(path)
			if !isIgnorable(config.IgnoreProjects, project) {
				projects = append(projects, project)
			}
		}
		return nil
	})
	return projects
}

func isIgnorable(projects []string, project string) bool {
	for _, name := range projects {
		if strings.HasSuffix(project, name) {
			return true
		}
	}
	return false
}
