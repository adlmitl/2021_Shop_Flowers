package logg

import (
	"log"
	"os"
	"shopflowers/src/pkg/util"
)

const pathToLoggFile = "./shopflowers.log"

// CommonLogger - Общая структура лог файла.
type CommonLogger struct {
	logInfo          *log.Logger
	logError         *log.Logger
	logErrorResponse *log.Logger
}

// NewCommonLogger - Конструктор CommonLogger.
func NewCommonLogger() *CommonLogger {
	file, err := logFile()
	if err != nil {
		log.Fatalf("\033[1;31m[E] Error log file not found: %s\033[0m", err.Error())
	}
	logInfo := log.New(file, " INFO ---", log.Lmsgprefix|log.LstdFlags)
	logError := log.New(file, "ERROR ---", log.Lmsgprefix|log.LstdFlags)
	logErrorResponse := log.New(file, "ERROR_RESPONSE ---", log.Lmsgprefix|log.LstdFlags)

	return &CommonLogger{
		logInfo:          logInfo,
		logError:         logError,
		logErrorResponse: logErrorResponse,
	}
}

// Info - Логирование информации.
func (l *CommonLogger) Info(msg string) {
	log.Printf("\033[1;34m[I] %s \033[0m", msg)
	l.logInfo.Printf("%s - %s", util.FileWithFuncAndLineNum(), msg)
}

// InfoWithArg - Логирование информации с аргументом.
func (l *CommonLogger) InfoWithArg(msg string, arg interface{}) {
	log.Printf("\033[1;34m[I] %s: %s\033[0m", msg, arg)
	l.logInfo.Printf("%s - %s: %s", util.FileWithFuncAndLineNum(), msg, arg)
}

// Error - Логирование ошибок.
func (l *CommonLogger) Error(path, err string) {
	log.Printf("%s \033[1;31m[E] %s: %s\033[0m", util.FileWithLineNum(), path, err)
	l.logError.Fatalf("%s - %s: %s", util.FileWithFuncAndLineNum(), path, err)
}

// ErrorWithArg - Логирование ошибок с аргументом.
func (l *CommonLogger) ErrorWithArg(msg string, arg interface{}) {
	log.Printf("%s \033[1;31m[E] %s: %v\033[0m", util.FileWithLineNum(), msg, arg)
	l.logError.Fatalf("%s - %s: %s", util.FileWithFuncAndLineNum(), msg, arg)
}

// ErrorResponse - Логирование статус ответа веб-сервера.
func (l *CommonLogger) ErrorResponse(msg string, statusCode int, err string) {
	log.Printf("%s \033[1;31m[E_RESPONSE] %s: %d: %s\033[0m", util.FileWithLineNum(), msg, statusCode, err)
	l.logError.Fatalf("%s - %s: %d: %s", util.FileWithFuncAndLineNum(), msg, statusCode, err)
}

// logFile - Лог файл.
func logFile() (*os.File, error) {
	file, err := os.OpenFile(pathToLoggFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("\033[1;31m[E] Error while opening logger file: %s\033[0m", err.Error())
	}

	return file, err
}
