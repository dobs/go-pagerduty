package main

import (
	"github.com/mitchellh/cli"
)

type UserContactMethodDelete struct {
}

func UserContactMethodDeleteCommand() (cli.Command, error) {
	return &UserContactMethodDelete{}, nil
}

func (c *UserContactMethodDelete) Help() string {
	return defaultHelpText
}

func (c *UserContactMethodDelete) Synopsis() string {
	return "Remove a user's contact method"
}

func (c *UserContactMethodDelete) Run(args []string) int {
	return 0
}
