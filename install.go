package main

import "github.com/ancientstraits/kelmet/cmd"

var InstallCmd = &cmd.Command{
	Name:      "install",
	Usage:     "[name] [chart] [flags]",
	ShortDesc: "install packages",
	LongDesc:  "install a Helm chart",
	LeastArgs: 1,
	Run: func(c *cmd.Command, args []string) error {
		return nil
	},
}

var UninstallCmd = &cmd.Command{
	Name:      "uninstall",
	Usage:     "[name] [chart] [flags]",
	ShortDesc: "uninstall packages",
	LongDesc:  "uninstall a Helm chart",
	LeastArgs: 1,
	Run: func(c *cmd.Command, args []string) error {
		return nil
	},
}
