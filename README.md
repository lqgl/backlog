# backlog
A custom switchable logger </br>
This a learning project. </br>
Reference from [spoor](https://github.com/phuhao00/spoor)

## Installation
If you'd like to install this locally, you can go get it.
````sh
# fetches this repo into $GOPATH
go get -u github.com/lqgl/backlog@latest
````

## 💡 Simple Usage
You can import `backlog` using:
````go
import "github.com/lqgl/backlog"
````
Then use one of the helpers below:

### fileWriter
Output logs to file
````go
fileWriter := backlog.NewFileWriter("log", 0, 0, 0)
logger := backlog.NewBackLog(backlog.Debug, "", 0, backlog.WithFileWriter(fileWriter))
logger.DebugF("Test DebugF...")
select {} // block the main thread
````

### consoleWriter
Output logs to console
````go
log := backlog.NewBackLog(Debug, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile, WithConsoleWriter(os.Stdout))
log.DebugF("Test DebugF...")
select {}
````
