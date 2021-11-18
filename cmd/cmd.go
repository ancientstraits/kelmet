// Package cmd helps create CLI applications.
// It's like Cobra, but a bit lighter.
package cmd

import (
	"fmt"
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

// Adds a subcommand to the command.
func (c *Command) AddCommand(cmds ...*Command) {
	c.subcmds = append(c.subcmds, cmds...)
}

// Executes the command.
func (c *Command) Execute() error {
	if args = args[1:]; len(args) < 1 && len(c.subcmds) > 0 {
		fmt.Printf("%s %s\n%s\n", c.Name, c.Usage, c.LongDesc)
		for _, subcmd := range c.subcmds {
			fmt.Printf("\t%s: %s\n", subcmd.Name, subcmd.ShortDesc)
		}
		return fmt.Errorf("%s doesn't have subcommand", c.Name)
	}
	for _, subcmd := range c.subcmds {
		if subcmd.Name == args[0] {
			return c.Run(subcmd, args[1:])
		}
	}
	if c.subcmds == nil {
		return c.Run(nil, args)
	} else {
		return fmt.Errorf("invalid command")
	}
}
