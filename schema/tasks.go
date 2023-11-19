package schema

import (
	"os"

	"github.com/invopop/jsonschema"

	"gitlab.kilic.dev/docker/vizier/pipe"
	. "gitlab.kilic.dev/libraries/plumber/v4"
)

func Generate(tl *TaskList[Pipe]) *Task[Pipe] {
	return tl.CreateTask("generate").
		Set(func(t *Task[Pipe]) error {
			schema := jsonschema.Reflect(&pipe.VizierConfig{})

			for k, v := range schema.Definitions {
				if k == "JsonDuration" {
					v.Type = "string"
					v.Required = nil
					v.Properties = nil
					v.AdditionalProperties = nil
				}
			}

			json, err := schema.MarshalJSON()

			if err != nil {
				return err
			}

			if err := os.WriteFile(t.Pipe.Output, json, 0600); err != nil {
				return err
			}

			t.Log.Infof("Generated json schema: %s", t.Pipe.Output)

			return nil
		})
}
