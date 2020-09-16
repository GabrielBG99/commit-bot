package logger

import (
	"log"
	"os"
	"time"
)

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

// WriteLog dump log into the specified file
func WriteLog(text string, path string) (int, error) {
	f, err := openFile(path)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer f.Close()

	now := time.Now().Format(time.RFC3339)
	msg := now + ": " + text + "\n"
	return f.WriteString(msg)
}
