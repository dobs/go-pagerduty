package main

import (
	"github.com/mitchellh/cli"
)

type TeamCreate struct {
}

func TeamCreateCommand() (cli.Command, error) {
	return &TeamCreate{}, nil
}

func (c *TeamCreate) Help() string {
	return defaultHelpText
}

func (c *TeamCreate) Synopsis() string {
	return "Create a new team"
}

func (c *TeamCreate) Run(args []string) int {
	return 0
}
