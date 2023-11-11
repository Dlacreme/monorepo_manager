package command

import (
	"errors"
	"fmt"
	"io"
	"monorepo_manager/src/config"
	"monorepo_manager/src/content"
	"os"
	"strings"
)

func exec(w io.Writer, conf *config.Config, params []string) error {
	ws, params := findWorkspace(conf, params)
	if ws == nil {
		return errors.New(content.NoWorkspaceInUse)
	}
	fmt.Printf("EXEC > [%s] => %+v\n", ws.Name, params)
	return nil
}

func findWorkspace(conf *config.Config, params []string) (*config.Workspace, []string) {
	var ws *config.Workspace = nil
	if !strings.Contains(params[0], "$") {
		wsName := os.Getenv(conf.EnvVarName)
		if wsName == "" {
			return nil, params
		}
		ws = conf.GetWorkspace(os.Getenv(conf.EnvVarName))
	}

	splitted := strings.Split(params[0], "$")
	if len(splitted) == 1 {
		ws = conf.GetRoot()
		params[0] = splitted[0]
	} else {
		ws = conf.GetWorkspace(splitted[0])
		params[0] = splitted[1]
	}
	return ws, params
}
