package modes

import (
	"os"

	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/iam"
	"github.com/awslabs/goformation/v7/cloudformation/lambda"
	"github.com/stefan79/gadgeto/pkg/resources"
)

type UploadLocation struct {
	Bucket *string
	Key    *string
}

type CloudFormationDeployContext[Client any] struct {
	Template *cloudformation.Template
	UploadLocation
	Environment          map[string]string
	TemplateFileLocation *string
}

func (c *CloudFormationDeployContext[Client]) Dispatch(factory resources.ResourceFactory[Client]) (Client, error) {
	return factory.Deploy(c.Template, c.Environment)
}

func NewDeployMode(templateFileName *string, s3Bucket *string) Mode[interface{}] {
	key := "lambda.zip"
	cfContext := &CloudFormationDeployContext[interface{}]{
		Template:             cloudformation.NewTemplate(),
		TemplateFileLocation: templateFileName,
		Environment:          make(map[string]string),
		UploadLocation: UploadLocation{
			Bucket: s3Bucket,
			Key:    &key,
		},
	}
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
	handler := "remote"
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
		Handler: &handler,
	}

}

func (c *CloudFormationDeployContext[Client]) Complete() error {
	c.generateLambdaDescriptor()
	yaml, err := c.Template.YAML()
	if err != nil {
		return err
	}
	err = os.WriteFile(*c.TemplateFileLocation, yaml, 0755)
	return err
}
