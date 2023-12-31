package gs3

import (
	"github.com/stefan79/gadgeto/pkg/resources/aws"
)

type (
	builderConf struct {
		aws.BaseAWSResource[Client]
		bucketName *string
	}

	S3Builder interface {
		WithBucketName(string) S3Builder
		Build() Client
	}
)

func (c *builderConf) WithBucketName(bucketName string) S3Builder {
	c.bucketName = &bucketName
	return c
}

func (c *builderConf) Build() Client {
	client, err := c.BaseAWSResource.Mode.Dispatch(c)
	if err != nil {
		panic(err)
	}
	return client

}
