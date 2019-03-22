package main

import (
	"github.com/mitchellh/cli"
)

type MaintenanceWindowList struct {
}

func MaintenanceWindowListCommand() (cli.Command, error) {
	return &MaintenanceWindowList{}, nil
}

func (c *MaintenanceWindowList) Help() string {
	return defaultHelpText
}

func (c *MaintenanceWindowList) Synopsis() string {
	return "List existing maintenance windows"
}

func (c *MaintenanceWindowList) Run(args []string) int {
	return 0
}
