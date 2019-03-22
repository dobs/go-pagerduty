package main

import (
	"github.com/mitchellh/cli"
)

type UserContactMethodShow struct {
}

func UserContactMethodShowCommand() (cli.Command, error) {
	return &UserContactMethodShow{}, nil
}

func (c *UserContactMethodShow) Help() string {
	return defaultHelpText
}

func (c *UserContactMethodShow) Synopsis() string {
	return "Get details about a user's contact method"
}

func (c *UserContactMethodShow) Run(args []string) int {
	return 0
}
