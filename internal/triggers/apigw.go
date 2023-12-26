package apigw

import (
	"github.com/stefan79/gadgeto/internal"
	"github.com/stefan79/gadgeto/pkg/context"
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
	Build(context.GagdgetoContext[interface{}]) internal.Trigger[Request, Response]
}

type ApiGWTrigger struct {
}

func (t *ApiGWTrigger) Handle(handler func(Request, Response) error) {
}

func (c *ApiGatewayConfig) WithMethod(method string) ApiGatewayBuilder {
	c.Method = &method
	return c
}

func (c *ApiGatewayConfig) Build(ctx context.GagdgetoContext[interface{}]) internal.Trigger[Request, Response] {
	return &ApiGWTrigger{}
}

func ApiGateway(name string) ApiGatewayBuilder {
	return &ApiGatewayConfig{
		Name: &name,
	}

}
