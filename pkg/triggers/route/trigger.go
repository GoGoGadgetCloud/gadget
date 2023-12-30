package route

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type (
	noopConf struct {
	}
	lambdaConf struct {
	}
)

func (r *noopConf) Handle(handler func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) {
	//do nothing
}

func (r *lambdaConf) Handle(handler func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) {
	lambda.Start(handler)
}
