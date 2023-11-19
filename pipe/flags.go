package pipe

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"slices"

	"github.com/urfave/cli/v2"
	. "gitlab.kilic.dev/libraries/plumber/v4"
	"gopkg.in/yaml.v3"
)

//revive:disable:line-length-limit

const (
	CATEGORY_CONFIG = "Config"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Category:    CATEGORY_CONFIG,
		Name:        "config-file",
		Usage:       "Configuration file to read from. json(https://raw.githubusercontent.com/cenk1cenk2/docker-vizier/main/schema.json)",
		Required:    false,
		Value:       "",
		EnvVars:     []string{"VIZIER_CONFIG_FILE"},
		Destination: &TL.Pipe.File,
	},

	&cli.StringFlag{
		Name:     "config",
		Usage:    "Steps to run for the application, will be ignored when configuration file is read. json(https://raw.githubusercontent.com/cenk1cenk2/docker-vizier/main/schema.json)",
		Required: false,
		EnvVars:  []string{"VIZIER_CONFIG"},
		Value:    "",
	},
}

//revive:disable:unused-parameter
func ProcessFlags(tl *TaskList[Pipe]) error {
	if v := tl.CliContext.String("config-file"); v != "" {
		file, err := os.ReadFile(v)

		if err != nil {
			return err
		}

		ext := path.Ext(v)

		if ext == ".json" {
			err := json.Unmarshal(file, &tl.Pipe.Config)

			if err != nil {
				return fmt.Errorf("Can not unmarshal from configuration file: %w", err)
			}
		} else if slices.Contains([]string{".yaml", ".yml"}, ext) {
			err := yaml.Unmarshal(file, &tl.Pipe.Config)

			if err != nil {
				return fmt.Errorf("Can not unmarshal from configuration file: %w", err)
			}
		} else {
			return fmt.Errorf("Can not handle configuration file with extension: %s", ext)
		}
	} else if v := tl.CliContext.String("config"); v != "" {
		if err := json.Unmarshal([]byte(v), &tl.Pipe.Config); err != nil {
			return fmt.Errorf("Can not unmarshal config: %w", err)
		}
	}

	return nil
}
