package pipe

import (
	. "gitlab.kilic.dev/libraries/plumber/v4"
)

type (
	Config struct {
		File string
	}

	Pipe struct {
		Config
		Steps [][]Step
	}
)

var TL = TaskList[Pipe]{
	Pipe: Pipe{},
}

func New(p *Plumber) *TaskList[Pipe] {
	return TL.New(p).
		SetRuntimeDepth(1).
		ShouldRunBefore(func(tl *TaskList[Pipe]) error {
			return ProcessFlags(tl)
		}).
		Set(func(tl *TaskList[Pipe]) Job {
			return tl.JobSequence(
				StepGenerator(tl).Job(),
			)
		})
}
