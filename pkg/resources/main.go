package resources

import (
	"strings"

	"github.com/awslabs/goformation/v7/cloudformation"
)

type ResourceType string

var LambdaFunction = ResourceType("AWSLambdaFunction")
var IamRole = ResourceType("AWSIAMRole")
var S3Bucket = ResourceType("AWSS3Bucket")
var Api = ResourceType("AWSApiGatewayV2Api")
var Route = ResourceType("AWSApiGatewayV2Route")
var Stage = ResourceType("AWSApiGatewayV2Stage")
var Deployment = ResourceType("AWSApiGatewayV2Deployment")
var Integration = ResourceType("AWSApiGatewayV2Integration")

type (
	CompletionHook func(*ResourceFactoryContext, *cloudformation.Template) error

	DeploymentRegistrations struct {
		CloudformationResources map[string]cloudformation.Resource
		EnvironmentVariables    map[string]string
		LambdaRoleActions       []map[string]interface{}
		CompletionHooks         []CompletionHook
	}

	ResourceFactoryContext struct {
		ApplicationName *string
		CommandName     *string
		LambdaName      *string
		LambdaKey       *string
		LambdaArn       *string
		LambdaRef       *string
	}

	DeploymentBuilder interface {
		RegisterCloudformationResource(string, cloudformation.Resource) error
		RegisterEnvironmentVariable(string, string) error
		RegisterLabmdaRoleAction(map[string]interface{}) error
		RegisterCompletionHook(CompletionHook) error
	}

	ResourceFactory[Resource any] interface {
		Deploy(*ResourceFactoryContext, DeploymentBuilder) (Resource, error)
		Connect(*ResourceFactoryContext) (Resource, error)
	}
)

func NewResourceBuilderContext(applicationName string, commandName string) *ResourceFactoryContext {
	lambdaKey := strings.Title(applicationName) + strings.Title(commandName)
	lambdaName := strings.Title(applicationName) + strings.Title(commandName) + strings.Title(string(LambdaFunction))
	return &ResourceFactoryContext{
		ApplicationName: &applicationName,
		CommandName:     &commandName,
		LambdaName:      &lambdaName,
		LambdaKey:       &lambdaKey,
		LambdaArn:       cloudformation.GetAttPtr(lambdaKey, "Arn"),
		LambdaRef:       cloudformation.RefPtr(lambdaKey),
	}
}

func (p *ResourceFactoryContext) GenerateCommandResourceName(resName string) string {
	// Convert each part to PascalCase
	appName := strings.Title(*p.ApplicationName)
	cmdName := strings.Title(*p.CommandName)
	resName = strings.Title(resName)

	// Concatenate the parts
	resourceName := appName + cmdName + resName

	return resourceName
}

func (p *ResourceFactoryContext) GenerateCommandResourceKey(resType ResourceType, resName string) string {
	// Convert each part to PascalCase

	resTypeAsString := strings.Title(string(resType))
	appName := strings.Title(*p.ApplicationName)
	cmdName := strings.Title(*p.CommandName)
	resName = strings.Title(resName)

	// Concatenate the parts
	resourceName := resTypeAsString + appName + cmdName + resName

	return resourceName
}

func (p *ResourceFactoryContext) GenerateAppResourceName(resName string) string {
	// Convert each part to PascalCase
	appName := strings.Title(*p.ApplicationName)
	resName = strings.Title(resName)

	// Concatenate the parts
	resourceName := appName + resName

	return resourceName
}

func (p *ResourceFactoryContext) GenerateAppResourceKey(resType ResourceType, resName string) string {
	// Convert each part to PascalCase
	resTypeAsString := strings.Title(string(resType))
	appName := strings.Title(*p.ApplicationName)
	resName = strings.Title(resName)

	// Concatenate the parts
	resourceName := resTypeAsString + appName + resName

	return resourceName
}

func StringPtr(input string) *string {
	return &input
}

func NewDeploymentBuilder() *DeploymentRegistrations {
	return &DeploymentRegistrations{
		CloudformationResources: make(map[string]cloudformation.Resource),
		EnvironmentVariables:    make(map[string]string),
		LambdaRoleActions:       make([]map[string]interface{}, 0),
		CompletionHooks:         make([]CompletionHook, 0),
	}
}

func (dr *DeploymentRegistrations) RegisterCloudformationResource(key string, resource cloudformation.Resource) error {
	dr.CloudformationResources[key] = resource
	return nil
}

func (dr *DeploymentRegistrations) RegisterEnvironmentVariable(key string, value string) error {
	dr.EnvironmentVariables[key] = value
	return nil
}

func (dr *DeploymentRegistrations) RegisterLabmdaRoleAction(action map[string]interface{}) error {
	dr.LambdaRoleActions = append(dr.LambdaRoleActions, action)
	return nil
}

func (dr *DeploymentRegistrations) RegisterCompletionHook(hook CompletionHook) error {
	dr.CompletionHooks = append(dr.CompletionHooks, hook)
	return nil
}
