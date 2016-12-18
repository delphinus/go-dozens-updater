package godo

import (
	"encoding/json"
	"os"

	"github.com/delphinus/godo/lib/dozens"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var commandZoneCreate = cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Usage:   "Create a new zone",
	Action:  doZoneCreate,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Usage: "Name of the zone (required)",
		},
		cli.BoolFlag{
			Name:  "add_google_apps, a",
			Usage: "Set if google apps is needed",
		},
		cli.StringFlag{
			Name:  "google_authorize, g",
			Value: "",
			Usage: "Set if TXT record is needed to confirm Google Apps",
		},
		cli.StringFlag{
			Name:  "mailaddress, m",
			Value: "",
			Usage: "SOA mail address. It will use the default address, if empty.",
		},
	},
}

func doZoneCreate(c *cli.Context) error {
	name := c.String("name")
	if name == "" {
		return errors.New("Please specify name of zone")
	}

	body := dozens.ZoneCreateBody{
		Name:            name,
		AddGoogleApps:   c.Bool("add_google_apps"),
		GoogleAuthorize: c.String("google_authorize"),
		MailAddress:     c.String("mailaddress"),
	}

	zone, err := dozens.ZoneCreate(body)
	if err != nil {
		return errors.Wrap(err, "error in PostCreate")
	}

	if err := json.NewEncoder(os.Stdout).Encode(zone); err != nil {
		return errors.Wrap(err, "error in Encode")
	}

	return nil
}
