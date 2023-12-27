package internal

import "context"

type (
	Trigger[REQ any, RES any] interface {
		Handle(func(context.Context, REQ) (RES, error))
	}
)

type GagdgetoContext interface {
}

func NewGadgetoContext() GagdgetoContext {
	return nil
}

func GetMainFunction(setup interface{}) func() error {
	return RunLocal
}

func RunLocal() error {
	return nil
}
