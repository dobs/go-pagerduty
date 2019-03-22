package main

import (
	"github.com/mitchellh/cli"
)

type ServiceDelete struct {
}

func ServiceDeleteCommand() (cli.Command, error) {
	return &ServiceDelete{}, nil
}

func (c *ServiceDelete) Help() string {
	return defaultHelpText
}

func (c *ServiceDelete) Synopsis() string {
	return "Delete an existing service"
}

func (c *ServiceDelete) Run(args []string) int {
	return 0
}
