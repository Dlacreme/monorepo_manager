package cli

import (
	"errors"
	"monorepo_manager/src/command"
	"monorepo_manager/src/global"
	"strings"
)

// represents a command line parsed & validated
type Line struct {
	Command command.Command
	Params  []string
}

// parse args into a valid Line or return an error if the input is invalid
func LineFromArgs(args []string) (*Line, error) {
	if len(args) < 1 {
		return nil, errors.New(global.ExpectArgument)
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

func detectCommand(arg string) (command.Command, error) {
	switch strings.ToLower(arg) {
	case "init":
		return command.Init, nil
	}
	return "", errors.New(global.InvalidCommandLine)
}

func parseCommandParams(command command.Command, args []string) ([]string, error) {
	return args, nil
}
