package impl

import (
	"github.com/stefan79/gadgeto/internal/resources"
	"github.com/stefan79/gadgeto/pkg/context"
)

type AWSLambdaRuntimeContext[Client any] struct {
}

func (c *AWSLambdaRuntimeContext[Client]) Dispatch(factory resources.ResourceFactory[Client]) (Client, error) {
	return factory.Connect()

}

func (c *AWSLambdaRuntimeContext[Client]) Complete() error {
	return nil
}

func NewRunContext() context.GadgetoContext[interface{}] {
	return &AWSLambdaRuntimeContext[interface{}]{}
}
