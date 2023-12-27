package aws

import "github.com/stefan79/gadgeto/pkg/context"

type BaseAWSResource[Client any] struct {
	Name           *string
	GadgetoContext context.GadgetoContext[Client]
}
