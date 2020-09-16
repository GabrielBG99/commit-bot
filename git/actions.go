package git

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/GabrielBG99/commit-bot/logger"
)

func Commit(project string) ([]byte, error) {
	msg := time.Now().Format(time.RFC3339) + " - Auto commit, by Commit Bot"
	command := fmt.Sprintf(`
		cd %v && \
		git add . && \
		git commit -m "%v"
	`, project, msg)
	return exec.Command("bash", "-c", command).Output()
}

func Push(project string) ([]byte, error) {
	command := fmt.Sprintf(`
		cd %v && \
		git push
	`, project)
	return exec.Command("bash", "-c", command).Output()
}

func CommitAndPush(project string, logFile string) {
	if _, err := Commit(project); err != nil {
		logger.WriteLog(fmt.Sprintf("Error: %v", err), logFile)
	} else if _, err := Push(project); err != nil {
		logger.WriteLog(fmt.Sprintf("Error: %v", err), logFile)
	}
}
