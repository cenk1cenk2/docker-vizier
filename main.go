package main

import (
	"github.com/urfave/cli/v2"

	"gitlab.kilic.dev/docker/vizier/pipe"
	"gitlab.kilic.dev/docker/vizier/schema"
	. "gitlab.kilic.dev/libraries/plumber/v4"
)

func main() {
	NewPlumber(
		func(p *Plumber) *cli.App {
			return &cli.App{
				Name:        CLI_NAME,
				Version:     VERSION,
				Usage:       DESCRIPTION,
				Description: DESCRIPTION,
				Flags:       p.AppendFlags(pipe.Flags),
				Before: func(ctx *cli.Context) error {
					p.EnableTerminator()

					return nil
				},
				Action: func(ctx *cli.Context) error {
					return pipe.TL.RunJobs(
						pipe.New(p).SetCliContext(ctx).Job(),
					)
				},
				Commands: cli.Commands{
					{
						Name:        "generate",
						Description: "Generate json schema",
						Flags:       p.AppendFlags(schema.Flags),
						Action: func(ctx *cli.Context) error {
							tl := &schema.TL

							return tl.RunJobs(
								schema.New(p).SetCliContext(ctx).Job(),
							)
						},
					},
				},
			}
		}).
		SetDocumentationOptions(DocumentationOptions{
			MarkdownOutputFile: "CLI.md",
			MarkdownBehead:     0,
			ExcludeFlags:       true,
			ExcludeHelpCommand: true,
		}).
		Run()
}
