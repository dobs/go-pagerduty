package main

import (
	"github.com/mitchellh/cli"
)

type LogEntryList struct {
}

func LogEntryListCommand() (cli.Command, error) {
	return &LogEntryList{}, nil
}

func (c *LogEntryList) Help() string {
	return defaultHelpText
}

func (c *LogEntryList) Synopsis() string {
	return "List all of the incident log entries across the entire account"
}

func (c *LogEntryList) Run(args []string) int {
	return 0
}
