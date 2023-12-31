package apigw

import "github.com/stefan79/gadgeto/pkg/modes"

type (
	runConfiguration struct {
		mode modes.Mode[APIGatewayProxyTrigger]
	}

	buildConfiguration struct {
		mode        modes.Mode[APIGatewayProxyTrigger]
		reference   string
		key         string
		builderConf *apiGatewayBuilderConfig
	}
)

func (bc *buildConfiguration) GetTriggerMode() modes.Mode[APIGatewayProxyTrigger] {
	return bc.mode
}

func (bc *buildConfiguration) GetReference() string {
	return bc.reference
}

func (bc *buildConfiguration) GetKey() string {
	return bc.key
}

func (bc *buildConfiguration) RegisterIntegration(key string) {
	bc.builderConf.integrationKeys = append(bc.builderConf.integrationKeys, key)
}

func (nc *runConfiguration) GetTriggerMode() modes.Mode[APIGatewayProxyTrigger] {
	return nc.mode
}

func (nc *runConfiguration) GetReference() string {
	panic("not implemented")
}

func (nc *runConfiguration) GetKey() string {
	panic("not implemented")
}

func (nc *runConfiguration) RegisterIntegration(key string) {
	panic("not implemented")
}
