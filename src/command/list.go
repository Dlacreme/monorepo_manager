package command

import (
	"fmt"
	"io"
	"monorepo_manager/src/config"
	"monorepo_manager/src/content"
)

func list(w io.Writer, conf *config.Config) error {
	if len(conf.Workspaces) == 0 {
		fmt.Fprintln(w, content.NoWorkspaceAvailable)
	}
	for _, ws := range conf.Workspaces {
		fmt.Fprintf(w, " - %s\n", ws.Name)
	}
	return nil
}
