package main

import (
	"github.com/mitchellh/cli"
)

type UserNotificationRuleCreate struct {
}

func UserNotificationRuleCreateCommand() (cli.Command, error) {
	return &UserNotificationRuleCreate{}, nil
}

func (c *UserNotificationRuleCreate) Help() string {
	return defaultHelpText
}

func (c *UserNotificationRuleCreate) Synopsis() string {
	return "Create a new notification rule"
}

func (c *UserNotificationRuleCreate) Run(args []string) int {
	return 0
}
