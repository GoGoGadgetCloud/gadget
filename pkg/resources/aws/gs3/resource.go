package gs3

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type (
	noopConf struct {
	}

	runConf struct {
		Client s3.Client
		Bucket string
	}
)

func createS3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	return client, nil
}

func (c *runConf) WriteToObject(key string, data []byte) error {
	_, err := c.Client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: &c.Bucket,
		Key:    &key,
		Body:   bytes.NewReader(data)})
	return err
}

func newRunClient(bucketName string) (Client, error) {
	client, err := createS3Client(context.Background())
	if err != nil {
		return nil, err
	}
	return &runConf{
		Client: *client,
		Bucket: bucketName,
	}, nil

}

func (c *noopConf) WriteToObject(key string, data []byte) error {
	return fmt.Errorf("NoOpClient")
}
