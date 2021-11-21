// Package cmd helps create CLI applications.
// It's like Cobra, but a bit lighter.
package cmd

import (
	"fmt"
	"log"
	"os"
)

var args = os.Args

// Command represents a command.
// Commands can have any amount of subcommands.
type Command struct {
	// The command's name.
	Name string

	// How to use the command.
	Usage string

	// A short description.
	ShortDesc string

	// A longer description.
	LongDesc string

	// The least amount of arguments in the command.
	LeastArgs int

	// How should the command be run?
	// If only its subcommands should be executed,
	// make it equal to `RunUseSubcommands`.
	Run     func(c *Command, args []string) error
	subcmds []*Command
}

// This function is for when
// your command should just
// execute the subcommand.
func RunUseSubcommands(c *Command, args []string) error {
	return c.Execute()
}

// Implements fmt.Stringer.
func (c *Command) String() string {
	ret := fmt.Sprintf("%s %s:\n  %s\n\n", c.Name, c.Usage, c.LongDesc)

	if len(c.subcmds) == 0 {
		ret += fmt.Sprintf("Needs at least %d arguments", c.LeastArgs)
		return ret
	}

	ret += "Available subcommands:\n"
	for _, subcmd := range c.subcmds {
		ret += fmt.Sprintf("  %s: %s\n", subcmd.Name, subcmd.ShortDesc)
	}
	return ret
}

// Adds a subcommand to the command.
func (c *Command) AddCommand(cmds ...*Command) {
	c.subcmds = append(c.subcmds, cmds...)
}

// Executes the command.
func (c *Command) Execute() error {
	args = args[1:]
	if len(c.subcmds) > 0 && len(args) == 0 ||
		len(args) < c.LeastArgs {
		log.Fatal(c)
	}
	for _, subcmd := range c.subcmds {
		if subcmd.Name == args[0] {
			return c.Run(subcmd, args[1:])
		}
	}
	if c.subcmds == nil {
		return c.Run(nil, args)
	} else {
		log.Fatal(c.String() + "Invalid subcommand\n\n")
	}
	return nil
}
