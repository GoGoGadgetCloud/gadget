package apigw

import "github.com/stefan79/gadgeto/pkg/modes"

type (
	apiGatewayBuilderConfig struct {
		mode                      modes.Mode[APIGatewayClient]
		baseMode                  modes.Mode[interface{}]
		name                      string
		method                    string
		corsAllowHeaders          []string
		corsAllowMethods          []string
		corsAllowOrigins          []string
		disableExecuteAPIEndpoint bool
		routeSelectionExpression  *string
		tags                      map[string]string
		integrationKeys           []string
	}

	ApiGatewayBuilder interface {
		WithMethod(string) ApiGatewayBuilder
		WithCORSAllowHeaders(headers ...string) ApiGatewayBuilder
		WithCORSAllowMethods(methods ...string) ApiGatewayBuilder
		WithCORSAllowOrigins(origins ...string) ApiGatewayBuilder
		WithDisbaledExecuteAPIEndpoint() ApiGatewayBuilder
		WithRouteSelectionExpression(string) ApiGatewayBuilder
		AddTag(key string, value string) ApiGatewayBuilder
		Build() APIGatewayClient
	}
)

func (c *apiGatewayBuilderConfig) WithMethod(method string) ApiGatewayBuilder {
	c.method = method
	return c
}

func (c *apiGatewayBuilderConfig) WithCORSAllowHeaders(headers ...string) ApiGatewayBuilder {
	c.corsAllowHeaders = headers
	return c
}

func (c *apiGatewayBuilderConfig) WithCORSAllowMethods(methods ...string) ApiGatewayBuilder {
	c.corsAllowMethods = methods
	return c
}

func (c *apiGatewayBuilderConfig) WithCORSAllowOrigins(origins ...string) ApiGatewayBuilder {
	c.corsAllowOrigins = origins
	return c
}

func (c *apiGatewayBuilderConfig) WithDisbaledExecuteAPIEndpoint() ApiGatewayBuilder {
	c.disableExecuteAPIEndpoint = true
	return c
}

func (c *apiGatewayBuilderConfig) WithRouteSelectionExpression(expression string) ApiGatewayBuilder {
	c.routeSelectionExpression = &expression
	return c
}

func (c *apiGatewayBuilderConfig) AddTag(key string, value string) ApiGatewayBuilder {
	c.tags[key] = value
	return c
}

func (c *apiGatewayBuilderConfig) Build() APIGatewayClient {
	trigger, err := c.mode.Dispatch(c)
	if err != nil {
		panic(err)
	}
	return trigger
}
