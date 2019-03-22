package main

import (
	"github.com/mitchellh/cli"
)

type ScheduleDelete struct {
}

func ScheduleDeleteCommand() (cli.Command, error) {
	return &ScheduleDelete{}, nil
}

func (c *ScheduleDelete) Help() string {
	return defaultHelpText
}

func (c *ScheduleDelete) Synopsis() string {
	return "Delete an on-call schedule"
}

func (c *ScheduleDelete) Run(args []string) int {
	return 0
}
