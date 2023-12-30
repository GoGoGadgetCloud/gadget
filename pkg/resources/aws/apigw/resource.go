package apigw

import "github.com/stefan79/gadgeto/pkg/modes"

type (
	runConf struct {
		mode modes.Mode[APIGatewayProxyTrigger]
	}

	buildConf struct {
		mode      modes.Mode[APIGatewayProxyTrigger]
		reference string
		key       string
	}
)

func (dc *buildConf) GetTriggerMode() modes.Mode[APIGatewayProxyTrigger] {
	return dc.mode
}

func (dc *buildConf) GetReference() string {
	return dc.reference
}

func (dc *buildConf) GetKey() string {
	return dc.key
}

func (nc *runConf) GetTriggerMode() modes.Mode[APIGatewayProxyTrigger] {
	return nc.mode
}

func (nc *runConf) GetReference() string {
	panic("not implemented")
}

func (nc *runConf) GetKey() string {
	panic("not implemented")
}
