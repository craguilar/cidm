package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
)

var logger *log.Logger

type ContextHook struct{}

func (hook ContextHook) Levels() []log.Level {
	return log.AllLevels
}

func (hook ContextHook) Fire(entry *log.Entry) error {
	if pc, file, line, ok := runtime.Caller(8); ok {
		funcName := runtime.FuncForPC(pc).Name()
		entry.Data["file"] = path.Base(file)
		entry.Data["func"] = path.Base(funcName)
		entry.Data["line"] = line
	}
	return nil
}

// Logger gets the root logger for this application
func Logger() *log.Logger {
	if logger == nil {
		logger = log.New()
		//logger.AddHook(ContextHook{})
		logger.Formatter = &log.JSONFormatter{}
		logger.Out = os.Stdout

	}
	return logger
}
