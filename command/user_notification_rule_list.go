package main

import (
	"github.com/mitchellh/cli"
)

type UserNotificationRuleList struct {
}

func UserNotificationRuleListCommand() (cli.Command, error) {
	return &UserNotificationRuleList{}, nil
}

func (c *UserNotificationRuleList) Help() string {
	return defaultHelpText
}

func (c *UserNotificationRuleList) Synopsis() string {
	return "List notification rules of your PagerDuty user"
}

func (c *UserNotificationRuleList) Run(args []string) int {
	return 0
}
