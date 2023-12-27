package gs3

import (
	"fmt"
	"os"

	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/s3"
	"github.com/stefan79/gadgeto/pkg/modes"
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

func (c *S3Config) Connect() (Client, error) {
	envKey := generateResourceName("s3", *c.Name)
	bucketName, ok := os.LookupEnv(envKey)
	if !ok {
		return nil, fmt.Errorf("environment variable %s not set", envKey)
	}
	return NewS3Client(bucketName)
}

func (c *S3Config) Deploy(template *cloudformation.Template, env map[string]string) (Client, error) {
	resourceName := generateResourceName("s3", *c.Name)
	bucket := &s3.Bucket{}
	if c.BucketName != nil {
		bucket.BucketName = c.BucketName
	}
	template.Resources[resourceName] = bucket
	env[resourceName] = cloudformation.Ref(resourceName)

	return &NoOpClient{}, nil
}

func generateResourceName(rType string, name string) string {
	return fmt.Sprintf("%s%s", rType, name)
}

func (c *S3Config) Build() Client {
	client, err := c.BaseAWSResource.Mode.Dispatch(c)
	if err != nil {
		panic(err)
	}
	return client

}
