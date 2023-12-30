package modes

import (
	"fmt"
	"os"

	"github.com/stefan79/gadgeto/pkg/resources"
)

type AWSLambdaRuntimeContext[Client any] struct {
	ResourceFactoryContext *resources.ResourceFactoryContext
}

func (c *AWSLambdaRuntimeContext[Client]) Dispatch(factory resources.ResourceFactory[Client]) (Client, error) {
	return factory.Connect(c.ResourceFactoryContext)

}
func (c *AWSLambdaRuntimeContext[Client]) Complete() {
}

func NewRunMode() Mode[interface{}] {
	applicationName, found := os.LookupEnv("GADGETO_APPLICATION_NAME")
	if !found {
		panic("GADGETO_APPLICATION_NAME not set")
	}
	commandName, found := os.LookupEnv("GADGETO_COMMAND_NAME")
	if !found {
		panic("GADGETO_COMMAND_NAME not set")
	}

	fmt.Println("ApplicationName: ", applicationName, "CommandName: ", commandName)

	return &AWSLambdaRuntimeContext[interface{}]{
		ResourceFactoryContext: &resources.ResourceFactoryContext{
			ApplicationName: &applicationName,
			CommandName:     &commandName,
		},
	}
}
