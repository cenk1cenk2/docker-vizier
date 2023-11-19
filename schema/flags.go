package schema

import (
	"github.com/urfave/cli/v2"
	. "gitlab.kilic.dev/libraries/plumber/v4"
)

//revive:disable:line-length-limit

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "output",
		Usage:       "Schema file to write to.",
		Required:    false,
		Value:       "schema.json",
		EnvVars:     []string{"VIZIER_SCHEMA_OUTPUT"},
		Destination: &TL.Pipe.Output,
	},
}

//revive:disable:unused-parameter
func ProcessFlags(tl *TaskList[Pipe]) error {
	return nil
}
