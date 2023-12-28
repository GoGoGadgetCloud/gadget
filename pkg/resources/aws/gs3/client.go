package gs3

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ClientConfig struct {
	Client s3.Client
	Bucket string
}

func NewS3Client(bucketName string) (Client, error) {
	client, err := createS3Client(context.Background())
	if err != nil {
		return nil, err
	}
	return &ClientConfig{
		Client: *client,
		Bucket: bucketName,
	}, nil

}

func createS3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	return client, nil
}

func (c *ClientConfig) WriteToObject(key string, data []byte) error {
	_, err := c.Client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: &c.Bucket,
		Key:    &key,
		Body:   bytes.NewReader(data)})
	return err
}
