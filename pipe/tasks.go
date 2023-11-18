package pipe

import (
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
						ShouldRunBefore(func(t *Task[Pipe]) error {
							for _, permission := range step.Permissions {
								err := handleStepPermission(t, permission)

								if err != nil {
									return err
								}
							}

							return nil
						}).
						Set(func(t *Task[Pipe]) error {
							for _, command := range step.Commands {
								handleStepCommand(t, command)
							}

							return nil
						}).
						SetJobWrapper(func(job Job) Job {
							if step.Delay.Duration > 0 {
								t.Log.Debugf(
									"Task is delayed: %s -> %s",
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
							if step.Parallel {
								t.Log.Debugf(
									"Task will run commands in parallel.",
								)

								return t.RunCommandJobAsJobParallel()
							}

							t.Log.Debugf(
								"Task will run commands in sequence.",
							)

							return t.RunCommandJobAsJobSequence()
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

func handleStepCommand(t *Task[Pipe], command VizierStepCommand) {
	run := strings.Split(command.Command, " ")

	t.CreateCommand(run[0], run[1:]...).
		Set(func(c *Command[Pipe]) error {
			if command.IgnoreError {
				c.SetIgnoreError()
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
}

func handleStepPermission(t *Task[Pipe], permission VizierStepPermission) error {
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
}

func handleStepPermissionForPath(t *Task[Pipe], permission VizierStepPermission, path string, info fs.FileInfo) error {
	if permission.Chown.User != nil && permission.Chown.Group != nil {
		err := os.Chown(path, int(*permission.Chown.User), int(*permission.Chown.Group))

		if err != nil {
			return err
		}

		t.Log.Tracef("Changed the owner of path: %s -> %d:%d", path, *permission.Chown.User, *permission.Chown.Group)
	}

	if info.IsDir() && permission.Chmod.Dir != nil {
		err := os.Chmod(path, *permission.Chmod.Dir)

		if err != nil {
			return err
		}

		t.Log.Tracef("Changed the permission of directory: %s -> %s", path, *permission.Chmod.Dir)
	} else if !info.IsDir() && permission.Chmod.File != nil {
		err := os.Chmod(path, *permission.Chmod.File)

		if err != nil {
			return err
		}

		t.Log.Tracef("Changed the permission of file: %s -> %s", path, *permission.Chmod.File)
	}

	return nil
}
