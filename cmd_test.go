package cmd

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func goodCommand(command *Command, args ...string) (err error) {
	return
}

func badCommand(command *Command, args ...string) (err error) {
	err = errors.New("error")
	return
}

func commandWithFlag(command *Command, args ...string) (err error) {
	value := command.Flag.String("foo", "", "usage of foo")
	command.Flag.Parse(args)

	if *value != "bar" {
		err = errors.New("error")
	}

	return
}

func TestExecuteCommand(t *testing.T) {
	registry := NewCommandRegistry()

	good := NewCommand("good", "this is usage for good command", "good command", goodCommand)
	bad := NewCommand("bad", "this is usage for bad command", "bad command", badCommand)

	registry.Register(good)
	registry.Register(bad)

	err := registry.Exec([]string{"good", "foo"})

	assert.Nil(t, err)

	err = registry.Exec([]string{"bad", "foo"})

	assert.NotNil(t, err)

	err = registry.Exec([]string{"xxxx", "foo"})

	assert.NotNil(t, err)
}

func TestUsage(t *testing.T) {
	registry := NewCommandRegistry()

	good := NewCommand("good", "this is usage for good command", "good command", goodCommand)
	bad := NewCommand("bad", "this is usage for bad command", "bad command", badCommand)

	registry.Register(good)
	registry.Register(bad)

	usage, err := registry.Usage()
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
	registry := NewCommandRegistry()
	command := NewCommand("test", "this is usage for test command", "test command", commandWithFlag)
	registry.Register(command)

	err := registry.Exec([]string{"test", "-foo", "bar"})

	assert.Nil(t, err)
}
