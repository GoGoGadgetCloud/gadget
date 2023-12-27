package resources

import "github.com/awslabs/goformation/v7/cloudformation"

type (
	ResourceFactory[Client any] interface {
		Deploy(tmpl *cloudformation.Template, env map[string]string) (Client, error)
		Connect() (Client, error)
	}
)
