package godo

import (
	"encoding/json"
	"os"

	"github.com/delphinus/godo/lib/dozens"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var commandZoneUpdate = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Usage:   "Update SOA of the specified zone",
	Action:  doZoneUpdate,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "id, i",
			Usage: "The zone ID to update",
		},
		cli.StringFlag{
			Name:  "soa, s",
			Usage: "The SOA mailaddress",
		},
	},
}

func doZoneUpdate(c *cli.Context) error {
	zoneID := c.String("id")
	mailAddress := c.String("soa")
	if zoneID == "" || mailAddress == "" {
		return errors.New("Please specify zone ID & SOA mail address")
	}

	zone, err := dozens.ZoneUpdate(Config.Token, zoneID, mailAddress)
	if err != nil {
		return errors.Wrap(err, "errorin ZoneUpdate")
	}

	if err := json.NewEncoder(os.Stdout).Encode(zone); err != nil {
		return errors.Wrap(err, "error in Encode")
	}
	return nil
}
