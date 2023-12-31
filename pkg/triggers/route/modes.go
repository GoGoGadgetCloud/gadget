package route

import (
	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/apigatewayv2"
	"github.com/awslabs/goformation/v7/cloudformation/iam"
	"github.com/stefan79/gadgeto/internal/iam_generator"
	"github.com/stefan79/gadgeto/pkg/resources"
	"github.com/stefan79/gadgeto/pkg/resources/aws/apigw"
)

// Connect implements resources.ResourceFactory.
func (*routeTriggerBuilderConf) Connect(ctx *resources.ResourceFactoryContext) (apigw.APIGatewayProxyTrigger, error) {
	return &lambdaConf{}, nil
}

// Deploy implements resources.ResourceFactory.
func (rtbc *routeTriggerBuilderConf) Deploy(ctx *resources.ResourceFactoryContext, tmpl *cloudformation.Template, env map[string]string) (apigw.APIGatewayProxyTrigger, resources.CompletionHook, error) {
	roleName := ctx.GenerateCommandResourceName(*rtbc.name + "Role")
	roleKey := ctx.GenerateCommandResourceKey(resources.IamRole, *rtbc.name+"Role")

	role := &iam.Role{
		AssumeRolePolicyDocument: iam_generator.GenerateAssumeRoleForService(iam_generator.ApiGateway),
		Path:                     resources.StringPtr("/"),
		Policies: []iam.Role_Policy{

			{
				PolicyName:     roleName,
				PolicyDocument: iam_generator.GeneratePermissionForReource(iam_generator.Allow, iam_generator.AllLambdaActions, *ctx.LambdaArn),
			},
		},
	}
	tmpl.Resources[roleKey] = role

	routeIntegrationKey := ctx.GenerateCommandResourceKey(resources.Integration, *rtbc.name+"Integration")
	timeoutInMillis := int(29000)
	routeIntegration := &apigatewayv2.Integration{
		ApiId:                rtbc.gateway.GetReference(),
		Description:          resources.StringPtr("Integration for " + *rtbc.name),
		ConnectionType:       resources.StringPtr("INTERNET"),
		CredentialsArn:       cloudformation.GetAttPtr(roleKey, "Arn"),
		PassthroughBehavior:  resources.StringPtr("WHEN_NO_MATCH"),
		TimeoutInMillis:      &timeoutInMillis,
		IntegrationMethod:    (*string)(rtbc.method),
		IntegrationType:      "AWS_PROXY",
		PayloadFormatVersion: resources.StringPtr("2.0"),
		IntegrationUri:       ctx.LambdaArn,
	}
	tmpl.Resources[routeIntegrationKey] = routeIntegration

	routeName := ctx.GenerateCommandResourceName(*rtbc.name)
	routeKey := ctx.GenerateCommandResourceKey(resources.Route, *rtbc.name)
	route := &apigatewayv2.Route{
		AWSCloudFormationDependsOn: []string{
			rtbc.gateway.GetKey(),
			*ctx.LambdaKey,
			routeIntegrationKey,
		},
		AuthorizationType: resources.StringPtr("NONE"),
		ApiId:             rtbc.gateway.GetReference(),
		OperationName:     &routeName,
		RouteKey:          *rtbc.routeKey,
		Target: cloudformation.JoinPtr(
			"/",
			[]string{
				"integrations",
				cloudformation.Ref(routeIntegrationKey),
			}),
	}
	tmpl.Resources[routeKey] = route
	return &noopConf{}, nil, nil
}
