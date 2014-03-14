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

func helloCommand(command *cemedine.Command, args ...string) (err error) {
    value := command.Flag.String("who", "", "name to say hello to")
    command.Flag.Parse(args)

	if *value != "" {
		fmt.Println("hello " + *value)
	}

	return
}

func main() {
    cemedine.Register(
        cemedine.NewCommand(
            "hello",
            "this is usage for hello command",
            "hello command",
            helloCommand,
        ),
    )

	err := cemedine.Exec(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
}
```

```
$ ./cmd

Commands:
    hello      hello command
```

```
$ ./cmd hello -h
Usage:
  -who="": name to say hello to
```

```
$ ./cmd hello -who hoge
hello hoge
```

## Documentation

- https://godoc.org/github.com/youpy/cemedine

## See Also

This package is based on the implementation of [goose](https://bitbucket.org/liamstask/goose/src/a9882a2ed799e698d21706769cd8db004ed68f79/cmd/goose/?at=master)

