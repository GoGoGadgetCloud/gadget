package gs3

import (
	"fmt"
	"os"

	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/s3"
	"github.com/awslabs/goformation/v7/cloudformation/tags"
	"github.com/stefan79/gadgeto/pkg/resources"
)

func (c *builderConf) Connect(ctx *resources.ResourceFactoryContext) (Client, error) {
	envKey := ctx.GenerateAppResourceKey(resources.S3Bucket, *c.bucketName)
	bucketName, ok := os.LookupEnv(envKey)
	if !ok {
		return nil, fmt.Errorf("environment variable %s not set", envKey)
	}
	return newRunClient(bucketName)
}

func (c *builderConf) Deploy(ctx *resources.ResourceFactoryContext, template *cloudformation.Template, env map[string]string) (Client, resources.CompletionHook, error) {
	bucketKey := ctx.GenerateAppResourceKey(resources.S3Bucket, *c.bucketName)

	bucket := &s3.Bucket{
		Tags: []tags.Tag{
			{
				Key:   "Foo",
				Value: "Bar",
			},
		},
	}
	if c.bucketName != nil {
		bucket.BucketName = c.bucketName
	}
	template.Resources[bucketKey] = bucket
	env[bucketKey] = cloudformation.Ref(bucketKey)

	return &noopConf{}, nil, nil
}
