package main

import (
	"github.com/mitchellh/cli"
)

type UserDelete struct {
}

func UserDeleteCommand() (cli.Command, error) {
	return &UserDelete{}, nil
}

func (c *UserDelete) Help() string {
	return defaultHelpText
}

func (c *UserDelete) Synopsis() string {
	return "Remove an existing user"
}

func (c *UserDelete) Run(args []string) int {
	return 0
}
