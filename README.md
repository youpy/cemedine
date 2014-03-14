# cemedine

A Go library to build subcommand for executable

## Usage

```go
package main

import (
    "fmt"
    "github.com/youpy/cemedine"
    "log"
    "os"
)

func runCommand(command *cemedine.Command, args ...string) (err error) {
    value := command.Flag.String("who", "", "name to say hello to")
    command.Flag.Parse(args)

	if *value != "" {
		fmt.Println("hello " + *value)
	}

	return
}

func main() {
    cemedine.Register(cemedine.NewCommand("hello", "this is usage", "hello command", runCommand))

	err := cemedine.Exec(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
}
```
