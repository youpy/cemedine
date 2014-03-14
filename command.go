package cemedine

import (
	"flag"
)

type Command struct {
	Run     func(cmd *Command, args ...string) error
	Flag    flag.FlagSet
	Name    string
	Usage   string
	Summary string
}

func NewCommand(name string, usage string, summary string, fn func(cmd *Command, args ...string) error) *Command {
	return &Command{
		Name:    name,
		Usage:   usage,
		Summary: summary,
		Run:     fn,
	}
}
