package main

import "github.com/ancientstraits/kelmet/cmd"

var PluginCmd = &cmd.Command{
	Name:      "plugin",
	Usage:     "[command]",
	ShortDesc: "manage plugins",
	LongDesc:  "manage and install plugins",
	Run:       cmd.RunUseSubcommands,
}

var PluginInstall = &cmd.Command{
	Name:      "install",
	Usage:     "[plugin]...",
	ShortDesc: "install one or more Helm plugins",
	LongDesc:  "install one or more Helm plugins",
	Run: func(c *cmd.Command, args []string) error {
		return nil
	},
}

var PluginUninstall = &cmd.Command{
	Name:      "uninstall",
	Usage:     "[plugin]...",
	ShortDesc: "uninstall one or more Helm plugins",
	LongDesc:  "uninstall one or more Helm plugins",
	Run: func(c *cmd.Command, args []string) error {
		return nil
	},
}

var PluginList = &cmd.Command{
	Name:      "list",
	Usage:     "",
	ShortDesc: "list installed Helm plugins",
	LongDesc:  "install one or more Helm plugins",
	Run: func(c *cmd.Command, args []string) error {
		return nil
	},
}

var PluginUpdate = &cmd.Command{
	Name:      "update",
	Usage:     "[plugin]...",
	ShortDesc: "update one or more Helm plugins",
	LongDesc:  "update one or more Helm plugins",
	Run: func(c *cmd.Command, args []string) error {
		return nil
	},
}

func init() {
	PluginCmd.AddCommand(
		PluginInstall,
		PluginUninstall,
		PluginList,
		PluginUpdate,
	)
}
