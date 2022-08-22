package main

import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelInfo LogLevel = iota
	LogLevelWarning
	LogLevelError
)

func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelInfo, LogLevelWarning, LogLevelError:
		return true
	default:
		return false
	}
}

type LogExtended struct {
	*log.Logger
	logLevel LogLevel
}

func (lg *LogExtended) SetLogLevel(lgLevel LogLevel) {
	if !lgLevel.IsValid() {
		return
	}
	lg.logLevel = lgLevel
}

func (lg *LogExtended) Infoln(msg string) {
	lg.print(LogLevelInfo, "INFO", msg)
}

func (lg *LogExtended) Warnln(msg string) {
	lg.print(LogLevelWarning, "WARNING", msg)
}

func (lg *LogExtended) Errorln(msg string) {
	lg.print(LogLevelError, "ERROR", msg)
}

func (lg *LogExtended) print(lvl LogLevel, prefix, msg string) {
	if lvl < lg.logLevel {
		return
	}

	var outputColor string
	switch lvl {
	case LogLevelError:
		outputColor = "\033[31m"
	case LogLevelWarning:
		outputColor = "\033[33m"
	case LogLevelInfo:
		outputColor = "\033[32m"
	}

	lg.Logger.Println(outputColor + prefix + " " + msg + "\033[0m")
}

func NewLogExtended() LogExtended {
	return LogExtended{
		logLevel: LogLevelError,
		Logger:   log.New(os.Stderr, "\033[36m[LOGGER] ", log.LstdFlags),
	}
}

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
}
