package main

import (
	"github.com/mitchellh/cli"
)

type TeamAddUser struct {
}

func TeamAddUserCommand() (cli.Command, error) {
	return &TeamAddUser{}, nil
}

func (c *TeamAddUser) Help() string {
	return defaultHelpText
}

func (c *TeamAddUser) Synopsis() string {
	return "Add a user to a team"
}

func (c *TeamAddUser) Run(args []string) int {
	return 0
}
