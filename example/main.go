package main

import (
	"github.com/lqgl/backlog"
	"log"
)

func main() {
	fileWriter := backlog.NewFileWriter("log", 0, 0, 0)
	logger := backlog.NewBackLog(backlog.Debug, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile, backlog.WithFileWriter(fileWriter))
	logger.DebugF("Test debugF...")
	select {} // block the main thread
}
