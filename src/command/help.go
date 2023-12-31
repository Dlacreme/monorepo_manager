package command

import (
	"fmt"
	"io"
)

func help(w io.Writer) error {
	fmt.Fprintf(w, `
NAME:
	mm - monorepo manager

DESCRIPTION:
	The current workspace is defined using the ENV variable MM_WORKSPACE.

COMMANDS:

	Usage:

	- list, l
		list all workspaces available

	- use, u $workspace
		define the workspace to use

	- workspace, w
		get the workspace currently used

	Config:

	- init, i
		create a default monorepo.toml file

	Other:

	- help, h
		display this helper

`)
	return nil
}
