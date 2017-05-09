package factories

import (
	"os/exec"

	"github.com/medbridge/mocking/mocks"
)

//ICommandFactory conforms to exec.Command signature. Used for mocking exec.Command
type ICommandFactory interface {
	Command(string, ...string) mocks.ICommand
}

//CommandTestFactory implements the ICommandFactory and mocks exec.Command
type CommandTestFactory struct {
	Commands [][]string
}

//Command mocks exec.Command
func (e *CommandTestFactory) Command(name string, arg ...string) mocks.ICommand {
	cmdString := []string{name}
	cmdString = append(cmdString, arg...)
	e.Commands = append(e.Commands, cmdString)
	cmd := mocks.Command{Command: cmdString}
	return &cmd
}

//CommandFactory wrapper for exec.Command. Used to provide a mocking interface
type CommandFactory struct{}

//Command wrapper for real exec.Command method
func (e *CommandFactory) Command(name string, arg ...string) mocks.ICommand {
	return exec.Command(name, arg...)
}
