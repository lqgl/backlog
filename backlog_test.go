package backlog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"testing"
)

func TestWithConsoleWriter(t *testing.T) {
	log := NewBackLog(Debug, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile, WithConsoleWriter(os.Stdout))
	log.DebugF("111111111")
	log.InfoF("222222222")
}

func TestWithFileWriter(t *testing.T) {
	fileWriter := NewFileWriter("log", 0, 0, 0)
	log := NewBackLog(Debug, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile, WithFileWriter(fileWriter))
	log.DebugF("111111111")
	select {}
}

func TestCallerSkip(t *testing.T) {
	_, file, line, ok := runtime.Caller(0)
	fmt.Println(file, line, ok)
}
