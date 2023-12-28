package modes

import (
	"fmt"
)

func DowncastMode[Client any](mode Mode[interface{}]) Mode[Client] {
	cloudformationTemplate, ok := mode.(*CloudFormationDeployContext[interface{}])
	if ok {
		return &CloudFormationDeployContext[Client]{
			Template:       cloudformationTemplate.Template,
			UploadLocation: cloudformationTemplate.UploadLocation,
			Environment:    cloudformationTemplate.Environment,
		}
	}
	_, ok = mode.(*AWSLambdaRuntimeContext[interface{}])
	if ok {
		return &AWSLambdaRuntimeContext[Client]{}
	}
	panic(fmt.Errorf("unknown context type"))
}
