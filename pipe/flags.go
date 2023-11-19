package pipe

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	. "gitlab.kilic.dev/libraries/plumber/v4"
)

//revive:disable:line-length-limit

const (
	CATEGORY_CONFIG = "Config"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Category:    CATEGORY_CONFIG,
		Name:        "config",
		Usage:       "Configuration file to read from. json(https://raw.githubusercontent.com/cenk1cenk2/docker-vizier/main/schema.json)",
		Required:    false,
		Value:       "",
		EnvVars:     []string{"VIZIER_CONFIG"},
		Destination: &TL.Pipe.Config.File,
	},

	&cli.StringFlag{
		Name:     "steps",
		Usage:    "Steps to run for the application, will be ignored when configuration file is read. json(https://raw.githubusercontent.com/cenk1cenk2/docker-vizier/main/schema.json)",
		Required: false,
		EnvVars:  []string{"VIZIER_STEPS"},
		Value:    "",
	},
}

//revive:disable:unused-parameter
func ProcessFlags(tl *TaskList[Pipe]) error {
	if v := tl.CliContext.String("config"); v != "" {
		file, err := os.ReadFile(v)

		if err != nil {
			return err
		}

		if err := json.Unmarshal(file, &tl.Pipe.Steps); err != nil {
			return fmt.Errorf("Can not unmarshal steps from configuration file: %w", err)
		}
	} else if v := tl.CliContext.String("steps"); v != "" {
		if err := json.Unmarshal([]byte(v), &tl.Pipe.Steps); err != nil {
			return fmt.Errorf("Can not unmarshal steps: %w", err)
		}
	}

	return nil
}
