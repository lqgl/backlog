package backlog

import (
	"fmt"
	"io"
	"strings"
)

type Level int

const (
	Debug Level = iota + 1
	Info
	Warn
	Error
	Fatal
)

type Logger interface {
	Output(callerSkip int, s string) error
	SetOutput(w io.Writer)
}

func (l *Level) Get() interface{} {
	return *l
}

func (l *Level) Set(s string) error {
	logLevel, err := ParseLogLevel(s)
	if err != nil {
		return err
	}
	*l = logLevel
	return nil
}

func (l Level) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	}
	return "invalid log level"
}

func ParseLogLevel(levelStr string) (Level, error) {
	switch strings.ToLower(levelStr) {
	case "debug":
		return Debug, nil
	case "info":
		return Info, nil
	case "warn":
		return Warn, nil
	case "error":
		return Error, nil
	case "fatal":
		return Fatal, nil
	}
	return 0, fmt.Errorf("invalid log level '%s' (debug, info, warn, error, fatal", levelStr)
}
