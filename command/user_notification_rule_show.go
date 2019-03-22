package main

import (
	"github.com/mitchellh/cli"
)

type UserNotificationRuleShow struct {
}

func UserNotificationRuleShowCommand() (cli.Command, error) {
	return &UserNotificationRuleShow{}, nil
}

func (c *UserNotificationRuleShow) Help() string {
	return defaultHelpText
}

func (c *UserNotificationRuleShow) Synopsis() string {
	return "Get details about a user's notification rule"
}

func (c *UserNotificationRuleShow) Run(args []string) int {
	return 0
}
