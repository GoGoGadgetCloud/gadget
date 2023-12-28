package modes

import (
	"github.com/stefan79/gadgeto/pkg/resources"
)

type Mode[Resource any] interface {
	Complete()
	Dispatch(factory resources.ResourceFactory[Resource]) (Resource, error)
}
