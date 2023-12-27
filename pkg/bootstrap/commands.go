package bootstrap

import (
	"github.com/urfave/cli/v2"
)

func Init(deployCmd func(*cli.Context) error) *cli.App {

	app := &cli.App{
		Name: "gadgeto compute",
		Commands: []*cli.Command{
			{
				Name:    "deploy",
				Aliases: []string{"d"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "output",
						Aliases:  []string{"o"},
						Usage:    "output file location",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "bucket",
						Aliases:  []string{"b"},
						Usage:    "s3 bucket",
						Required: true,
					},
				},
				Action: deployCmd,
			},
		},
	}

	return app

}
