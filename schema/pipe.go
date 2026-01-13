package schema

import (
	. "github.com/cenk1cenk2/plumber/v6"
)

type (
	Pipe struct {
		Output string
	}
)

var TL = TaskList{}

var P = &Pipe{}

func New(p *Plumber) *TaskList {
	return TL.New(p).
		SetRuntimeDepth(1).
		ShouldRunBefore(func(tl *TaskList) error {
			return p.Validate(P)
		}).
		Set(func(tl *TaskList) Job {
			return JobSequence(
				Generate(tl).Job(),
			)
		})
}
