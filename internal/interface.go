package internal

type (
	Trigger[REQ any, RES any] interface {
		Handle(func(REQ, RES) error)
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
