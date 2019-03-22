package main

import (
	"github.com/mitchellh/cli"
)

type ScheduleUpdate struct {
}

func ScheduleUpdateCommand() (cli.Command, error) {
	return &ScheduleUpdate{}, nil
}

func (c *ScheduleUpdate) Help() string {
	return defaultHelpText
}

func (c *ScheduleUpdate) Synopsis() string {
	return "Update an existing on-call schedule"
}

func (c *ScheduleUpdate) Run(args []string) int {
	return 0
}
