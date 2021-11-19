package main

import (
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
	root.AddCommand(
		PluginCmd,
		SearchCmd,
	)

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
