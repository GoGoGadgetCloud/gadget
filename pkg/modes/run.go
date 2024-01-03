package modes

import (
	"fmt"
	"os"

	"github.com/stefan79/gadgeto/pkg/resources"
)

type AWSLambdaRuntimeMode[Resource any] struct {
	ResourceFactoryContext *resources.ResourceFactoryContext
}

func (c *AWSLambdaRuntimeMode[Resource]) Dispatch(factory resources.ResourceFactory[Resource]) (Resource, error) {
	return factory.Connect(c.ResourceFactoryContext)

}
func (c *AWSLambdaRuntimeMode[Resource]) Complete() {
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

	return &AWSLambdaRuntimeMode[interface{}]{
		ResourceFactoryContext: &resources.ResourceFactoryContext{
			ApplicationName: &applicationName,
			CommandName:     &commandName,
		},
	}
}
