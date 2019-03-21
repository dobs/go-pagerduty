package main

import (
	"github.com/mitchellh/cli"
)

type IncidentNoteCreate struct {
}

func IncidentNoteCreateCommand() (cli.Command, error) {
	return &IncidentNoteCreate{}, nil
}

func (c *IncidentNoteCreate) Help() string {
	return defaultHelpText
}

func (c *IncidentNoteCreate) Synopsis() string {
	return "Create a new note for the specified incident"
}

func (c *IncidentNoteCreate) Run(args []string) int {
	return 0
}
