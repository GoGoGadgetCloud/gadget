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

type lambdaFunctionDescriptor struct {
	deploymentSourceBucket *string
	deployementSourceKey   *string
	handlerName            *string
}

type CloudFormationDeployMode[Client any] struct {
	Logger                   common.Logger
	Handler                  *string
	TemplateFileLocation     *string
	ResourceFactoryContext   *resources.ResourceFactoryContext
	DeploymentRegistrations  *resources.DeploymentRegistrations
	lambdaFunctionDescriptor *lambdaFunctionDescriptor
}

func (c *CloudFormationDeployMode[Resource]) Dispatch(factory resources.ResourceFactory[Resource]) (Resource, error) {
	return factory.Deploy(c.ResourceFactoryContext, c.DeploymentRegistrations)
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

	lfd := &lambdaFunctionDescriptor{
		deploymentSourceBucket: param.S3Bucket,
		deployementSourceKey:   param.S3Key,
		handlerName:            param.Handler,
	}

	rfc := resources.NewResourceBuilderContext(
		*param.ApplicationPrefix,
		*param.CommandPrefix,
	)

	cfContext := &CloudFormationDeployMode[interface{}]{
		Logger:                   common.NewDefaultLogger(),
		TemplateFileLocation:     param.TemplateFileName,
		Handler:                  param.Handler,
		ResourceFactoryContext:   rfc,
		lambdaFunctionDescriptor: lfd,
		DeploymentRegistrations:  resources.NewDeploymentBuilder(),
	}

	return cfContext
}

func generateLambdaDescriptor(
	lambdaFunctionDescriptor *lambdaFunctionDescriptor,
	rfc *resources.ResourceFactoryContext,
	env map[string]string,
	template *cloudformation.Template,
) {
	lambdaRoleKey := rfc.GenerateCommandResourceKey("iamrole", "role")
	lambdaRoleName := rfc.GenerateCommandResourceName("iamrole")

	role := &iam.Role{
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
	template.Resources[lambdaRoleKey] = role

	env["GADGETO_APPLICATION_NAME"] = *rfc.ApplicationName
	env["GADGETO_COMMAND_NAME"] = *rfc.CommandName

	runtime := "go1.x"
	template.Resources[*rfc.LambdaKey] = &lambda.Function{
		FunctionName: rfc.LambdaName,
		Code: &lambda.Function_Code{
			S3Bucket: lambdaFunctionDescriptor.deploymentSourceBucket,
			S3Key:    lambdaFunctionDescriptor.deployementSourceKey,
		},
		Environment: &lambda.Function_Environment{
			Variables: env,
		},
		Role:    cloudformation.GetAtt(lambdaRoleKey, "Arn"),
		Runtime: &runtime,
		Handler: lambdaFunctionDescriptor.handlerName,
		Tags: []tags.Tag{
			{
				Key:   "Foo",
				Value: "Bar",
			},
		},
	}
}

func (c *CloudFormationDeployMode[Client]) Complete() {
	template := cloudformation.NewTemplate()
	for key, resource := range c.DeploymentRegistrations.CloudformationResources {
		template.Resources[key] = resource
	}

	for _, hook := range c.DeploymentRegistrations.CompletionHooks {
		hook(c.ResourceFactoryContext, template)
	}
	generateLambdaDescriptor(c.lambdaFunctionDescriptor, c.ResourceFactoryContext, c.DeploymentRegistrations.EnvironmentVariables, template)

	yaml, err := template.YAML()
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
