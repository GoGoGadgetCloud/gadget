package bootstrap

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/stefan79/gadgeto/pkg/modes"
	"github.com/urfave/cli/v2"
)

type BootStrapContext struct {
	Context modes.Mode[interface{}]
}

func NewContext() modes.Mode[interface{}] {
	errorLogger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: false,
		TimeFormat:      time.Kitchen,
	})
	bs := PrepareBootstrapContext()
	if isLambdaRuntime() {
		fmt.Println("Detected a Lambda Runtime, using Lambda Runtime Context")
		return modes.NewRunMode()
	}
	app := Init(bs.InitCloudformationDeployment)
	if err := app.Run(os.Args); err != nil {
		errorLogger.Error("Failed to start in desired mode", "err", err)
	}
	if bs.Context == nil {
		os.Exit(1)
	}
	return bs.Context
}

func isLambdaRuntime() bool {
	_, ok := os.LookupEnv("AWS_EXECUTION_ENV")
	return ok
}

func lookup(key string, ctx *cli.Context) *string {

	if ctx.IsSet(key) {
		value := ctx.String(key)
		return &value
	}
	return nil
}

func (bs *BootStrapContext) InitCloudformationDeployment(cCtx *cli.Context) error {
	param := &modes.DeployModeParam{
		TemplateFileName:  lookup("template", cCtx),
		Handler:           lookup("handler", cCtx),
		S3Bucket:          lookup("s3bucket", cCtx),
		S3Key:             lookup("s3key", cCtx),
		ApplicationPrefix: lookup("application", cCtx),
		CommandPrefix:     lookup("command", cCtx),
	}
	bs.Context = modes.NewDeployMode(param)
	return nil
}

func PrepareBootstrapContext() *BootStrapContext {
	return &BootStrapContext{}
}
