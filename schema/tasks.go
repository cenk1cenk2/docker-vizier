package schema

import (
	"os"

	"github.com/invopop/jsonschema"

	. "github.com/cenk1cenk2/plumber/v6"
	"gitlab.kilic.dev/docker/vizier/pipe"
)

func Generate(tl *TaskList) *Task {
	return tl.CreateTask("generate").
		Set(func(t *Task) error {
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

			if err := os.WriteFile(P.Output, json, 0600); err != nil {
				return err
			}

			t.Log.Infof("Generated json schema: %s", P.Output)

			return nil
		})
}
