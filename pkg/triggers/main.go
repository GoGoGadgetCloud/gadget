package triggers

import "context"

type (
	Trigger[REQ any, RES any] interface {
		Handle(func(context.Context, REQ) (RES, error))
	}
)
