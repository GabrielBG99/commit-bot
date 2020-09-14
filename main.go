package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type config struct {
	RootFolder     string   `yaml:"root"`
	IgnoreProjects []string `yaml:"ignore"`
	LogFilePath    string   `yaml:"log_path"`
}

func readConfig() (*config, error) {
	configPath := "config.yml"
	config := config{}
	data, err := ioutil.ReadFile(configPath)

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if config.RootFolder == "" {
		return nil, errors.New("Root Folder not set")
	}
	return &config, nil
}

func isIgnorable(projects []string, project string) bool {
	for _, name := range projects {
		if strings.HasSuffix(project, name) {
			return true
		}
	}
	return false
}

func getProjects(config *config) []string {
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

func currentTime() string {
	t := time.Now()
	return fmt.Sprintf(
		"%d-%02d-%02d %02d:%02d:%02d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
	)
}

func openFile(path string) (*os.File, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		f.Close()
	}
	return os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
}

func writeLog(text string, path string) (int, error) {
	f, err := openFile(path)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer f.Close()

	msg := currentTime() + ": " + text + "\n"
	return f.WriteString(msg)
}

func commit(project string) ([]byte, error) {
	msg := currentTime() + " - Auto commit, by Commit Bot"
	command := fmt.Sprintf(`
		cd %v && \
		git add . && \
		git commit -m "%v"
	`, project, msg)
	return exec.Command("bash", "-c", command).Output()
}

func push(project string) ([]byte, error) {
	command := fmt.Sprintf(`
		cd %v && \
		git push
	`, project)
	return exec.Command("bash", "-c", command).Output()
}

func commitAndPush(project string, logFile string) {
	if _, err := commit(project); err != nil {
		writeLog(fmt.Sprintf("Error: %v", err), logFile)
	} else if _, err := push(project); err != nil {
		writeLog(fmt.Sprintf("Error: %v", err), logFile)
	}
}

func main() {
	config, err := readConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	projects := getProjects(config)
	for _, path := range projects {
		commitAndPush(path, config.LogFilePath)
	}
}
