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
				switch k {
				case "SerializedDuration":
					v.Type = "string"
					v.Required = nil
					v.Properties = nil
					v.AdditionalProperties = nil
				case "TemplatableBoolean":
					v.Type = ""
					v.OneOf = []*jsonschema.Schema{
						{
							Type: "string",
						},
						{
							Type: "boolean",
						},
					}
					v.Required = nil
					v.Properties = nil
					v.AdditionalProperties = nil
				}
			}

			schema.AllOf = []*jsonschema.Schema{
				{
					Ref: "#/$defs/VizierConfig",
				},
			}

			schema.Ref = ""

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
