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
		Usage:       "Configuration file to read from.",
		Required:    false,
		Value:       "",
		EnvVars:     []string{"VIZIER_CONFIG"},
		Destination: &TL.Pipe.Config.File,
	},

	&cli.StringFlag{
		Name: "steps",
		Usage: `Steps to run for the application, will be ignored when configuration file is read. json([]struct {
  name?: string
  commands?: []struct {
    cwd?: string
    command: string
    retry?: struct {
      retries?: number
      always?: boolean
      delay?: string
    }
    ignore_error?: boolean
    log?: struct {
      stdout?: VizierLogLevels
      stderr?: VizierLogLevels
      lifetime?: VizierLogLevels
    }
    environment?: map[string]string
    run_as?: struct {
      user?: string
      group?: string
    }
  }
  permissions?: []struct {
    path: string
    chown?: struct {
      user?: string
      group?: string
    }
    chmod?: struct {
      file?: string
      dir?: string
    }
    recursive?: boolean
  }
  delay?: string
  background?: boolean
  parallel?: boolean
})
`,
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
