package command

import (
	"errors"
	"monorepo_manager/src/constant"
)

type Command string

// list available commands
const (
	Init Command = "init"
)

const ()

func Run(cmd Command, params []string) error {
	switch cmd {
	case Init:
		return initialize()
	}
	return errors.New(constant.CommandNotFound)
}
