package cemedine

import (
	"bytes"
	"errors"
	"fmt"
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
	if len(args) == 0 {
		var usage string

		usage, err = r.Usage()
		if err != nil {
			return
		}

		fmt.Println(usage)

		return
	}

	name := args[0]

	for _, command := range r.commands {
		if command.Name == name {
			err = command.Run(command, args[1:]...)
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

var commandLine = NewCommandRegistry()

func Register(command *Command) {
	commandLine.Register(command)
}

func Exec(args []string) (err error) {
	err = commandLine.Exec(args)
	return
}

func Usage() (usage string, err error) {
	usage, err = commandLine.Usage()
	return
}
