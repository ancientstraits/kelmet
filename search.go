package main

import "github.com/ancientstraits/kelmet/cmd"

var SearchCmd = &cmd.Command{
	Name:      "search",
	Usage:     "[hub | repo]",
	ShortDesc: "search packages",
	LongDesc:  "search for Helm charts",
	Run:       cmd.RunUseSubcommands,
}

var SearchHub = &cmd.Command{
	Name:      "hub",
	Usage:     "[package_name]",
	ShortDesc: "search packages in Hub",
	LongDesc:  "search for Helm charts in Artifact Hub",
	Run: func(c *cmd.Command, args []string) error {
		return nil
	},
}

var SearchRepo = &cmd.Command{
	Name:      "repo",
	Usage:     "[package_name]",
	ShortDesc: "search packages in repo",
	LongDesc:  "search for Helm charts in repositories",
	Run: func(c *cmd.Command, args []string) error {
		return nil
	},
}

func init() {
	SearchCmd.AddCommand(
		SearchHub,
		SearchRepo,
	)
}
