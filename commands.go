package main

import (
	"fmt"
)

type CLICommand struct {
	name        string
	description string
}

func (cmd CLICommand) PrintCommand() {
	fmt.Printf("%s: %s\n", cmd.name, cmd.description)
}
