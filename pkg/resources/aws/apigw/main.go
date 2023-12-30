package apigw

import (
	"github.com/aws/aws-lambda-go/events"

	"github.com/stefan79/gadgeto/pkg/modes"
	"github.com/stefan79/gadgeto/pkg/triggers"
)

var Ignore = BasePath("ignore")
var Prepend = BasePath("prepend")
var Split = BasePath("split")

type (
	BasePath string

	APIGatewayProxyTrigger triggers.Trigger[events.APIGatewayProxyRequest, events.APIGatewayProxyResponse]

	APIGatewayClient interface {
		GetTriggerMode() modes.Mode[APIGatewayProxyTrigger]
		GetReference() string
		GetKey() string
	}
)

func NewApiGatewayClient(name string, mode modes.Mode[interface{}]) ApiGatewayBuilder {
	return &apiGatewayBuilderConfig{
		name:                      name,
		mode:                      modes.DowncastMode[APIGatewayClient](mode),
		baseMode:                  mode,
		corsAllowHeaders:          []string{"*"},
		corsAllowMethods:          []string{"*"},
		corsAllowOrigins:          []string{"*"},
		disableExecuteAPIEndpoint: false,
		routeSelectionExpression:  nil,
		tags:                      map[string]string{},
	}
}
