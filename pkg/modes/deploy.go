package modes

import (
	"fmt"
	"os"

	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/iam"
	"github.com/awslabs/goformation/v7/cloudformation/lambda"
	"github.com/awslabs/goformation/v7/cloudformation/tags"
	"github.com/stefan79/gadgeto/internal/iam_generator"
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
	Environment            map[string]string
	CompletionHooks        []resources.CompletionHook
	Handler                *string
	TemplateFileLocation   *string
	ResourceFactoryContext *resources.ResourceFactoryContext
}

func (c *CloudFormationDeployContext[Client]) Dispatch(factory resources.ResourceFactory[Client]) (Client, error) {
	fmt.Println("Dispatch Passing", *c.ResourceFactoryContext.ApplicationName, *c.ResourceFactoryContext.CommandName)
	client, hook, err := factory.Deploy(c.ResourceFactoryContext, c.Template, c.Environment)
	if err != nil {
		return client, err

	}
	if hook != nil {
		c.CompletionHooks = append(c.CompletionHooks, hook)
	}
	return client, nil
}

type DeployModeParam struct {
	TemplateFileName  *string
	Handler           *string
	S3Bucket          *string
	S3Key             *string
	ApplicationPrefix *string
	CommandPrefix     *string
}

func NewDeployMode(param *DeployModeParam) Mode[interface{}] {

	cfContext := &CloudFormationDeployContext[interface{}]{
		Logger:               common.NewDefaultLogger(),
		Template:             cloudformation.NewTemplate(),
		TemplateFileLocation: param.TemplateFileName,
		Handler:              param.Handler,
		Environment:          make(map[string]string),
		CompletionHooks:      make([]resources.CompletionHook, 0),
		UploadLocation: UploadLocation{
			Bucket: param.S3Bucket,
			Key:    param.S3Key,
		},
		ResourceFactoryContext: &resources.ResourceFactoryContext{
			ApplicationName: param.ApplicationPrefix,
			CommandName:     param.CommandPrefix,
		},
	}

	cfContext.generateLambdaDescriptor()
	return cfContext
}

func (c *CloudFormationDeployContext[Client]) generateLambdaDescriptor() {
	lambdaRoleKey := c.ResourceFactoryContext.GenerateCommandResourceKey("iamrole", "role")
	lambdaRoleName := c.ResourceFactoryContext.GenerateCommandResourceName("iamrole")
	lambdaFunctionKey := c.ResourceFactoryContext.GenerateCommandResourceKey("lambdafunction", "lambda")
	lambdaFunctionName := c.ResourceFactoryContext.GenerateCommandResourceName("function")
	c.Template.Resources[lambdaRoleKey] = &iam.Role{
		AssumeRolePolicyDocument: iam_generator.GenerateAssumeRoleForService(iam_generator.Lambda),
		Policies: []iam.Role_Policy{
			{
				PolicyName:     lambdaRoleName + "Policy",
				PolicyDocument: iam_generator.GeneratePermissionForReource(iam_generator.Allow, iam_generator.AllActions, "*"),
			},
		},
		Tags: []tags.Tag{
			{
				Key:   "Foo",
				Value: "Bar",
			},
		},
	}
	runtime := "go1.x"

	c.Environment["GADGETO_APPLICATION_NAME"] = *c.ResourceFactoryContext.ApplicationName
	c.Environment["GADGETO_COMMAND_NAME"] = *c.ResourceFactoryContext.CommandName

	c.Template.Resources[lambdaFunctionKey] = &lambda.Function{
		FunctionName: &lambdaFunctionName,
		Code: &lambda.Function_Code{
			S3Bucket: c.UploadLocation.Bucket,
			S3Key:    c.UploadLocation.Key,
		},
		Environment: &lambda.Function_Environment{
			Variables: c.Environment,
		},
		Role:    cloudformation.GetAtt(lambdaRoleKey, "Arn"),
		Runtime: &runtime,
		Handler: c.Handler,
		Tags: []tags.Tag{
			{
				Key:   "Foo",
				Value: "Bar",
			},
		},
	}
	c.ResourceFactoryContext.LambdaRef = resources.StringPtr(cloudformation.Ref(lambdaFunctionKey))
	c.ResourceFactoryContext.LambdaArn = resources.StringPtr(cloudformation.GetAtt(lambdaFunctionKey, "Arn"))
	c.ResourceFactoryContext.LambdaKey = resources.StringPtr(lambdaFunctionKey)
}

func (c *CloudFormationDeployContext[Client]) Complete() {
	for _, hook := range c.CompletionHooks {
		hook(c.ResourceFactoryContext, c.Template)
	}
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
