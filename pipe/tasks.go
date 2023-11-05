package pipe

import (
	"strings"
	"syscall"

	. "gitlab.kilic.dev/libraries/plumber/v4"
)

func StepGenerator(tl *TaskList[Pipe]) *Task[Pipe] {
	return tl.CreateTask().
		Set(func(t *Task[Pipe]) error {
			for _, step := range t.Pipe.Steps {
				// every slice element runs in sequence
				st := t.CreateSubtask().
					ShouldRunAfter(func(t *Task[Pipe]) error {
						return t.RunSubtasks()
					}).
					EnableTerminator().
					AddSelfToTheParentAsSequence()

				// multiple elements in a step run as parallel
				for _, s := range step {
					func(s Step) {
						st.CreateSubtask(s.Name).
							ShouldRunBefore(func(t *Task[Pipe]) error {
								if s.Delay.Duration > 0 {
									t.Log.Warnf(
										"Task was delayed: %s",
										s.Delay.String(),
									)
								}

								return nil
							}).
							Set(func(t *Task[Pipe]) error {
								for _, command := range s.Commands {
									func(command string) {
										c := strings.Split(command, " ")

										t.CreateCommand(c[0], c[1:]...).
											Set(func(c *Command[Pipe]) error {
												if s.IgnoreError {
													c.SetIgnoreError()
												}

												c.Command.SysProcAttr = &syscall.SysProcAttr{}

												if s.Syscall.Enable {
													c.Command.SysProcAttr.Credential = &syscall.Credential{
														Uid: s.Syscall.Uid,
														Gid: s.Syscall.Gid,
													}
												}

												return nil
											}).
											AppendEnvironment(s.Environment).
											SetDir(s.Cwd).
											SetRetries(s.Retry.Retries, s.Retry.Always, s.Retry.Delay.Duration).
											SetLogLevel(s.Log.Stdout, s.Log.Stderr, s.Log.Lifetime).
											EnableTerminator().
											AddSelfToTheTask()
									}(command)
								}

								return nil
							}).
							SetJobWrapper(func(job Job) Job {
								return TL.JobDelay(job, s.Delay.Duration)
							}).
							ShouldRunAfter(func(t *Task[Pipe]) error {
								return t.RunCommandJobAsJobSequence()
							}).
							EnableTerminator().
							AddSelfToTheParentAsParallel()
					}(s)
				}
			}

			return nil
		}).
		EnableTerminator().
		ShouldRunAfter(func(t *Task[Pipe]) error {
			return t.RunSubtasks()
		})
}
