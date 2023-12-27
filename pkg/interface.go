package main

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
