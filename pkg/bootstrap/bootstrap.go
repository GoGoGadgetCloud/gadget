package bootstrap

import (
	"fmt"
	"os"

	"github.com/stefan79/gadgeto/pkg/context"
	"github.com/stefan79/gadgeto/pkg/context/impl"
	"github.com/urfave/cli/v2"
)

type BootStrapContext struct {
	Context context.GadgetoContext[interface{}]
}

func NewContext() context.GadgetoContext[interface{}] {
	bs := PrepareBootstrapContext()
	if isLambdaRuntime() {
		fmt.Println("Detected a Lambda Runtime, using Lambda Runtime Context")
		return impl.NewRunContext()
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
	bs.Context = impl.NewDeploymentContext(&output, &bucket)
	return nil
}

func PrepareBootstrapContext() *BootStrapContext {
	return &BootStrapContext{}
}
