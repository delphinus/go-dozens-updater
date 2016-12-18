package godo

import (
	"encoding/json"
	"os"

	"github.com/delphinus/godo/lib/dozens"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var commandZoneDelete = cli.Command{
	Name:    "delete",
	Aliases: []string{"d"},
	Usage:   "Delete the specified zone",
	Action:  doZoneDelete,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "id, i",
			Usage: "The zone ID to delete",
		},
	},
}

func doZoneDelete(c *cli.Context) error {
	zoneID := c.String("id")
	if zoneID == "" {
		return errors.New("Please specify zone ID")
	}

	zone, err := dozens.ZoneDelete(Config.Token, zoneID)
	if err != nil {
		return errors.Wrap(err, "errorin ZoneDelete")
	}

	if err := json.NewEncoder(os.Stdout).Encode(zone); err != nil {
		return errors.Wrap(err, "error in Encode")
	}
	return nil
}
