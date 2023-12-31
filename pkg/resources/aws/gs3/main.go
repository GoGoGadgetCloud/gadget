package gs3

import (
	"github.com/stefan79/gadgeto/pkg/modes"
	"github.com/stefan79/gadgeto/pkg/resources/aws"
)

type (
	Client interface {
		WriteToObject(string, []byte) error
	}
)

func S3(mode modes.Mode[interface{}], name string) S3Builder {
	return &builderConf{
		BaseAWSResource: aws.BaseAWSResource[Client]{
			Name: &name,
			Mode: modes.DowncastMode[Client](mode),
		},
	}
}
