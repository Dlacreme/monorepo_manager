package cli

import (
	"fmt"
	"monorepo_manager/src/command"
	"monorepo_manager/src/content"
	"monorepo_manager/src/test"
	"testing"
)

func TestLineFromArgs(t *testing.T) {

	t.Run("ignore case", func(t *testing.T) {
		subject, err := LineFromArgs([]string{"InIT"})
		test.AssertNoError(t, err)
		test.AssertStringEq(t, string(subject.Command), string(command.Init))
	})

	t.Run("fail gracefully if no parameters available", func(t *testing.T) {
		_, err := LineFromArgs([]string{})
		test.AssertErrorEq(t, err, fmt.Sprint("mm ", content.ExpectArgument))
	})

	t.Run("can parse 'init'", func(t *testing.T) {
		subject, err := LineFromArgs([]string{"init"})
		test.AssertNoError(t, err)
		test.AssertStringEq(t, string(subject.Command), string(command.Init))

		subject, err = LineFromArgs([]string{"i"})
		test.AssertNoError(t, err)
		test.AssertStringEq(t, string(subject.Command), string(command.Init))
	})

	t.Run("can parse 'help'", func(t *testing.T) {
		subject, err := LineFromArgs([]string{"help"})
		test.AssertNoError(t, err)
		test.AssertStringEq(t, string(subject.Command), string(command.Help))

		subject, err = LineFromArgs([]string{"h"})
		test.AssertNoError(t, err)
		test.AssertStringEq(t, string(subject.Command), string(command.Help))
	})

	t.Run("can parse 'list'", func(t *testing.T) {
		subject, err := LineFromArgs([]string{"list"})
		test.AssertNoError(t, err)
		test.AssertStringEq(t, string(subject.Command), string(command.List))

		subject, err = LineFromArgs([]string{"l"})
		test.AssertNoError(t, err)
		test.AssertStringEq(t, string(subject.Command), string(command.List))
	})

	t.Run("can parse 'workspace'", func(t *testing.T) {
		subject, err := LineFromArgs([]string{"workspace"})
		test.AssertNoError(t, err)
		test.AssertStringEq(t, string(subject.Command), string(command.Workspace))

		subject, err = LineFromArgs([]string{"w"})
		test.AssertNoError(t, err)
		test.AssertStringEq(t, string(subject.Command), string(command.Workspace))
	})

	t.Run("can parse 'use'", func(t *testing.T) {
		subject, err := LineFromArgs([]string{"use", "hello"})
		test.AssertNoError(t, err)
		test.AssertStringEq(t, string(subject.Command), string(command.Use))
		test.AssertNumEq(t, len(subject.Params), 1)
		test.AssertStringEq(t, subject.Params[0], "hello")

		subject, err = LineFromArgs([]string{"u", "hello"})
		test.AssertNoError(t, err)
		test.AssertStringEq(t, string(subject.Command), string(command.Use))
	})

	t.Run("return an error if use has no params", func(t *testing.T) {
		_, err := LineFromArgs([]string{"use"})
		test.AssertErrorEq(t, err, fmt.Sprint(command.Use, " ", content.ExpectArgument))
	})

}
