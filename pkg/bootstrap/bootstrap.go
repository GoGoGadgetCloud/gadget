package bootstrap

import (
	"fmt"
	"os"

	"github.com/stefan79/gadgeto/pkg/modes"
	"github.com/urfave/cli/v2"
)

type BootStrapContext struct {
	Context modes.Mode[interface{}]
}

func NewContext() modes.Mode[interface{}] {
	bs := PrepareBootstrapContext()
	if isLambdaRuntime() {
		fmt.Println("Detected a Lambda Runtime, using Lambda Runtime Context")
		return modes.NewRunMode()
	}
	app := Init(bs.InitCloudformationDeployment)
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
	if bs.Context == nil {
		panic("no context initialized")
	}
	return bs.Context
}

func isLambdaRuntime() bool {
	_, ok := os.LookupEnv("AWS_EXECUTION_ENV")
	return ok
}

func (bs *BootStrapContext) InitCloudformationDeployment(cCtx *cli.Context) error {
	output := cCtx.String("output")
	bucket := cCtx.String("bucket")
	bs.Context = modes.NewDeployMode(&output, &bucket)
	return nil
}

func PrepareBootstrapContext() *BootStrapContext {
	return &BootStrapContext{}
}
