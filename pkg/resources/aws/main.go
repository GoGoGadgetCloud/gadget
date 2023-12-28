package aws

import "github.com/stefan79/gadgeto/pkg/modes"

type BaseAWSResource[Client any] struct {
	Name *string
	Mode modes.Mode[Client]
}
