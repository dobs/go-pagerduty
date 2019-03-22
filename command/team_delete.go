package main

import (
	"github.com/mitchellh/cli"
)

type TeamDelete struct {
}

func TeamDeleteCommand() (cli.Command, error) {
	return &TeamDelete{}, nil
}

func (c *TeamDelete) Help() string {
	return defaultHelpText
}

func (c *TeamDelete) Synopsis() string {
	return "Remove an existing team"
}

func (c *TeamDelete) Run(args []string) int {
	return 0
}
