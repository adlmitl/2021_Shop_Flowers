package logg

import (
	"log"
	"os"
	"shopflowers/src/pkg/util"
)

const pathToLoggFile = "./shopflowers.log"

type Logg struct {
	logInfo  *log.Logger
	logError *log.Logger
}

func NewLogg() *Logg {
	file, err := logFile()
	if err != nil {
		log.Fatalf("\033[1;31m[E] Error log file not found: %s\033[0m", err.Error())
	}

	logInfo := log.New(file, " INFO ---", log.Lmsgprefix|log.LstdFlags)
	logError := log.New(file, "ERROR --- ", log.Lmsgprefix|log.LstdFlags)

	return &Logg{
		logInfo:  logInfo,
		logError: logError,
	}
}

func logFile() (*os.File, error) {
	file, err := os.OpenFile(pathToLoggFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("\033[1;31m[E] Error while opening logger file: %s\033[0m", err.Error())
	}

	return file, err
}

// LogInfo - Log info.
func (l *Logg) LogInfo(message string) {
	log.Printf("\033[1;34m[I] %s\033[0m", message)
	l.logInfo.Printf(" %s - %s", util.FileWithFuncAndLineNum(), message)
}

// LogInfoWithArgs - Log infof with args.
func (l *Logg) LogInfoWithArgs(message string, args ...interface{}) {
	log.Printf("\033[1;34m[I] %s: %v\033[0m", message, args)
	l.logInfo.Printf(" %s - %v", util.FileWithFuncAndLineNum(), args)
}

// LogError - Log Errorf with args.
func (l *Logg) LogError(message string, args ...interface{}) {
	log.Printf("%s \033[1;31m[E] %v: %s\033[0m", util.FileWithLineNum(), message, args)
	l.logError.Printf("%s - %s: %v", util.FileWithFuncAndLineNum(), message, args)
	os.Exit(1)
}
