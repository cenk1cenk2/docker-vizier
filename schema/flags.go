package schema

import (
	. "github.com/cenk1cenk2/plumber/v6"
	"github.com/urfave/cli/v3"
)

//revive:disable:line-length-limit

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:     "output",
		Usage:    "Schema file to write to.",
		Required: false,
		Value:    "schema.json",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar("VIZIER_SCHEMA_OUTPUT"),
		),
		Destination: &P.Output,
	},
}

//revive:disable:unused-parameter
func ProcessFlags(tl *TaskList) error {
	return nil
}
