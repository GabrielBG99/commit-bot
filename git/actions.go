package git

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/GabrielBG99/commit-bot/logger"
)

func add(project string) ([]byte, error) {
	return exec.Command("git", "-C", project, "add", ".").Output()
}

func commit(project string) ([]byte, error) {
	add(project)
	msg := time.Now().Format(time.RFC3339) + " - Auto commit, by Commit Bot"
	return exec.Command("git", "-C", project, "commit", "-m", msg).Output()
}

func push(project string) ([]byte, error) {
	return exec.Command("git", "-C", project, "push").Output()
}

func commitAndPush(project string, logFile string) {
	if _, err := commit(project); err != nil {
		logger.WriteLog(fmt.Sprintf("Error: %v", err), logFile)
	} else if _, err := push(project); err != nil {
		logger.WriteLog(fmt.Sprintf("Error: %v", err), logFile)
	}
}
