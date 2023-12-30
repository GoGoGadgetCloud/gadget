package bootstrap

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func Init(deployCmd func(*cli.Context) error) *cli.App {
	name := filepath.Base(os.Args[0])
	app := &cli.App{
		Name:        name,
		Usage:       "A gadget enriched application",
		Description: "If not run in a Lambda Runtime, you must configure the run mode",
		Commands: []*cli.Command{
			{
				Name:      "deployment",
				Usage:     "manages cloudformation / gadget templates",
				UsageText: "Use the subcommands to generate the cloudformation and gadget templates ",
				Aliases:   []string{"d"},
				Subcommands: []*cli.Command{
					{
						Name:    "generate",
						Aliases: []string{"g"},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "template",
								Usage:    "where the template should be generated to",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "handler",
								Usage:    "handler to use in the function definition",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "application",
								Usage:    "application prefix to use in the created resources",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "command",
								Usage:    "command prefix to use in the created resources",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "s3bucket",
								Usage:    "s3 bucket name to use in the function source",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "s3key",
								Usage:    "s3 key to use in the function source",
								Required: true,
							},
						},
						Action: deployCmd,
					},
				},
			},
		},
	}

	return app

}
