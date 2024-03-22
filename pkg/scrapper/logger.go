package scrapper

import (
	"log"
	"os"
)

var logger *log.Logger

func loggerInit(sessionId string) {
	file, err := openLogFile("logs/log_" + sessionId + ".txt")
	logger = log.New(log.Writer(), "[LOG] ", log.Ltime)
	if err != nil {
		log.Fatal(err)
	}
	logger.SetOutput(file)
}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
