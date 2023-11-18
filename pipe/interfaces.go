package pipe

import (
	"os"

	. "gitlab.kilic.dev/libraries/plumber/v4"
)

type (
	VizierPermission struct {
		User  *uint32 `json:"user,omitempty"  validate:"required"`
		Group *uint32 `json:"group,omitempty"`
	}

	VizierChmod struct {
		File *os.FileMode `json:"file,omitempty"`
		Dir  *os.FileMode `json:"dir,omitempty"`
	}
)

type (
	VizierStepCommandRetry struct {
		Retries int          `json:"retries,omitempty" validate:"gt=0"`
		Always  bool         `json:"always,omitempty"`
		Delay   JsonDuration `json:"delay,omitempty"`
	}

	VizierStepCommandLogLevel struct {
		Stdout   LogLevel `json:"stdout,omitempty"   validate:"oneof=fatal error warn info debug trace"`
		Stderr   LogLevel `json:"stderr,omitempty"   validate:"oneof=fatal error warn info debug trace"`
		Lifetime LogLevel `json:"lifetime,omitempty" validate:"oneof=fatal error warn info debug trace"`
	}

	VizierStepCommandRunAs struct {
		VizierPermission
	}

	VizierStepCommand struct {
		Cwd         string                    `json:"cwd,omitempty"          validate:"omitempty,dir"`
		Command     string                    `json:"command"                validate:"required"`
		Retry       VizierStepCommandRetry    `json:"retry,omitempty"        validate:"omitempty"`
		IgnoreError bool                      `json:"ignore_error,omitempty"`
		Log         VizierStepCommandLogLevel `json:"log,omitempty"          validate:"omitempty"`
		Environment map[string]string         `json:"environment,omitempty"`
		RunAs       *VizierStepCommandRunAs   `json:"run_as,omitempty"       validate:"omitempty"`
	}

	VizierStepPermission struct {
		Path      *string          `json:"path,omitempty"      validate:"required"`
		Chown     VizierPermission `json:"chown,omitempty"     validate:"omitempty"`
		Chmod     VizierChmod      `json:"chmod,omitempty"     validate:"omitempty"`
		Recursive bool             `json:"recursive,omitempty"`
	}

	VizierStep struct {
		Name        string                 `json:"name,omitempty"`
		Commands    []VizierStepCommand    `json:"commands,omitempty"    validate:"omitempty,dive"`
		Permissions []VizierStepPermission `json:"permissions,omitempty" validate:"omitempty,dive"`
		Delay       JsonDuration           `json:"delay,omitempty"`
		Background  bool                   `json:"background,omitempty"`
		Parallel    bool                   `json:"parallel,omitempty"`
	}
)
