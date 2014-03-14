package cemedine_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/youpy/cemedine"
	"testing"
)

func goodCommand(command *cemedine.Command, args ...string) (err error) {
	return
}

func badCommand(command *cemedine.Command, args ...string) (err error) {
	err = errors.New("error")
	return
}

func commandWithFlag(command *cemedine.Command, args ...string) (err error) {
	value := command.Flag.String("foo", "", "usage of foo")
	command.Flag.Parse(args)

	if *value != "bar" {
		err = errors.New("error")
	}

	return
}

func TestExecuteCommand(t *testing.T) {
	cemedine.ResetForTesting()

	good := cemedine.NewCommand("good", "this is usage for good command", "good command", goodCommand)
	bad := cemedine.NewCommand("bad", "this is usage for bad command", "bad command", badCommand)

	cemedine.Register(good)
	cemedine.Register(bad)

	err := cemedine.Exec([]string{"good", "foo"})

	assert.Nil(t, err)

	err = cemedine.Exec([]string{"bad", "foo"})

	assert.NotNil(t, err)

	err = cemedine.Exec([]string{"xxxx", "foo"})

	assert.NotNil(t, err)
}

func TestUsage(t *testing.T) {
	cemedine.ResetForTesting()

	good := cemedine.NewCommand("good", "this is usage for good command", "good command", goodCommand)
	bad := cemedine.NewCommand("bad", "this is usage for bad command", "bad command", badCommand)

	cemedine.Register(good)
	cemedine.Register(bad)

	usage, err := cemedine.Usage()
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
	cemedine.ResetForTesting()

	command := cemedine.NewCommand("test", "this is usage for test command", "test command", commandWithFlag)
	cemedine.Register(command)

	err := cemedine.Exec([]string{"test", "-foo", "bar"})

	assert.Nil(t, err)
}
