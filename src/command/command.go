package command

import (
	"errors"
	"monorepo_manager/src/config"
	"monorepo_manager/src/content"
	"os"
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
	out := os.Stdout
	switch cmd {
	case Init:
		return initialize(out)
	case Help:
		return help(out)
	}
	// following commands require a valid config file
	conf, err := config.Load()
	if err != nil {
		return err
	}
	switch cmd {
	case List:
		return list(conf, out)
	case Use:
		return use(conf, out)
	case Workspace:
		return workspace(conf, out)
	}
	return errors.New(content.CommandNotFound)
}
