package impl

import (
	"fmt"

	"github.com/stefan79/gadgeto/pkg/context"
)

func WithContext[Client any](ctx context.GadgetoContext[interface{}]) context.GadgetoContext[Client] {
	cloudformationTemplate, ok := ctx.(*CloudFormationDeployContext[interface{}])
	if ok {
		return &CloudFormationDeployContext[Client]{
			Template:       cloudformationTemplate.Template,
			UploadLocation: cloudformationTemplate.UploadLocation,
			Environment:    cloudformationTemplate.Environment,
		}
	}
	_, ok = ctx.(*AWSLambdaRuntimeContext[interface{}])
	if ok {
		return &AWSLambdaRuntimeContext[Client]{}
	}
	panic(fmt.Errorf("unknown context type"))
}
