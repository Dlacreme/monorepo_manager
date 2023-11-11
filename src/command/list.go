package command

import (
	"fmt"
	"io"
	"monorepo_manager/src/config"
	"monorepo_manager/src/content"
)

func list(conf *config.Config, w io.Writer) error {
	if len(conf.Workspaces) == 0 {
		fmt.Fprintln(w, content.NoWorkspaceAvailable)
	}
	for _, ws := range conf.Workspaces {
		fmt.Fprintf(w, " - %s\n", ws.Name)
	}
	return nil
}
