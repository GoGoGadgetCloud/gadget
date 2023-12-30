package route

import (
	"github.com/stefan79/gadgeto/pkg/resources"
	"github.com/stefan79/gadgeto/pkg/resources/aws/apigw"
)

var GET Method = "GET"
var POST Method = "POST"
var PUT Method = "PUT"
var DELETE Method = "DELETE"
var PATCH Method = "PATCH"
var HEAD Method = "HEAD"
var OPTIONS Method = "OPTIONS"
var ANY Method = "*"

type (
	Method string
)

func NewTrigger(name string, gateway apigw.APIGatewayClient) RouteTriggerBuilder {
	return &routeTriggerBuilderConf{
		name:     &name,
		gateway:  gateway,
		routeKey: resources.StringPtr("$default"),
	}
}
