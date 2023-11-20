package pipe

import (
	"os"

	. "gitlab.kilic.dev/libraries/plumber/v4"
)

//revive:disable:line-length-limit

type (
	VizierChown struct {
		User  *uint32 `json:"user,omitempty"  yaml:"user"  validate:"required" jsonschema:"required,oneof_type=string"`
		Group *uint32 `json:"group,omitempty" yaml:"group"                     jsonschema:"oneof_type=string"`
	}

	VizierChmod struct {
		File *os.FileMode `json:"file,omitempty" yaml:"file" jsonschema:"oneof_type=string"`
		Dir  *os.FileMode `json:"dir,omitempty"  yaml:"dir"  jsonschema:"oneof_type=string"`
	}
)

type (
	VizierStepCommandRetry struct {
		Retries int                `json:"retries,omitempty" yaml:"retries" validate:"gte=0"`
		Always  bool               `json:"always,omitempty"  yaml:"always"`
		Delay   SerializedDuration `json:"delay,omitempty"   yaml:"delay"                    jsonschema:"type=string"`
	}

	VizierStepCommandLogLevel struct {
		Stdout      LogLevel `json:"stdout,omitempty"      yaml:"stdout"      jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace" default:"info"`
		Stderr      LogLevel `json:"stderr,omitempty"      yaml:"stderr"      jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace" default:"warn"`
		Lifetime    LogLevel `json:"lifetime,omitempty"    yaml:"lifetime"    jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace" default:"info"`
		Permissions LogLevel `json:"permissions,omitempty" yaml:"permissions" jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace" default:"debug"`
		Delay       LogLevel `json:"delay,omitempty"       yaml:"delay"       jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace" default:"warn"`
		Background  LogLevel `json:"background,omitempty"  yaml:"background"  jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace" default:"debug"`
	}

	VizierStepPermissionLogLevel struct {
		Chown LogLevel `json:"chown,omitempty" yaml:"chown" jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace" default:"trace"`
		Chmod LogLevel `json:"chmod,omitempty" yaml:"chmod" jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace" default:"trace"`
	}

	VizierStepTemplateLogLevel struct {
		Generation LogLevel `json:"generation,omitempty" yaml:"generation" jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace" default:"info"`
		Context    LogLevel `json:"ctx,omitempty"        yaml:"ctx"        jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace" default:"trace"`
		Chown      LogLevel `json:"chown,omitempty"      yaml:"chown"      jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace" default:"debug"`
		Chmod      LogLevel `json:"chmod,omitempty"      yaml:"chmod"      jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace" default:"debug"`
	}

	VizierStepLogLevel struct {
		Delay      LogLevel `json:"delay,omitempty"      yaml:"delay"      jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace" default:"warn"`
		Background LogLevel `json:"background,omitempty" yaml:"background" jsonschema:"type=string,enum=fatal,enum=error,enum=warn,enum=info,enum=debug,enum=trace" default:"debug"`
	}

	VizierStepCommandRunAs struct {
		VizierChown
	}

	VizierStepCommandPipe struct {
		Stdin bool `json:"stdin,omitempty" yaml:"stdin"`
	}

	VizierStepCommandScript struct {
		Inline *string     `json:"inline,omitempty" yaml:"inline" validate:"required_without=File"`
		File   *string     `json:"file,omitempty"   yaml:"file"   validate:"required_without=Inline,omitempty,file"`
		Ctx    interface{} `json:"ctx,omitempty"    yaml:"ctx"`
	}

	VizierStepCommandHealth struct {
		EnsureIsAlive bool `json:"ensureIsAlive,omitempty" yaml:"ensureIsAlive" validate:"omitempty"`
		IgnoreError   bool `json:"ignoreError,omitempty"   yaml:"ignoreError"`
	}

	VizierStepCommand struct {
		Name          string                    `json:"name,omitempty"          yaml:"name"`
		Cwd           string                    `json:"cwd,omitempty"           yaml:"cwd"           validate:"omitempty,dir"`
		Command       string                    `json:"command"                 yaml:"command"       validate:"required"      jsonschema:"required"`
		Script        *VizierStepCommandScript  `json:"script,omitempty"        yaml:"script"        validate:"omitempty"`
		Retry         VizierStepCommandRetry    `json:"retry,omitempty"         yaml:"retry"         validate:"omitempty"`
		Environment   map[string]string         `json:"environment,omitempty"   yaml:"environment"`
		RunAs         *VizierStepCommandRunAs   `json:"runAs,omitempty"         yaml:"runAs"         validate:"omitempty"`
		Health        VizierStepCommandHealth   `json:"health,omitempty"        yaml:"health"        validate:"omitempty"`
		ShouldDisable TemplatableBoolean        `json:"shouldDisable,omitempty" yaml:"shouldDisable" validate:"omitempty"     jsonschema:"oneof_type=string;boolean"`
		Parallel      bool                      `json:"parallel,omitempty"      yaml:"parallel"`
		Delay         SerializedDuration        `json:"delay,omitempty"         yaml:"delay"                                  jsonschema:"type=string"`
		Background    bool                      `json:"background,omitempty"    yaml:"background"`
		Log           VizierStepCommandLogLevel `json:"log,omitempty"           yaml:"log"           validate:"omitempty"`
		Pipe          VizierStepCommandPipe     `json:"pipe,omitempty"          yaml:"pipe"          validate:"omitempty"`
	}

	VizierStepPermission struct {
		Path          *string                      `json:"path,omitempty"          yaml:"path"          validate:"required"  jsonschema:"required"`
		Chown         VizierChown                  `json:"chown,omitempty"         yaml:"chown"         validate:"omitempty"`
		Chmod         VizierChmod                  `json:"chmod,omitempty"         yaml:"chmod"         validate:"omitempty"`
		Recursive     bool                         `json:"recursive,omitempty"     yaml:"recursive"`
		ShouldDisable TemplatableBoolean           `json:"shouldDisable,omitempty" yaml:"shouldDisable" validate:"omitempty" jsonschema:"oneof_type=string;boolean"`
		Parallel      bool                         `json:"parallel,omitempty"      yaml:"parallel"`
		Log           VizierStepPermissionLogLevel `json:"log,omitempty"           yaml:"log"           validate:"omitempty"`
	}

	VizierStepTemplate struct {
		Inline        *string                    `json:"inline,omitempty"        yaml:"inline"        validate:"required_without=Input"`
		Input         *string                    `json:"input,omitempty"         yaml:"input"         validate:"required_without=Inline,file"`
		Output        string                     `json:"output,omitempty"        yaml:"output"        validate:"required"                     jsonschema:"required"`
		Ctx           interface{}                `json:"ctx,omitempty"           yaml:"ctx"`
		Chmod         VizierChmod                `json:"chmod,omitempty"         yaml:"chmod"`
		Chown         VizierChown                `json:"chown,omitempty"         yaml:"chown"         validate:"omitempty"`
		ShouldDisable TemplatableBoolean         `json:"shouldDisable,omitempty" yaml:"shouldDisable" validate:"omitempty"                    jsonschema:"oneof_type=string;boolean"`
		Parallel      bool                       `json:"parallel,omitempty"      yaml:"parallel"`
		Log           VizierStepTemplateLogLevel `json:"log,omitempty"           yaml:"log"           validate:"omitempty"`
	}

	VizierStep struct {
		Name          string                 `json:"name,omitempty"          yaml:"name"`
		Commands      []VizierStepCommand    `json:"commands,omitempty"      yaml:"commands"      validate:"omitempty,dive"`
		Permissions   []VizierStepPermission `json:"permissions,omitempty"   yaml:"permissions"   validate:"omitempty,dive"`
		Templates     []VizierStepTemplate   `json:"templates,omitempty"     yaml:"templates"     validate:"omitempty,dive"`
		ShouldDisable TemplatableBoolean     `json:"shouldDisable,omitempty" yaml:"shouldDisable" validate:"omitempty"      jsonschema:"oneof_type=string;boolean"`
		Delay         SerializedDuration     `json:"delay,omitempty"         yaml:"delay"                                   jsonschema:"type=string"`
		Background    bool                   `json:"background,omitempty"    yaml:"background"`
		Parallel      bool                   `json:"parallel,omitempty"      yaml:"parallel"`
		Log           VizierStepLogLevel     `json:"log,omitempty"           yaml:"log"           validate:"omitempty"`
	}

	VizierConfig struct {
		Steps []VizierStep `json:"steps" yaml:"steps" validate:"required,dive" jsonschema:"required"`
	}
)
