package gs3

import (
	"fmt"

	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/s3"
	"github.com/stefan79/gadgeto/internal/resources/aws"
	"github.com/stefan79/gadgeto/internal/resources/aws/gs3/internal"
	"github.com/stefan79/gadgeto/pkg/context"
)

type (
	S3Builder interface {
		WithBucketName(string) S3Builder
		Build() Client
	}

	S3Config struct {
		aws.BaseAWSResource[Client]
		BucketName *string
	}
)

func S3(ctx context.GagdgetoContext[interface{}], name string) S3Builder {
	return &S3Config{
		BaseAWSResource: aws.BaseAWSResource[Client]{
			Name:           &name,
			GadgetoContext: context.WithContext[Client](ctx),
		},
	}
}

func (c *S3Config) WithBucketName(bucketName string) S3Builder {
	c.BucketName = &bucketName
	return c
}

func (c *S3Config) Connect() (Client, error) {
	return &internal.NoOpClient{}, nil
}

func (c *S3Config) Deploy(template cloudformation.Template) (Client, error) {
	template.Resources[generateResourceName("AWS::S3::Bucket", *c.Name)] = &s3.Bucket{
		BucketName: c.BucketName,
	}
	return &internal.NoOpClient{}, nil
}

func generateResourceName(rType string, name string) string {
	return fmt.Sprintf("%s.%s", rType, name)
}

func (c *S3Config) Build() Client {
	client, err := c.BaseAWSResource.GadgetoContext.Dispatch(c)
	if err != nil {
		panic(err)
	}
	return client

}
