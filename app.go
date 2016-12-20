package godo

import (
	"github.com/urfave/cli"
)

// NewApp will return godo App
func NewApp() *cli.App {
	before := func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			return nil
		}
		return SetupConfig()
	}

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
			Before:  before,
			Subcommands: []cli.Command{
				commandZoneList,
				commandZoneCreate,
				commandZoneUpdate,
				commandZoneDelete,
			},
		},
		{
			Name:    "record",
			Aliases: []string{"r"},
			Usage:   "Manage records",
			Before:  before,
			Subcommands: []cli.Command{
				commandRecordList,
				commandRecordDelete,
				commandRecordUpdate,
			},
		},
	}
	return app
}
