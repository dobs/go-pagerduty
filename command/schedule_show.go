package main

import (
	"github.com/mitchellh/cli"
)

type ScheduleShow struct {
}

func ScheduleShowCommand() (cli.Command, error) {
	return &ScheduleShow{}, nil
}

func (c *ScheduleShow) Help() string {
	return defaultHelpText
}

func (c *ScheduleShow) Synopsis() string {
	return "Show detailed information about a schedule"
}

func (c *ScheduleShow) Run(args []string) int {
	return 0
}
