package mocks

import (
	"io"
)

//ICommand conforms to exec.Cmd struct. Used for mocking exec.Cmd
type ICommand interface {
	CombinedOutput() ([]byte, error)
	StdinPipe() (io.WriteCloser, error)
}

//Command mocks exec.Cmd struct
type Command struct {
	Command []string
}

func (e *Command) CombinedOutput() ([]byte, error) {
	return nil, nil
}

func (e *Command) StdinPipe() (io.WriteCloser, error) {
	var stub StdinPipe
	return &stub, nil
}

//StdinPipe mocks Cmd.StdinPipe method
type StdinPipe struct {
	Closed bool
	Bytes  []byte
}

func (t *StdinPipe) Close() error {
	t.Closed = true
	return nil
}

func (t *StdinPipe) Write(p []byte) (int, error) {
	t.Bytes = p
	return 1, nil
}
