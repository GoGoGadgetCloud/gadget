package main

import (
	"github.com/stefan79/gadgeto/internal/resources/aws/gs3"
	apigw "github.com/stefan79/gadgeto/internal/triggers"
	"github.com/stefan79/gadgeto/pkg/context"
)

type (
	Setup struct {
		MyS3Bucket gs3.Client
	}
)

func (s *Setup) handleCreateCustomerCall(request apigw.Request, response apigw.Response) error {
	key, ok := request.QueryParams["key"]
	var err error
	if !ok {
		response.ResponseCode = 400
	} else {
		err = s.MyS3Bucket.WriteToObject(key, request.Body)
	}
	return err
}

func main() {
	ctx := context.NewGadgetoContext()

	setup := &Setup{
		MyS3Bucket: gs3.S3(ctx, "myBucket").WithBucketName("my-bucket").Build(),
	}
	apigw.ApiGateway("CreateCustomer").WithMethod("POST").Build(ctx).Handle(setup.handleCreateCustomerCall)
	err := ctx.Complete()
	if err != nil {
		panic(err)
	}
}
