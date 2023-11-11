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

func TestWorkspace(t *testing.T) {
	t.Run("return the workspace in use", func(t *testing.T) {
		var out bytes.Buffer
		conf := config.Config{
			EnvVarName: "MM_TEST_WORKSPACE",
			Name:       "test config",
			Workspaces: []config.Workspace{
				{Name: "workspace"},
			},
		}
		err := workspace(&out, &conf)
		os.Setenv(conf.EnvVarName, "workspace")

		test.AssertNoError(t, err)
		test.AssertStringEq(t, out.String(), "workspace\n")
	})

	t.Run("display a meaningful message if no workspace in use", func(t *testing.T) {
		var out bytes.Buffer
		conf := config.Config{
			EnvVarName: "MM_TEST_WORKSPACE",
			Name:       "test config",
			Workspaces: []config.Workspace{
				{Name: "workspace"},
			},
		}
		os.Setenv(conf.EnvVarName, "")
		err := workspace(&out, &conf)

		test.AssertNoError(t, err)
		test.AssertStringEq(t, out.String(), fmt.Sprintln(content.NoWorkspaceInUse))
	})

	t.Run("returns an error if the workspace doesn't exist", func(t *testing.T) {
		var out bytes.Buffer
		conf := config.Config{
			EnvVarName: "MM_TEST_WORKSPACE",
			Name:       "test config",
			Workspaces: []config.Workspace{
				{Name: "workspace"},
			},
		}
		os.Setenv(conf.EnvVarName, "non-existant-workspace")
		err := workspace(&out, &conf)

		s := fmt.Sprintln("non-existant-workspace", content.WorkspaceInUseNotFound)
		test.AssertErrorEq(t, err, s)
	})
}
