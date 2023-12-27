package triggers

import (
	"context"

	"github.com/stefan79/gadgeto/internal"
	gcontext "github.com/stefan79/gadgeto/pkg/context"
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
	Build(gcontext.GadgetoContext[interface{}]) internal.Trigger[Request, Response]
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

func (c *ApiGatewayConfig) Build(ctx gcontext.GadgetoContext[interface{}]) internal.Trigger[Request, Response] {
	return &ApiGWTrigger{}
}

func ApiGateway(name string) ApiGatewayBuilder {
	return &ApiGatewayConfig{
		Name: &name,
	}

}
