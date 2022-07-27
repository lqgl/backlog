package main

import "github.com/lqgl/backlog"

func main() {
	fileWriter := backlog.NewFileWriter("log", 0, 0, 0)
	logger := backlog.NewBackLog(backlog.Debug, "", 0, backlog.WithFileWriter(fileWriter))
	logger.DebugF("Test debugF...")
	select {} // block the main thread
}
