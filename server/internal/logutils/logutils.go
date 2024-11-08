package logutils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func SetupLogging() (*os.File, error) {
	logDir := os.Getenv("LOG_FILE_DIR")
	if logDir == "" {
		logDir = "./logs"
	}

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("error creating log directory: %v", err)
	}

	currentTime := time.Now().Format("2006-01-02")
	logFilePath := fmt.Sprintf("%s/%s.log", logDir, currentTime)

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	log.SetOutput(file)
	return file, nil
}
