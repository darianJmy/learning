package log

import (
	"log"
	"os"
	"path/filepath"
)

const (
	consoleLog = "gin-practise.console"
)

var (
	ConsoleInfo    *log.Logger
	ConsoleWarning *log.Logger
	ConsoleError   *log.Logger
)

func ConsoleLogger(logDir string) error {
	logFile, err := os.OpenFile(filepath.Join(logDir, consoleLog), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return err
	}

	ConsoleInfo = log.New(logFile, "[INFO]\t", log.LstdFlags)
	ConsoleWarning = log.New(logFile, "[WARNING]\t", log.LstdFlags)
	ConsoleError = log.New(logFile, "[ERROR]\t", log.LstdFlags)
	return nil
}
