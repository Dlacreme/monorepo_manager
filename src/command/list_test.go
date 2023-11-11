package command

import (
	"bytes"
	"fmt"
	"monorepo_manager/src/config"
	"monorepo_manager/src/content"
	"monorepo_manager/src/test"
	"testing"
)

func TestList(t *testing.T) {
	t.Run("list all workspaces available", func(t *testing.T) {
		var out bytes.Buffer
		conf := config.Config{
			Name: "test config",
			Workspaces: []config.Workspace{
				{
					Name:   "workspace 1",
					Folder: "workspace_1",
				},
				{
					Name:   "workspace 2",
					Folder: "workspace_2",
				},
			},
		}
		err := list(&out, &conf)
		test.AssertNoError(t, err)
		test.AssertStringEq(t, out.String(), " - workspace 1\n - workspace 2\n")
	})

	t.Run("display a meaningful message if no workspace available", func(t *testing.T) {
		var out bytes.Buffer
		conf := config.Config{
			Name:       "test config",
			Workspaces: []config.Workspace{},
		}
		err := list(&out, &conf)
		test.AssertNoError(t, err)
		test.AssertStringEq(t, out.String(), fmt.Sprintln(content.NoWorkspaceAvailable))
	})
}
