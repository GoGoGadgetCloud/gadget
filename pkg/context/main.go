package context

import (
	"github.com/stefan79/gadgeto/internal/resources"
)

type GadgetoContext[Client any] interface {
	Complete() error
	Dispatch(factory resources.ResourceFactory[Client]) (Client, error)
}
