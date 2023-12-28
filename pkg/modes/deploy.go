package modes

import (
	"fmt"
	"os"

	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/iam"
	"github.com/awslabs/goformation/v7/cloudformation/lambda"
	"github.com/stefan79/gadgeto/pkg/common"
	"github.com/stefan79/gadgeto/pkg/resources"
)

type UploadLocation struct {
	Bucket *string
	Key    *string
}

type CloudFormationDeployContext[Client any] struct {
	Logger   common.Logger
	Template *cloudformation.Template
	UploadLocation
	Environment          map[string]string
	Handler              *string
	TemplateFileLocation *string
}

func (c *CloudFormationDeployContext[Client]) Dispatch(factory resources.ResourceFactory[Client]) (Client, error) {
	return factory.Deploy(c.Template, c.Environment)
}

func NewDeployMode(templateFileName *string, handler *string, s3Bucket *string, s3Key *string) Mode[interface{}] {

	cfContext := &CloudFormationDeployContext[interface{}]{
		Logger:               common.NewDefaultLogger(),
		Template:             cloudformation.NewTemplate(),
		TemplateFileLocation: templateFileName,
		Handler:              handler,
		Environment:          make(map[string]string),
		UploadLocation: UploadLocation{
			Bucket: s3Bucket,
			Key:    s3Key,
		},
	}
	cfContext.Logger.Out().Info(
		"Run Mode Deploy",
		"template", *templateFileName,
		"bucket", *s3Bucket,
		"key", *s3Key,
	)
	return cfContext
}

func (c *CloudFormationDeployContext[Client]) generateLambdaDescriptor() {
	c.Template.Resources["LambdaRole"] = &iam.Role{
		AssumeRolePolicyDocument: map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				{
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"Service": "lambda.amazonaws.com",
					},
					"Action": "sts:AssumeRole",
				},
			},
		},
		Policies: []iam.Role_Policy{
			{
				PolicyName: "AllPrivileges",
				PolicyDocument: map[string]interface{}{
					"Version": "2012-10-17",
					"Statement": []map[string]interface{}{
						{
							"Effect":   "Allow",
							"Action":   "*",
							"Resource": "*",
						},
					},
				},
			},
		},
	}
	runtime := "go1.x"
	c.Template.Resources["LambdaFunction"] = &lambda.Function{
		Code: &lambda.Function_Code{
			S3Bucket: c.UploadLocation.Bucket,
			S3Key:    c.UploadLocation.Key,
		},
		Environment: &lambda.Function_Environment{
			Variables: c.Environment,
		},
		Role:    cloudformation.GetAtt("LambdaRole", "Arn"),
		Runtime: &runtime,
		Handler: c.Handler,
	}

}

func (c *CloudFormationDeployContext[Client]) Complete() {
	c.generateLambdaDescriptor()
	yaml, err := c.Template.YAML()
	if err != nil {
		err = fmt.Errorf("error generating yaml: %w", err)
		c.Logger.Err().Error(err)
		os.Exit(1)
	}
	err = os.WriteFile(*c.TemplateFileLocation, yaml, 0755)
	if err != nil {
		err = fmt.Errorf("error saving yaml: %w", err)
		c.Logger.Err().Error(err, "location", *c.TemplateFileLocation)
		os.Exit(1)
	}

}
