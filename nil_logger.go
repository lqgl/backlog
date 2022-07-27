package backlog

import "io"

type NilLogger struct{}

func (n NilLogger) Output(callerSkip int, s string) error {
	return nil
}

func (n *NilLogger) SetOutput(w io.Writer) {

}
