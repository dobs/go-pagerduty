package main

import (
	"github.com/mitchellh/cli"
)

type UserContactMethodUpdate struct {
}

func UserContactMethodUpdateCommand() (cli.Command, error) {
	return &UserContactMethodUpdate{}, nil
}

func (c *UserContactMethodUpdate) Help() string {
	return defaultHelpText
}

func (c *UserContactMethodUpdate) Synopsis() string {
	return "Update a user's contact method"
}

func (c *UserContactMethodUpdate) Run(args []string) int {
	return 0
}
