package cli

import (
	"monorepo_manager/src/command"
	"monorepo_manager/src/global"
	"monorepo_manager/src/test"
	"testing"
)

func TestLineFromArgs(t *testing.T) {

	t.Run("can parse 'init'", func(t *testing.T) {
		subject, err := LineFromArgs([]string{"init"})
		test.AssertNoError(t, err)
		test.AssertStringEq(t, string(subject.Command), string(command.Init))
	})

	t.Run("ignore case", func(t *testing.T) {
		subject, err := LineFromArgs([]string{"InIT"})
		test.AssertNoError(t, err)
		test.AssertStringEq(t, string(subject.Command), string(command.Init))
	})

	t.Run("fail gracefully if no parameters available", func(t *testing.T) {
		_, err := LineFromArgs([]string{})
		test.AssertErrorEq(t, err, global.ExpectArgument)
	})

}
