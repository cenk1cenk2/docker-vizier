package pipe

import (
	"os"

	. "gitlab.kilic.dev/libraries/plumber/v4"
)

type (
	VizierChown struct {
		User  *uint32 `json:"user,omitempty"  validate:"required" jsonschema:"required,oneof_type=string"`
		Group *uint32 `json:"group,omitempty"                     jsonschema:"oneof_type=string"`
	}

	VizierChmod struct {
		File *os.FileMode `json:"file,omitempty" jsonschema:"oneof_type=string"`
		Dir  *os.FileMode `json:"dir,omitempty"  jsonschema:"oneof_type=string"`
	}
)

type (
	VizierStepCommandRetry struct {
		Retries int          `json:"retries,omitempty" validate:"gte=0"`
		Always  bool         `json:"always,omitempty"`
		Delay   JsonDuration `json:"delay,omitempty"                    jsonschema:"type=string"`
	}

	VizierStepCommandLogLevel struct {
		Stdout   LogLevel `json:"stdout,omitempty"   jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace"`
		Stderr   LogLevel `json:"stderr,omitempty"   jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace"`
		Lifetime LogLevel `json:"lifetime,omitempty" jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace"`
	}

	VizierStepCommandRunAs struct {
		VizierChown
	}

	VizierStepCommandHealth struct {
		EnsureIsAlive bool `json:"ensureIsAlive,omitempty" validate:"omitempty"`
		IgnoreError   bool `json:"ignoreError,omitempty"`
	}

	VizierStepCommand struct {
		Name        string                    `json:"name,omitempty"`
		Cwd         string                    `json:"cwd,omitempty"         validate:"omitempty,dir"`
		Command     string                    `json:"command"               validate:"required"      jsonschema:"required"`
		Retry       VizierStepCommandRetry    `json:"retry,omitempty"       validate:"omitempty"`
		Log         VizierStepCommandLogLevel `json:"log,omitempty"         validate:"omitempty"`
		Environment map[string]string         `json:"environment,omitempty"`
		RunAs       *VizierStepCommandRunAs   `json:"runAs,omitempty"       validate:"omitempty"`
		Health      VizierStepCommandHealth   `json:"health,omitempty"      validate:"omitempty"`
	}

	VizierStepPermission struct {
		Path      *string     `json:"path,omitempty"      validate:"required"  jsonschema:"required"`
		Chown     VizierChown `json:"chown,omitempty"     validate:"omitempty"`
		Chmod     VizierChmod `json:"chmod,omitempty"     validate:"omitempty"`
		Recursive bool        `json:"recursive,omitempty"`
	}

	VizierStepTemplate struct {
		Input  string      `json:"input,omitempty"  validate:"required,file" jsonschema:"required"`
		Output string      `json:"output,omitempty" validate:"required"      jsonschema:"required"`
		Ctx    interface{} `json:"ctx,omitempty"`
		Chmod  VizierChmod `json:"chmod,omitempty"`
		Chown  VizierChown `json:"chown,omitempty"  validate:"omitempty"`
	}

	VizierStep struct {
		Name        string                 `json:"name,omitempty"`
		Commands    []VizierStepCommand    `json:"commands,omitempty"    validate:"omitempty,dive"`
		Permissions []VizierStepPermission `json:"permissions,omitempty" validate:"omitempty,dive"`
		Templates   []VizierStepTemplate   `json:"templates,omitempty"   validate:"omitempty,dive"`
		Delay       JsonDuration           `json:"delay,omitempty"                                 jsonschema:"type=string"`
		Background  bool                   `json:"background,omitempty"`
		Parallel    bool                   `json:"parallel,omitempty"`
	}

	VizierConfig struct {
		Steps []VizierStep `json:"steps" validate:"required,dive" jsonschema:"required"`
	}
)
