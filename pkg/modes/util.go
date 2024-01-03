package modes

import (
	"fmt"
)

func DowncastMode[Resource any](mode Mode[interface{}]) Mode[Resource] {
	cfdc, ok := mode.(*CloudFormationDeployMode[interface{}])
	if ok {
		return &CloudFormationDeployMode[Resource]{
			Handler:                  cfdc.Handler,
			ResourceFactoryContext:   cfdc.ResourceFactoryContext,
			Logger:                   cfdc.Logger,
			DeploymentRegistrations:  cfdc.DeploymentRegistrations,
			TemplateFileLocation:     cfdc.TemplateFileLocation,
			lambdaFunctionDescriptor: cfdc.lambdaFunctionDescriptor,
		}
	}
	lambdaRuntimeContext, ok := mode.(*AWSLambdaRuntimeMode[interface{}])
	if ok {
		return &AWSLambdaRuntimeMode[Resource]{
			ResourceFactoryContext: lambdaRuntimeContext.ResourceFactoryContext,
		}
	}
	panic(fmt.Errorf("unknown context type"))
}
