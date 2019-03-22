package main

import (
	"github.com/mitchellh/cli"
)

type IncidentSnooze struct {
}

func IncidentSnoozeCommand() (cli.Command, error) {
	return &IncidentSnooze{}, nil
}

func (c *IncidentSnooze) Help() string {
	return defaultHelpText
}

func (c *IncidentSnooze) Synopsis() string {
	return "Snooze an incident"
}

func (c *IncidentSnooze) Run(args []string) int {
	return 0
}
