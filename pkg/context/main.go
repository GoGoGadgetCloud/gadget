package context

import (
	"fmt"

	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/spf13/cobra"
	"github.com/stefan79/gadgeto/internal/resources"
)

type GagdgetoContext[Client any] interface {
	Complete() error
	Dispatch(factory resources.ResourceFactory[Client]) (Client, error)
}

type CloudFormationDeployContext[Client any] struct {
	Template cloudformation.Template
}

func (c *CloudFormationDeployContext[Client]) Complete() error {
	yaml, err := c.Template.YAML()
	if err == nil {
		fmt.Println("Template: ", string(yaml))
	} else {
		fmt.Println("Error: ", err)
	}
	return err
}

func (c *CloudFormationDeployContext[Client]) Dispatch(factory resources.ResourceFactory[Client]) (Client, error) {
	return factory.Deploy(c.Template)
}

var rootCommand = &cobra.Command{
	Use:   "gadgeto",
	Short: "Gadgeto Meta State Machine",
	Run: func(*cobra.Command, []string) {
		fmt.Println("Hello World")
	},
}

func NewGadgetoContext() GagdgetoContext[interface{}] {
	if err := rootCommand.Execute(); err != nil {
		panic(err)
	}
	return &CloudFormationDeployContext[interface{}]{
		Template: *cloudformation.NewTemplate(),
	}
}

func WithContext[Client any](ctx GagdgetoContext[interface{}]) GagdgetoContext[Client] {
	cloudformationTemplate, ok := ctx.(*CloudFormationDeployContext[interface{}])
	if !ok {
		panic("wrong context type")
	}
	return &CloudFormationDeployContext[Client]{
		Template: cloudformationTemplate.Template,
	}
}
