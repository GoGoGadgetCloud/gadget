package modes

import (
	"github.com/stefan79/gadgeto/pkg/resources"
)

type AWSLambdaRuntimeContext[Client any] struct {
}

func (c *AWSLambdaRuntimeContext[Client]) Dispatch(factory resources.ResourceFactory[Client]) (Client, error) {
	return factory.Connect()

}

func (c *AWSLambdaRuntimeContext[Client]) Complete() error {
	return nil
}

func NewRunMode() Mode[interface{}] {
	return &AWSLambdaRuntimeContext[interface{}]{}
}
