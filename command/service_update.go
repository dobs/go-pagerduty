package main

import (
	"github.com/mitchellh/cli"
)

type ServiceUpdate struct {
}

func ServiceUpdateCommand() (cli.Command, error) {
	return &ServiceUpdate{}, nil
}

func (c *ServiceUpdate) Help() string {
	return defaultHelpText
}

func (c *ServiceUpdate) Synopsis() string {
	return "Update an existing service"
}

func (c *ServiceUpdate) Run(args []string) int {
	return 0
}
