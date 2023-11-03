package pipe

type (
	StepRetry struct {
		Retries int            `json:"retries,omitempty"`
		Always  bool           `json:"always,omitempty"`
		Delay   StringDuration `json:"delay,omitempty"`
	}

	Step struct {
		Name        string            `json:"name"                  validate:"required"`
		Cwd         string            `json:"cwd"                   validate:"dir"`
		Commands    []string          `json:"commands"              validate:"required"`
		Delay       StringDuration    `json:"delay,omitempty"`
		Retry       StepRetry         `json:"retry,omitempty"`
		Environment map[string]string `json:"environment,omitempty"`
	}
)
