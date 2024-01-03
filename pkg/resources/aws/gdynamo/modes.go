package gdynamo

import (
	"github.com/stefan79/gadgeto/pkg/resources"
)

// Connect implements resources.ResourceFactory.
func (*builderConf[ID, Type]) Connect(ctx *resources.ResourceFactoryContext) (DynamoDBResource[ID, Type], error) {
	panic("unimplemented")
}

// Deploy implements resources.ResourceFactory.
func (*builderConf[ID, Type]) Deploy(ctx *resources.ResourceFactoryContext, db resources.DeploymentBuilder) (DynamoDBResource[ID, Type], error) {
	return nil, nil
}
