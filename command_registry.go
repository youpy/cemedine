package cmd

import (
	"bytes"
	"errors"
	"text/template"
)

type CommandRegistry struct {
	commands []*Command
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		commands: []*Command{},
	}
}

func (r *CommandRegistry) Register(command *Command) {
	r.commands = append(r.commands, command)
}

func (r *CommandRegistry) Exec(args []string) (err error) {
	name := args[0]

	for _, command := range r.commands {
		if command.Name == name {
			command.Flag.Parse(args[1:])
			err = command.Run(command, command.Flag.Args()...)
			return
		}
	}

	err = errors.New("Commmand not found: " + name)

	return
}

func (r *CommandRegistry) Usage() (usage string, err error) {
	var b bytes.Buffer
	var usageTmpl = template.Must(template.New("usage").Parse(
		`
Commands:{{range .}}
    {{.Name | printf "%-10s"}} {{.Summary}}{{end}}
`))

	err = usageTmpl.Execute(&b, r.commands)
	if err != nil {
		return
	}

	usage = b.String()

	return
}
