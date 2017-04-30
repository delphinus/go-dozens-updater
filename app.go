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
				commandRecordCreate,
				commandRecordDelete,
				commandRecordUpdate,
			},
		},
		{
			Name:   "renew",
			Usage:  "Renew record entry if needed",
			Before: before,
			Action: renew,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "zone,z",
					Usage: "Zone name to update",
				},
				cli.StringFlag{
					Name:  "domain, d",
					Usage: "Domain to update IP Address",
				},
				cli.BoolFlag{
					Name:  "show-ip-only, s",
					Usage: "Show current IP address setting",
				},
				cli.BoolFlag{
					Name:  "ipv6, 6",
					Usage: "Use ipv6 instead of ipv4",
				},
			},
		},
	}
	return app
}
