package command

import (
	"monorepo_manager/src/content"
	"monorepo_manager/src/test"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("can execute 'init'", func(t *testing.T) {
		err := Run("init", []string{})
		test.AssertNoError(t, err)
	})

	t.Run("returns an error if the command is unknown", func(t *testing.T) {
		err := Run("unknowncommand", []string{})
		test.AssertErrorEq(t, err, content.CommandNotFound)
	})

}
