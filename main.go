package main

import (
	"fmt"
	"os"
)

type Cmd struct {
	desc string
	exec func(args ...string) bool
}
type CmdList map[string]Cmd

func (c CmdList) Execute() bool {
	for name, cmd := range c {
		if name == os.Args[1] {
			return cmd.exec(os.Args[2:]...)
		}
	}
	return false
}

func main() {
	program := CmdList {
		"hello": Cmd {
			desc: "Print hello",
			exec: func(args ...string) bool {
				fmt.Println("hello")
				return true
			},
		},
	}
	program.Execute()
}
