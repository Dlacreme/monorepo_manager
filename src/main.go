package main

import (
	"fmt"
	"monorepo_manager/src/cli"
	"monorepo_manager/src/command"
	"os"
)

func main() {
	line, err := cli.LineFromArgs(os.Args[1:])
	if err != nil {
		fmt.Printf("(exit) Error:\n%s\n", err)
		return
	}
	command.Run(line.Command, line.Params)
}
