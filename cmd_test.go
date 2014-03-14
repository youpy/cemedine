package cmd_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/youpy/go-cmd"
	"testing"
)

func goodCommand(command *cmd.Command, args ...string) (err error) {
	return
}

func badCommand(command *cmd.Command, args ...string) (err error) {
	err = errors.New("error")
	return
}

func commandWithFlag(command *cmd.Command, args ...string) (err error) {
	value := command.Flag.String("foo", "", "usage of foo")
	command.Flag.Parse(args)

	if *value != "bar" {
		err = errors.New("error")
	}

	return
}

func TestExecuteCommand(t *testing.T) {
	cmd.ResetForTesting()

	good := cmd.NewCommand("good", "this is usage for good command", "good command", goodCommand)
	bad := cmd.NewCommand("bad", "this is usage for bad command", "bad command", badCommand)

	cmd.Register(good)
	cmd.Register(bad)

	err := cmd.Exec([]string{"good", "foo"})

	assert.Nil(t, err)

	err = cmd.Exec([]string{"bad", "foo"})

	assert.NotNil(t, err)

	err = cmd.Exec([]string{"xxxx", "foo"})

	assert.NotNil(t, err)
}

func TestUsage(t *testing.T) {
	cmd.ResetForTesting()

	good := cmd.NewCommand("good", "this is usage for good command", "good command", goodCommand)
	bad := cmd.NewCommand("bad", "this is usage for bad command", "bad command", badCommand)

	cmd.Register(good)
	cmd.Register(bad)

	usage, err := cmd.Usage()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, `
Commands:
    good       good command
    bad        bad command
`, usage)
}

func TestCommandWithFlag(t *testing.T) {
	cmd.ResetForTesting()

	command := cmd.NewCommand("test", "this is usage for test command", "test command", commandWithFlag)
	cmd.Register(command)

	err := cmd.Exec([]string{"test", "-foo", "bar"})

	assert.Nil(t, err)
}
