package apigw

import (
	"context"

	"github.com/stefan79/gadgeto/pkg/modes"
	"github.com/stefan79/gadgeto/pkg/triggers"
)

type ApiGatewayConfig struct {
	Name   *string
	Method *string
}

type Request struct {
	QueryParams map[string]string
	Body        []byte
}

type Response struct {
	ResponseCode int
}

type ApiGatewayBuilder interface {
	WithMethod(string) ApiGatewayBuilder
	Build(modes.Mode[interface{}]) triggers.Trigger[Request, Response]
}

type ApiGWTrigger struct {
}

// Handle implements internal.Trigger.
func (*ApiGWTrigger) Handle(func(context.Context, Request) (Response, error)) {
	panic("unimplemented")
}

func (c *ApiGatewayConfig) WithMethod(method string) ApiGatewayBuilder {
	c.Method = &method
	return c
}

func (c *ApiGatewayConfig) Build(mode modes.Mode[interface{}]) triggers.Trigger[Request, Response] {
	return &ApiGWTrigger{}
}

func ApiGateway(name string) ApiGatewayBuilder {
	return &ApiGatewayConfig{
		Name: &name,
	}

}
