package gs3

import (
	"errors"
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

func (c *builderConf) Deploy(ctx *resources.ResourceFactoryContext, db resources.DeploymentBuilder) (Client, error) {
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
	rcrErr := db.RegisterCloudformationResource(bucketKey, bucket)
	revErr := db.RegisterEnvironmentVariable(bucketKey, cloudformation.Ref(bucketKey))

	return &noopConf{}, errors.Join(rcrErr, revErr)
}
