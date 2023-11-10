package cli

import (
	"errors"
	"monorepo_manager/src/command"
	"monorepo_manager/src/content"
	"strings"
)

func LineFromArgs(args []string) (*Line, error) {
	if len(args) < 1 {
		return nil, errors.New(content.ExpectArgument)
	}
	command, err := detectCommand(args[0])
	if err != nil {
		return nil, err
	}
	commandArgs, err := parseCommandParams(command, args[1:])
	if err != nil {
		return nil, err
	}

	l := Line{
		Command: command,
		Params:  commandArgs,
	}
	return &l, nil
}

type Line struct {
	Command command.Command
	Params  []string
}

func detectCommand(arg string) (command.Command, error) {
	arg = strings.ToLower(arg)
	switch strings.ToLower(arg) {
	case "init":
		return command.Init, nil
	}
	return "", errors.New(content.CommandNotFound)
}

func parseCommandParams(command command.Command, args []string) ([]string, error) {
	return args, nil
}
