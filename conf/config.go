package conf

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config default config structure
type Config struct {
	RootFolder     string   `yaml:"root" exemple:"/home/<user>/Documents"`
	IgnoreProjects []string `yaml:"ignore" exemple:"[proj-1,proj-2]"`
	LogFilePath    string   `yaml:"log_path" exemple:"/home/<user>/.commit-bot.logs"`
}

// ReadConfig get config from file
func ReadConfig() (*Config, error) {
	config := Config{}
	configPath := GetConfigFilePath()
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if config.RootFolder == "" {
		return nil, errors.New("Root Folder not set")
	}
	return &config, nil
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
