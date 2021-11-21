package main

import (
	"fmt"
	"os"
	"path"

	"github.com/ancientstraits/kelmet/cmd"
	"github.com/docker/docker/daemon/graphdriver/copy"
)

var PluginCmd = &cmd.Command{
	Name:      "plugin",
	Usage:     "[command]",
	ShortDesc: "manage plugins",
	LongDesc:  "manage and install plugins",
	Run:       cmd.RunUseSubcommands,
}

func checkPluginPath() error {
	if os.Getenv("HELM_PLUGINS") == "" {
		homepath, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		os.Setenv("HELM_PLUGINS", path.Join(homepath, ".local/share/helm/plugins"))
	}
	return nil
}

var PluginInstall = &cmd.Command{
	Name:      "install",
	Usage:     "[plugin]...",
	ShortDesc: "install one or more Helm plugins",
	LongDesc:  "install one or more Helm plugins",
	LeastArgs: 1,
	Run: func(c *cmd.Command, args []string) error {
		if err := checkPluginPath(); err != nil {
			return err
		}

		plugin := path.Join(os.Getenv("HELM_PLUGINS"), args[0])
		if err := copy.DirCopy(args[0], plugin, copy.Content, false); err != nil {
			return err
		}
		fmt.Printf("%s successfully installed to %s\n", args[0], plugin)

		return nil
	},
}

var PluginUninstall = &cmd.Command{
	Name:      "uninstall",
	Usage:     "[plugin]...",
	ShortDesc: "uninstall one or more Helm plugins",
	LongDesc:  "uninstall one or more Helm plugins",
	LeastArgs: 1,
	Run: func(c *cmd.Command, args []string) error {
		if err := checkPluginPath(); err != nil {
			return err
		}

		plugin := path.Join(os.Getenv("HELM_PLUGINS"), args[0])
		err := os.RemoveAll(plugin)
		if err != nil {
			return err
		}

		fmt.Printf("Successfully uninstalled %s\n", plugin)
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
	LeastArgs: 1,
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
