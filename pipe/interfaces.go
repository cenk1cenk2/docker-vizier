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

	StepSyscallCredentials struct {
		Enable bool   `json:"enable,omitempty" validate:"bool"`
		Uid    uint32 `json:"uid,omitempty"    validate:"uint32"`
		Gid    uint32 `json:"gid,omitempty"    validate:"uint32"`
	}

	Step struct {
		Name        string                 `json:"name,omitempty"`
		Cwd         string                 `json:"cwd,omitempty"          validate:"dir"`
		Commands    []string               `json:"commands"               validate:"required"`
		Delay       StringDuration         `json:"delay,omitempty"`
		Retry       StepRetry              `json:"retry,omitempty"`
		IgnoreError bool                   `json:"ignore_error,omitempty"`
		Log         StepLogLevel           `json:"log,omitempty"`
		Environment map[string]string      `json:"environment,omitempty"`
		RunAs       StepSyscallCredentials `json:"run_as,omitempty"`
	}
)
