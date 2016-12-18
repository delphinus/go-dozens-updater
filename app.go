package godo

import (
	"github.com/urfave/cli"
)

// NewApp will return godo App
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Usage = "Manage DNS with dozens"
	app.Version = Version
	app.Author = "delphinus"
	app.Email = "delphinus@remora.cx"
	app.Commands = []cli.Command{
		{
			Name:    "zone",
			Aliases: []string{"z"},
			Usage:   "Manage zones",
			Subcommands: []cli.Command{
				commandZoneList,
				commandZoneCreate,
			},
		},
	}
	app.Before = func(c *cli.Context) error {
		if c.Command.Name != "" {
			SetupConfig()
		}
		return nil
	}
	return app
}
