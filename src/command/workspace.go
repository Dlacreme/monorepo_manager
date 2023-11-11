package command

import (
	"errors"
	"fmt"
	"io"
	"monorepo_manager/src/config"
	"monorepo_manager/src/content"
	"os"
)

func workspace(w io.Writer, config *config.Config) error {
	used_ws := os.Getenv(ENV_VARNAME)
	if used_ws == "" {
		fmt.Fprintln(w, content.NoWorkspaceInUse)
		return nil
	}
	if !config.IsWorkspaceExisting(used_ws) {
		return errors.New(fmt.Sprintln(used_ws, content.WorkspaceInUseNotFound))
	}
	fmt.Fprintln(w, used_ws)
	return nil
}
