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
							Set(func(t *Task[Pipe]) error {
								for _, command := range s.Commands {
									func(command string) {
										c := strings.Split(command, " ")

										t.CreateCommand(c[0], c[1:]...).
											AppendEnvironment(s.Environment).
											SetDir(s.Cwd).
											SetRetries(s.Retry.Retries, s.Retry.Always, s.Retry.Delay.Duration).
											AddSelfToTheTask()
									}(command)
								}

								if s.Delay.Duration > 0 {
									t.SetJobWrapper(func(job Job) Job {
										return TL.JobDelay(job, s.Delay.Duration)
									})
								}
								return nil
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
