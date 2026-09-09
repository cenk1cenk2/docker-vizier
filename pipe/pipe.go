package pipe

import (
	"context"

	. "github.com/cenk1cenk2/plumber/v7"
)

type (
	Pipe struct {
		File   string
		Config VizierConfig `validate:"required"`
	}
)

var TL = TaskList{}

var P = &Pipe{}

func New(p *Plumber) *TaskList {
	return TL.New(p).
		SetRuntimeDepth(1).
		ShouldRunBefore(func(_ context.Context, tl *TaskList) error {
			return p.Validate(P)
		}).
		Set(func(tl *TaskList) Job {
			return StepGenerator(tl)
		})
}
