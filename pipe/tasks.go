package pipe

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"syscall"

	. "github.com/cenk1cenk2/plumber/v7"
	"github.com/mattn/go-shellwords"
)

func StepGenerator(tl *TaskList) Job {
	job := CreateEmptyJob()

	for _, step := range P.Config.Steps {
		func(step VizierStep) {
			task := tl.CreateTask(step.Name).
				ShouldDisable(func(_ *Task) bool {
					return step.ShouldDisable.bool
				}).
				Set(func(ctx context.Context, t *Task) error {
					if len(step.Permissions) > 0 {
						st := t.CreateSubtask("permissions").
							ShouldRunAfter(func(ctx context.Context, t *Task) error {
								return t.RunSubtasks(ctx)
							}).
							AddSelfToTheParentAsSequence()

						for _, permission := range step.Permissions {
							handleStepPermission(st, permission).
								AddSelfToTheParent(func(pt *Task, st *Task) {
									pt.ExtendSubtask(func(job Job) Job {
										if permission.Parallel {
											return JobParallel(job, st.Job())
										}

										return JobSequence(job, st.Job())
									})
								})
						}
					}

					if len(step.Templates) > 0 {
						st := t.CreateSubtask("templates").
							ShouldRunAfter(func(ctx context.Context, t *Task) error {
								return t.RunSubtasks(ctx)
							}).
							AddSelfToTheParentAsSequence()

						for _, template := range step.Templates {
							handleTemplate(st, template).
								AddSelfToTheParent(func(pt *Task, st *Task) {
									pt.ExtendSubtask(func(job Job) Job {
										if template.Parallel {
											return JobParallel(job, st.Job())
										}

										return JobSequence(job, st.Job())
									})
								})
						}
					}

					if len(step.Commands) > 0 {
						st := t.CreateSubtask().
							ShouldRunAfter(func(ctx context.Context, t *Task) error {
								return t.RunSubtasks(ctx)
							}).
							AddSelfToTheParentAsSequence()

						for _, command := range step.Commands {
							handleStepCommand(st, command).
								AddSelfToTheParent(func(pt *Task, st *Task) {
									pt.ExtendSubtask(func(job Job) Job {
										if command.Parallel {
											return JobParallel(job, st.Job())
										}

										return JobSequence(job, st.Job())
									})
								})
						}
					}

					return nil
				}).
				SetJobWrapper(func(job Job, t *Task) Job {
					if step.Delay.Duration > 0 {
						t.Log.Log(
							context.Background(),
							step.Log.Delay.GetLogLevel(),
							fmt.Sprintf("Task will run with delay: %s", step.Delay.String()),
						)

						job = JobDelay(job, step.Delay.Duration)
					}

					if step.Background {
						t.Log.Log(
							context.Background(),
							step.Log.Background.GetLogLevel(),
							"Task will run in the background.",
						)

						job = JobBackground(job, t.Log)
					}

					return job
				}).
				ShouldRunAfter(func(ctx context.Context, t *Task) error {
					return t.RunSubtasks(ctx)
				})

			if step.Parallel {
				job = JobParallel(job, task.Job())
			} else {
				job = JobSequence(job, task.Job())
			}
		}(step)
	}

	return job
}

func handleStepCommand(t *Task, command VizierStepCommand) *Task {
	return t.CreateSubtask(command.Name).
		ShouldDisable(func(_ *Task) bool {
			return command.ShouldDisable.bool
		}).
		Set(func(ctx context.Context, t *Task) error {
			run, err := shellwords.Parse(command.Command)
			if err != nil {
				return fmt.Errorf("failed to parse command %q: %w", command.Command, err)
			}

			t.CreateCommand(run[0], run[1:]...).
				Set(func(_ context.Context, c *Command) error {
					if command.Health.IgnoreError {
						c.SetIgnoreError()
					}

					if command.Health.EnsureIsAlive {
						c.EnsureIsAlive()
					}

					return nil
				}).
				SetScript(func(_ *Command) *CommandScript {
					if command.Script != nil {
						if command.Script.Inline != nil {
							return &CommandScript{
								Inline: *command.Script.Inline,
								Ctx:    command.Script.Ctx,
							}
						} else if command.Script.File != nil {
							return &CommandScript{
								File: *command.Script.File,
								Ctx:  command.Script.Ctx,
							}
						}
					}

					return nil
				}).
				AppendEnvironment(command.Environment).
				SetDir(command.Cwd).
				SetRetries(
					&CommandRetry{
						Tries:  command.Retry.Retries,
						Always: command.Retry.Always,
						Delay:  command.Retry.Delay.Duration,
					}).
				SetCredential(func(c *Command, credential *syscall.Credential) *syscall.Credential {
					if command.RunAs != nil {
						if command.RunAs.User != nil {
							c.Log.Log(
								context.Background(),
								command.Log.Permissions.GetLogLevel(),
								fmt.Sprintf("Will run the command with uid: %d", *command.RunAs.User),
							)

							credential.Uid = *command.RunAs.User
						}

						if command.RunAs.Group != nil {
							c.Log.Log(
								context.Background(),
								command.Log.Permissions.GetLogLevel(),
								fmt.Sprintf("Will run the command with gid: %d", *command.RunAs.Group),
							)

							credential.Gid = *command.RunAs.Group
						}

						return credential
					}

					return nil
				}).
				SetLogLevel(command.Log.Stdout, command.Log.Stderr, command.Log.Lifetime).
				SetJobWrapper(func(job Job, c *Command) Job {
					if command.Delay.Duration > 0 {
						t.Log.Log(
							context.Background(),
							command.Log.Delay.GetLogLevel(),
							fmt.Sprintf("Command will run with delay: %s -> %s", c.GetFormattedCommand(), command.Delay.String()),
						)

						job = JobDelay(job, command.Delay.Duration)
					}

					if command.Background {
						t.Log.Log(
							context.Background(),
							command.Log.Background.GetLogLevel(),
							fmt.Sprintf("Command will run in the background: %s", c.GetFormattedCommand()),
						)

						job = JobBackground(job, c.Log)
					}

					return job
				}).
				EnableTerminator().
				AddSelfToTheTask()

			return nil
		}).
		ShouldRunAfter(func(ctx context.Context, t *Task) error {
			return t.RunCommandJobAsJobSequence(ctx)
		})
}

func handleStepPermission(t *Task, permission VizierStepPermission) *Task {
	return t.CreateSubtask(*permission.Path).
		ShouldDisable(func(_ *Task) bool {
			return permission.ShouldDisable.bool
		}).
		Set(func(ctx context.Context, t *Task) error {
			if !permission.Recursive {
				info, err := os.Lstat(*permission.Path)

				if err != nil {
					return err
				}

				return applyStepPermissionForPath(ctx, t, permission, *permission.Path, info)
			}

			return filepath.Walk(*permission.Path, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}

				return applyStepPermissionForPath(ctx, t, permission, path, info)
			})
		})
}

func handleTemplate(t *Task, template VizierStepTemplate) *Task {
	if template.Input != nil {
		return t.CreateSubtask(fmt.Sprintf("%s -> %s", *template.Input, template.Output)).
			ShouldDisable(func(_ *Task) bool {
				return template.ShouldDisable.bool
			}).
			Set(func(ctx context.Context, t *Task) error {
				tpl, err := os.ReadFile(*template.Input)

				if err != nil {
					return err
				}

				return applyStepTemplateForInline(ctx, t, template, string(tpl))
			})
	} else if template.Inline != nil {
		return t.CreateSubtask(fmt.Sprintf("%s -> %s", "inline", template.Output)).
			Set(func(ctx context.Context, t *Task) error {
				return applyStepTemplateForInline(ctx, t, template, *template.Inline)
			})
	}

	return nil
}

func applyStepPermissionForPath(ctx context.Context, t *Task, permission VizierStepPermission, path string, info fs.FileInfo) error {
	if permission.Chown.User != nil && permission.Chown.Group != nil {
		err := os.Chown(path, int(*permission.Chown.User), int(*permission.Chown.Group))

		if err != nil {
			return err
		}

		t.Log.Log(ctx, permission.Log.Chown.GetLogLevel(), fmt.Sprintf("Changed the owner of path: %s -> %d:%d", path, *permission.Chown.User, *permission.Chown.Group))
	}

	if info.IsDir() && permission.Chmod.Dir != nil {
		err := os.Chmod(path, *permission.Chmod.Dir)

		if err != nil {
			return err
		}

		t.Log.Log(ctx, permission.Log.Chmod.GetLogLevel(), fmt.Sprintf("Changed the permission of directory: %s -> %s", path, *permission.Chmod.Dir))
	} else if !info.IsDir() && permission.Chmod.File != nil {
		err := os.Chmod(path, *permission.Chmod.File)

		if err != nil {
			return err
		}

		t.Log.Log(ctx, permission.Log.Chmod.GetLogLevel(), fmt.Sprintf("Changed the permission of file: %s -> %s", path, *permission.Chmod.File))
	}

	return nil
}

func applyStepTemplateForInline(ctx context.Context, t *Task, template VizierStepTemplate, tpl string) error {
	render, err := InlineTemplate(tpl, template.Ctx)

	if err != nil {
		return err
	}

	t.Log.Log(ctx, template.Log.Generation.GetLogLevel(), "Created file from template.")
	t.Log.Log(ctx, template.Log.Context.GetLogLevel(), fmt.Sprintf("Injected context: %+v", template.Ctx))

	if err := os.WriteFile(template.Output, []byte(render), 0600); err != nil {
		return err
	}

	return handleStepPermission(t, VizierStepPermission{
		Path:  &template.Output,
		Chown: template.Chown,
		Chmod: template.Chmod,
		Log: VizierStepPermissionLogLevel{
			Chown: template.Log.Chown,
			Chmod: template.Log.Chmod,
		},
		Recursive: false,
	}).
		Run(ctx)
}
