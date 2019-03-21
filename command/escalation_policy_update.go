package main

import (
	"github.com/mitchellh/cli"
)

type EscalationPolicyUpdateUpdate struct {
}

func EscalationPolicyUpdateCommand() (cli.Command, error) {
	return &EscalationPolicyUpdateUpdate{}, nil
}

func (c *EscalationPolicyUpdateUpdate) Help() string {
	return defaultHelpText
}

func (c *EscalationPolicyUpdateUpdate) Synopsis() string {
	return "Updates an existing escalation policy and rules"
}

func (c *EscalationPolicyUpdateUpdate) Run(args []string) int {
	return 0
}
