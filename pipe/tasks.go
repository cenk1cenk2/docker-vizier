package pipe

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	. "gitlab.kilic.dev/libraries/plumber/v4"
)

func StepGenerator(tl *TaskList[Pipe]) *Task[Pipe] {
	return tl.CreateTask().
		Set(func(t *Task[Pipe]) error {
			for _, step := range t.Pipe.Steps {
				func(step VizierStep) {
					t.CreateSubtask(step.Name).
						Set(func(t *Task[Pipe]) error {
							if len(step.Permissions) > 0 {
								st := t.CreateSubtask("permissions").
									ShouldRunAfter(func(t *Task[Pipe]) error {
										return t.RunSubtasks()
									}).
									AddSelfToTheParentAsSequence()

								for _, permission := range step.Permissions {
									handleStepPermission(st, permission).AddSelfToTheParentAsParallel()
								}
							}

							if len(step.Templates) > 0 {
								st := t.CreateSubtask("templates").
									ShouldRunAfter(func(t *Task[Pipe]) error {
										return t.RunSubtasks()
									}).
									AddSelfToTheParentAsSequence()

								for _, template := range step.Templates {
									handleTemplate(st, template).AddSelfToTheParentAsParallel()
								}
							}

							if len(step.Commands) > 0 {
								for _, command := range step.Commands {
									st := handleStepCommand(t, command)

									if step.Parallel {
										st.AddSelfToTheParentAsParallel()
									} else {
										st.AddSelfToTheParentAsSequence()
									}
								}
							}

							return nil
						}).
						SetJobWrapper(func(job Job) Job {
							if step.Delay.Duration > 0 {
								t.Log.Warnf(
									"Task will run with delay: %s -> %s",
									step.Name,
									step.Delay.String(),
								)

								job = TL.JobDelay(job, step.Delay.Duration)
							}

							if step.Background {
								t.Log.Debugf(
									"Task will run in the background: %s",
									step.Name,
								)

								job = TL.JobBackground(job)
							}

							return job
						}).
						ShouldRunAfter(func(t *Task[Pipe]) error {
							return t.RunSubtasks()
						}).
						AddSelfToTheParentAsSequence()
				}(step)
			}

			return nil
		}).
		ShouldRunAfter(func(t *Task[Pipe]) error {
			return t.RunSubtasks()
		})
}

func handleStepCommand(t *Task[Pipe], command VizierStepCommand) *Task[Pipe] {
	return t.CreateSubtask(command.Name).
		Set(func(t *Task[Pipe]) error {
			run := strings.Split(command.Command, " ")

			t.CreateCommand(run[0], run[1:]...).
				Set(func(c *Command[Pipe]) error {
					if command.Health.IgnoreError {
						c.SetIgnoreError()
					}

					if command.Health.EnsureIsAlive {
						c.EnsureIsAlive()
					}

					if command.RunAs != nil {
						c.Command.SysProcAttr = &syscall.SysProcAttr{
							Credential: &syscall.Credential{},
						}

						if command.RunAs.User != nil {
							c.Log.Debugf(
								"Will run the command with uid: %d",
								*command.RunAs.User,
							)
							c.Command.SysProcAttr.Credential.Uid = *command.RunAs.User
						}

						if command.RunAs.Group != nil {
							c.Log.Debugf(
								"Will run the command with gid: %d",
								*command.RunAs.Group,
							)
							c.Command.SysProcAttr.Credential.Gid = *command.RunAs.Group
						}
					}

					return nil
				}).
				AppendEnvironment(command.Environment).
				SetDir(command.Cwd).
				SetRetries(command.Retry.Retries, command.Retry.Always, command.Retry.Delay.Duration).
				SetLogLevel(command.Log.Stdout, command.Log.Stderr, command.Log.Lifetime).
				EnableTerminator().
				AddSelfToTheTask()

			return nil
		}).
		ShouldRunAfter(func(t *Task[Pipe]) error {
			return t.RunCommandJobAsJobSequence()
		})
}

func handleStepPermission(t *Task[Pipe], permission VizierStepPermission) *Task[Pipe] {
	return t.CreateSubtask(*permission.Path).
		Set(func(t *Task[Pipe]) error {
			if !permission.Recursive {
				info, err := os.Lstat(*permission.Path)

				if err != nil {
					return err
				}

				return handleStepPermissionForPath(t, permission, *permission.Path, info)
			}

			return filepath.Walk(*permission.Path, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}

				return handleStepPermissionForPath(t, permission, path, info)
			})
		})
}

func handleTemplate(t *Task[Pipe], template VizierStepTemplate) *Task[Pipe] {
	return t.CreateSubtask(fmt.Sprintf("%s -> %s", template.Input, template.Output)).
		Set(func(t *Task[Pipe]) error {
			tpl, err := os.ReadFile(template.Input)

			if err != nil {
				return err
			}

			render, err := InlineTemplate(string(tpl), template.Ctx)

			if err != nil {
				return err
			}

			t.Log.Infof("Created file from template.")
			t.Log.Debugf("Injected context: %+v", template.Ctx)

			if err := os.WriteFile(template.Output, []byte(render), 0600); err != nil {
				return err
			}

			return handleStepPermission(t, VizierStepPermission{
				Path:      &template.Output,
				Chown:     template.Chown,
				Chmod:     template.Chmod,
				Recursive: false,
			}).
				Run()
		})
}

func handleStepPermissionForPath(t *Task[Pipe], permission VizierStepPermission, path string, info fs.FileInfo) error {
	if permission.Chown.User != nil && permission.Chown.Group != nil {
		err := os.Chown(path, int(*permission.Chown.User), int(*permission.Chown.Group))

		if err != nil {
			return err
		}

		t.Log.Debugf("Changed the owner of path: %s -> %d:%d", path, *permission.Chown.User, *permission.Chown.Group)
	}

	if info.IsDir() && permission.Chmod.Dir != nil {
		err := os.Chmod(path, *permission.Chmod.Dir)

		if err != nil {
			return err
		}

		t.Log.Debugf("Changed the permission of directory: %s -> %s", path, *permission.Chmod.Dir)
	} else if !info.IsDir() && permission.Chmod.File != nil {
		err := os.Chmod(path, *permission.Chmod.File)

		if err != nil {
			return err
		}

		t.Log.Debugf("Changed the permission of file: %s -> %s", path, *permission.Chmod.File)
	}

	return nil
}
