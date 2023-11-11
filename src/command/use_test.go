package command

import (
	"bytes"
	"fmt"
	"monorepo_manager/src/config"
	"monorepo_manager/src/content"
	"monorepo_manager/src/test"
	"os"
	"testing"
)

func TestUse(t *testing.T) {
	t.Run("set the workspace", func(t *testing.T) {
		var out bytes.Buffer
		conf := config.Config{
			Name: "test config",
			Workspaces: []config.Workspace{
				{Name: "workspace"},
			},
		}
		err := use(&out, &conf, []string{"workspace"})

		test.AssertNoError(t, err)
		test.AssertStringEq(t, out.String(), fmt.Sprintln(content.NowUsingWorkspace, "workspace"))
		test.AssertStringEq(t, os.Getenv(ENV_VARNAME), "workspace")
	})

	t.Run("returns an error if the workspace doesn't exist", func(t *testing.T) {
		var out bytes.Buffer
		conf := config.Config{
			Name:       "test config",
			Workspaces: []config.Workspace{},
		}
		err := use(&out, &conf, []string{"nonexisting"})
		test.AssertErrorEq(t, err, fmt.Sprintln("nonexisting", content.WorkspaceNotFound))
	})
}
