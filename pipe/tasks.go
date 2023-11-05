package pipe

import (
	"strings"

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
											ShouldRunBefore(func(c *Command[Pipe]) error {
												c.Log.WithField(LOG_FIELD_STATUS, "RUN").
													Logf(s.Log.Lifetime, "%s", command)

												return nil
											}).
											ShouldRunAfter(func(c *Command[Pipe]) error {
												c.Log.WithField(LOG_FIELD_STATUS, "END").
													Logf(s.Log.Lifetime, "%s", command)

												return nil
											}).
											Set(func(c *Command[Pipe]) error {
												if s.IgnoreError {
													c.SetIgnoreError()
												}

												return nil
											}).
											AppendEnvironment(s.Environment).
											SetDir(s.Cwd).
											SetRetries(s.Retry.Retries, s.Retry.Always, s.Retry.Delay.Duration).
											SetLogLevel(s.Log.Stdout, s.Log.Stderr, LOG_LEVEL_DEBUG).
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
							AddSelfToTheParentAsParallel()
					}(s)
				}
			}

			return nil
		}).
		ShouldRunAfter(func(t *Task[Pipe]) error {
			return t.RunSubtasks()
		})
}
