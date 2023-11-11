package command

import (
	"errors"
	"monorepo_manager/src/content"
)

type Command string

// list available commands
const (
	Help      Command = "help"
	Init      Command = "init"
	List      Command = "list"
	Use       Command = "use"
	Workspace Command = "workspace"
)

const ()

func Run(cmd Command, params []string) error {
	switch cmd {
	case Init:
		return initialize()

	case Help:
		return help()
	}
	return errors.New(content.CommandNotFound)
}
