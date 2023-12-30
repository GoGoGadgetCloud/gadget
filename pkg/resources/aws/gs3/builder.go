package gs3

import (
	"fmt"
	"os"

	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/s3"
	"github.com/awslabs/goformation/v7/cloudformation/tags"
	"github.com/stefan79/gadgeto/pkg/modes"
	"github.com/stefan79/gadgeto/pkg/resources"
	"github.com/stefan79/gadgeto/pkg/resources/aws"
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

func S3(mode modes.Mode[interface{}], name string) S3Builder {
	return &S3Config{
		BaseAWSResource: aws.BaseAWSResource[Client]{
			Name: &name,
			Mode: modes.DowncastMode[Client](mode),
		},
	}
}

func (c *S3Config) WithBucketName(bucketName string) S3Builder {
	c.BucketName = &bucketName
	return c
}

func (c *S3Config) Connect(ctx *resources.ResourceFactoryContext) (Client, error) {
	envKey := ctx.GenerateAppResourceKey(resources.S3Bucket, *c.BucketName)
	bucketName, ok := os.LookupEnv(envKey)
	if !ok {
		return nil, fmt.Errorf("environment variable %s not set", envKey)
	}
	return NewS3Client(bucketName)
}

func (c *S3Config) Deploy(ctx *resources.ResourceFactoryContext, template *cloudformation.Template, env map[string]string) (Client, error) {
	bucketKey := ctx.GenerateAppResourceKey(resources.S3Bucket, *c.BucketName)

	bucket := &s3.Bucket{
		Tags: []tags.Tag{
			{
				Key:   "Foo",
				Value: "Bar",
			},
		},
	}
	if c.BucketName != nil {
		bucket.BucketName = c.BucketName
	}
	template.Resources[bucketKey] = bucket
	env[bucketKey] = cloudformation.Ref(bucketKey)

	return &NoOpClient{}, nil
}

func (c *S3Config) Build() Client {
	client, err := c.BaseAWSResource.Mode.Dispatch(c)
	if err != nil {
		panic(err)
	}
	return client

}
