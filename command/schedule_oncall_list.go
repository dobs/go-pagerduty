package main

import (
	"github.com/mitchellh/cli"
)

type ScheduleOncallList struct {
}

func ScheduleOncallListCommand() (cli.Command, error) {
	return &ScheduleOncallList{}, nil
}

func (c *ScheduleOncallList) Help() string {
	return defaultHelpText
}

func (c *ScheduleOncallList) Synopsis() string {
	return "List incidents"
}

func (c *ScheduleOncallList) Run(args []string) int {
	return 0
}
