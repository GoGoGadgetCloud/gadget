package apigw

import (
	"errors"
	"fmt"

	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/apigatewayv2"
	"github.com/stefan79/gadgeto/pkg/modes"
	"github.com/stefan79/gadgeto/pkg/resources"
)

// Connect implements resources.ResourceFactory.
func (agbc *apiGatewayBuilderConfig) Connect(ctx *resources.ResourceFactoryContext) (APIGatewayClient, error) {
	return &runConfiguration{
		mode: modes.DowncastMode[APIGatewayProxyTrigger](agbc.baseMode),
	}, nil
}

// Deploy implements resources.ResourceFactory.
func (agbc *apiGatewayBuilderConfig) Deploy(ctx *resources.ResourceFactoryContext, db resources.DeploymentBuilder) (APIGatewayClient, error) {
	apiGWKey := ctx.GenerateAppResourceKey(resources.Api, agbc.name)
	apiGWName := ctx.GenerateAppResourceName(agbc.name)
	apiGW := &apigatewayv2.Api{
		Name: &apiGWName,
		CorsConfiguration: &apigatewayv2.Api_Cors{
			AllowHeaders: agbc.corsAllowHeaders,
			AllowMethods: agbc.corsAllowMethods,
			AllowOrigins: agbc.corsAllowOrigins,
		},
		Description:              resources.StringPtr(fmt.Sprintf("Generated by Gadget to implement the trigger for %s:%s", *ctx.ApplicationName, *ctx.CommandName)),
		ProtocolType:             resources.StringPtr("HTTP"),
		RouteSelectionExpression: agbc.routeSelectionExpression,
		Tags:                     agbc.tags,
	}
	rcfrErr := db.RegisterCloudformationResource(apiGWKey, apiGW)

	complConf := &completionConfiguration{
		apiName:         apiGWName,
		apiReference:    cloudformation.Ref(apiGWKey),
		integrationKeys: agbc.integrationKeys,
	}

	rchErr := db.RegisterCompletionHook(complConf.deploymentCompletion)

	return &buildConfiguration{
		mode:        modes.DowncastMode[APIGatewayProxyTrigger](agbc.baseMode),
		reference:   cloudformation.Ref(apiGWKey),
		key:         apiGWKey,
		builderConf: agbc,
	}, errors.Join(rcfrErr, rchErr)
}

type completionConfiguration struct {
	apiName         string
	apiReference    string
	integrationKeys []string
}

func (cc *completionConfiguration) deploymentCompletion(ctx *resources.ResourceFactoryContext, tmpl *cloudformation.Template) error {
	apiGWDeploymentKey := ctx.GenerateAppResourceKey(resources.Deployment, cc.apiName+"Deployment")
	apiGWDeployment := &apigatewayv2.Deployment{
		AWSCloudFormationDependsOn: cc.integrationKeys,
		ApiId:                      cc.apiReference,
		Description:                resources.StringPtr(fmt.Sprintf("Generated by Gadget to implement the trigger for %s:%s", *ctx.ApplicationName, *ctx.CommandName)),
	}
	tmpl.Resources[apiGWDeploymentKey] = apiGWDeployment

	apiGWStageKey := ctx.GenerateAppResourceKey(resources.Stage, cc.apiName+"Stage")
	apiGWStageName := ctx.GenerateAppResourceName(cc.apiName + "Stage")
	apiGWStage := &apigatewayv2.Stage{
		StageName:    apiGWStageName,
		Description:  resources.StringPtr(fmt.Sprintf("Generated by Gadget to implement the trigger for %s:%s", *ctx.ApplicationName, *ctx.CommandName)),
		ApiId:        cc.apiReference,
		DeploymentId: cloudformation.RefPtr(apiGWDeploymentKey),
	}
	tmpl.Resources[apiGWStageKey] = apiGWStage

	return nil

}
