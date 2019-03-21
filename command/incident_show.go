package main

import (
	"github.com/mitchellh/cli"
)

type IncidentShow struct {
}

func IncidentShowCommand() (cli.Command, error) {
	return &IncidentShow{}, nil
}

func (c *IncidentShow) Help() string {
	return defaultHelpText
}

func (c *IncidentShow) Synopsis() string {
	return "Show detailed information about an incident"
}

func (c *IncidentShow) Run(args []string) int {
	return 0
}
