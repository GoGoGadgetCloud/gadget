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
	ResourceFactoryContext struct {
		ApplicationName *string
		CommandName     *string
		LambdaRef       *string
		LambdaArn       *string
		LambdaKey       *string
	}

	ResourceFactory[Client any] interface {
		Deploy(ctx *ResourceFactoryContext, tmpl *cloudformation.Template, env map[string]string) (Client, error)
		Connect(ctx *ResourceFactoryContext) (Client, error)
	}
)

func NewResourceBuilderContext(applicationName string, commandName string) *ResourceFactoryContext {
	return &ResourceFactoryContext{
		ApplicationName: &applicationName,
		CommandName:     &commandName,
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
