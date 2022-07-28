package backlog

import (
	"fmt"
	"io"
	"log"
)

type BackLog struct {
	l        Logger
	cfgLevel Level
	prefix   string
	flag     int
}

type Option func(backlog *BackLog)

// WithFileWriter log to file
func WithFileWriter(writer *FileWriter) Option {
	return func(backlog *BackLog) {
		writer.level = backlog.cfgLevel
		backlog.l.SetOutput(writer)
	}
}

// WithConsoleWriter log to console
func WithConsoleWriter(writer io.Writer) Option {
	return func(backlog *BackLog) {
		backlog.l.SetOutput(writer)
	}
}

func NewBackLog(cfgLevel Level, prefix string, flag int, opts ...Option) *BackLog {
	logger := log.New(io.Discard, prefix, flag)
	b := &BackLog{
		l:        logger,
		cfgLevel: cfgLevel,
	}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

func (b *BackLog) checkLevel(level Level) bool {
	if level >= b.cfgLevel {
		return false
	}
	return true
}

// DebugF Log line format: [IWEF]mmdd hh:mm:ss.uuuuuu threadid file:line] msg
func (b *BackLog) DebugF(f string, args ...interface{}) {
	if b.checkLevel(Debug) {
		return
	}
	b.l.Output(2, fmt.Sprintf(Debug.String()+" "+f, args...))
}

func (b *BackLog) InfoF(f string, args ...interface{}) {
	if b.checkLevel(Info) {
		return
	}
	b.l.Output(2, fmt.Sprintf(Info.String()+" "+f, args...))
}

func (b *BackLog) WarnF(f string, args ...interface{}) {
	if b.checkLevel(Warn) {
		return
	}
	b.l.Output(2, fmt.Sprintf(Warn.String()+" "+f, args...))
}

func (b *BackLog) ErrorF(f string, args ...interface{}) {
	if b.checkLevel(Error) {
		return
	}
	b.l.Output(2, fmt.Sprintf(Error.String()+" "+f, args...))
}

func (b *BackLog) FatalF(f string, args ...interface{}) {
	if b.checkLevel(Fatal) {
		return
	}
	b.l.Output(2, fmt.Sprintf(Fatal.String()+" "+f, args...))
}
