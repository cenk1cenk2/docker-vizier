package pipe

import (
	. "gitlab.kilic.dev/libraries/plumber/v4"
)

type (
	StepRetry struct {
		Retries int            `json:"retries,omitempty"`
		Always  bool           `json:"always,omitempty"`
		Delay   StringDuration `json:"delay,omitempty"`
	}

	StepLogLevel struct {
		Stdout   LogLevel `json:"stdout,omitempty"   validate:"int"`
		Stderr   LogLevel `json:"stderr,omitempty"   validate:"int"`
		Lifetime LogLevel `json:"lifetime,omitempty" validate:"int"`
	}

	Step struct {
		Name        string            `json:"name"                   validate:"required"`
		Cwd         string            `json:"cwd"                    validate:"dir"`
		Commands    []string          `json:"commands"               validate:"required"`
		Delay       StringDuration    `json:"delay,omitempty"`
		Retry       StepRetry         `json:"retry,omitempty"`
		IgnoreError bool              `json:"ignore_error,omitempty"`
		Log         StepLogLevel      `json:"log,omitempty"`
		Environment map[string]string `json:"environment,omitempty"`
	}
)
