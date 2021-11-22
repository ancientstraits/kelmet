package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/ancientstraits/kelmet/cmd"
	"gopkg.in/yaml.v2"
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
		if err := os.Symlink(args[0], plugin); err != nil {
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

type PluginConf struct {
	Name            string `yaml:"name"`
	Version         string `yaml:"version"`
	Usage           string `yaml:"usage"`
	Description     string `yaml:"description"`
	IgnoreFlags     bool   `yaml:"ignoreFlags"`
	Command         string `yaml:"command"`
	PlatformCommand []struct {
		Os      string `yaml:"os"`
		Arch    string `yaml:"arch"`
		Command string `yaml:"command"`
	} `yaml:"platformCommand"`
}

// p is a YAML file.
func genPlugin(p string) *cmd.Command {
	pcontent, err := os.ReadFile(p)
	if err != nil {
		panic(err)
	}

	pc := PluginConf{}
	if err := yaml.Unmarshal([]byte(pcontent), &pc); err != nil {
		panic(err)
	}

	ret := &cmd.Command{
		Name:      pc.Name,
		Usage:     pc.Usage,
		ShortDesc: pc.Description,
		LongDesc:  pc.Description,
		LeastArgs: 0,
		Run: func(c *cmd.Command, args []string) error {
			cmdargs := strings.Split(pc.Command, " ")
			com := exec.Command(cmdargs[0], cmdargs[1:]...)
			com.Stdin = os.Stdin
			com.Stderr = os.Stderr
			com.Stdout = os.Stdout
			return com.Run()
		},
	}
	return ret
}

var Plugins = []*cmd.Command{}

func init() {
	PluginCmd.AddCommand(
		PluginInstall,
		PluginUninstall,
		PluginList,
		PluginUpdate,
	)

	if err := checkPluginPath(); err != nil {
		panic(err)
	}

	plugdir := os.Getenv("HELM_PLUGINS")
	entries, err := os.ReadDir(plugdir)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		metadata := path.Join(plugdir, entry.Name(), "plugin.yaml")
		if _, err := os.Stat(metadata); errors.Is(err, os.ErrNotExist) {
			continue
		}

		Plugins = append(Plugins, genPlugin(metadata))
	}
}
