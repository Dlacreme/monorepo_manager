package command

import (
	"errors"
	"fmt"
	"io"
	"monorepo_manager/src/config"
	"monorepo_manager/src/content"
	"os"
)

func use(w io.Writer, config *config.Config, params []string) error {
	target_ws := params[0]
	if !config.IsWorkspaceExisting(target_ws) {
		return errors.New(fmt.Sprintln(target_ws, content.WorkspaceNotFound))
	}
	err := os.Setenv(config.EnvVarName, target_ws)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, content.NowUsingWorkspace, target_ws)
	return nil
}
