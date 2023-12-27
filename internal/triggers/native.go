package triggers

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/stefan79/gadgeto/internal"
	gcontext "github.com/stefan79/gadgeto/pkg/context"
	"github.com/stefan79/gadgeto/pkg/context/impl"
)

type NoOpTriggerConfig[Request any, Response any] struct {
}

// Handle implements internal.Trigger.
func (*NoOpTriggerConfig[Request, Response]) Handle(func(context.Context, Request) (Response, error)) {
	//do nothing
}

type LambdaTriggerConfig[Request any, Response any] struct {
}

// Handle implements internal.Trigger.
func (*LambdaTriggerConfig[Request, Response]) Handle(handler func(context.Context, Request) (Response, error)) {
	lambda.Start(handler)
}

type NativeTriggerBuilder[Request any, Response any] interface {
	Build() internal.Trigger[Request, Response]
}

type NativeTriggerBuilderConfig[Request any, Response any] struct {
	Context gcontext.GadgetoContext[internal.Trigger[Request, Response]]
}

// Connect implements resources.ResourceFactory.
func (*NativeTriggerBuilderConfig[Request, Response]) Connect() (internal.Trigger[Request, Response], error) {
	return &LambdaTriggerConfig[Request, Response]{}, nil
}

// Deploy implements resources.ResourceFactory.
func (*NativeTriggerBuilderConfig[Request, Response]) Deploy(tmpl *cloudformation.Template, env map[string]string) (internal.Trigger[Request, Response], error) {
	return &NoOpTriggerConfig[Request, Response]{}, nil
}

type NativeTriggerConfig[Request any, Response any] struct {
}

// Build implements NativeTriggerBuilder.
func (c *NativeTriggerBuilderConfig[Request, Response]) Build() internal.Trigger[Request, Response] {
	trigger, err := c.Context.Dispatch(c)
	if err != nil {
		panic(err)
	}
	return trigger
}

func NewNativeTrigger[Request any, Response any](name string, ctx gcontext.GadgetoContext[interface{}]) NativeTriggerBuilder[Request, Response] {
	return &NativeTriggerBuilderConfig[Request, Response]{
		Context: impl.WithContext[internal.Trigger[Request, Response]](ctx),
	}
}
