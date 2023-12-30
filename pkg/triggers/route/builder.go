package route

import (
	"github.com/stefan79/gadgeto/pkg/resources"
	"github.com/stefan79/gadgeto/pkg/resources/aws/apigw"
)

type (
	routeTriggerBuilderConf struct {
		name     *string
		gateway  apigw.APIGatewayClient
		routeKey *string
		method   *Method
	}

	//TODO: Add support for AWS_IAM using a Resource
	//TODO: Add support for AuthorizationType JSW sing a Resource
	//TODO: Add support for AuthorizationType Custom Lambda using a Resource
	//TODO: Add support for WebSockets
	//TODO: Add support for encoded payloads
	//TODO: Add support for RequestParameters
	RouteTriggerBuilder interface {
		WithKey(method Method, key string) RouteTriggerBuilder
		WithDefaultKey() RouteTriggerBuilder
		Build() apigw.APIGatewayProxyTrigger
	}
)

func (r *routeTriggerBuilderConf) WithKey(method Method, key string) RouteTriggerBuilder {
	r.method = &method
	r.routeKey = resources.StringPtr(string(method) + " " + key)
	return r
}

func (r *routeTriggerBuilderConf) WithDefaultKey() RouteTriggerBuilder {
	r.method = &ANY
	r.routeKey = resources.StringPtr("$default")
	return r
}

func (r *routeTriggerBuilderConf) Build() apigw.APIGatewayProxyTrigger {
	trigger, err := r.gateway.GetTriggerMode().Dispatch(r)
	if err != nil {
		panic(err)
	}
	return trigger
}
