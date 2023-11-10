package command

import "fmt"

func help() error {
	fmt.Printf(`
Usage:

  - list, l
  	list all workspaces available

  - use, u $workspace
  	define the workspace to use

  - repository
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
