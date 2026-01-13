package pipe

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v3"
)

//revive:disable:line-length-limit

const (
	CATEGORY_CONFIG = "Config"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Category: CATEGORY_CONFIG,
		Name:     "config-file",
		Usage:    "Configuration file to read from. json(https://raw.githubusercontent.com/cenk1cenk2/docker-vizier/main/schema.json)",
		Required: false,
		Value:    "",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar("VIZIER_CONFIG_FILE"),
		),
		Destination:      &P.File,
		ValidateDefaults: false,
		Validator: func(v string) error {
			file, err := os.ReadFile(v)
			if err != nil {
				return err
			}

			switch path.Ext(v) {
			case ".json":
				err := json.Unmarshal(file, &P.Config)
				if err != nil {
					return fmt.Errorf("Can not unmarshal from configuration file: %w", err)
				}
			case ".yaml", ".yml":
				err := yaml.Unmarshal(file, &P.Config)
				if err != nil {
					return fmt.Errorf("Can not unmarshal from configuration file: %w", err)
				}

			}

			return fmt.Errorf("Can not handle configuration file with extension: %s", path.Ext(v))
		},
	},

	&cli.StringFlag{
		Name:     "config",
		Usage:    "Steps to run for the application, will be ignored when configuration file is read. json(https://raw.githubusercontent.com/cenk1cenk2/docker-vizier/main/schema.json)",
		Required: false,
		Sources: cli.NewValueSourceChain(
			cli.EnvVar("VIZIER_CONFIG"),
		),
		Value:            "",
		ValidateDefaults: false,
		Validator: func(v string) error {
			if err := json.Unmarshal([]byte(v), &P.Config); err != nil {
				return fmt.Errorf("Can not unmarshal config: %w", err)
			}

			return nil

		},
	},
}
