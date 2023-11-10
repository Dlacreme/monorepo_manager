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
		fmt.Printf("Error:\n%s\n", err)
		return
	}

	err = command.Run(line.Command, line.Params)
	if err != nil {
		fmt.Printf("Failed to execute '%s':\n%s\n", line.Command, err)
	}
}
