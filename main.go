package main

import (
	"fmt"
	"log"

	"github.com/ancientstraits/kelmet/cmd"
)

func main() {
	root := &cmd.Command{
		Name:      "kelmet",
		Usage:     "kelmet",
		ShortDesc: "a Kubernetes package manager",
		LongDesc:  "a Kubernetes package manager that is compatible with Helm",
		Run:       cmd.RunUseSubcommands,
	}
	hello := &cmd.Command{
		Name:      "hello",
		Usage:     "hello",
		ShortDesc: "echo hello",
		LongDesc:  "prints out hello",
		Run: func(c *cmd.Command, args []string) error {
			fmt.Println("hello")
			return nil
		},
	}
	root.AddCommand(hello)

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
