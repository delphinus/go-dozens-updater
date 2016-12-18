package godo

import (
	"encoding/json"
	"os"

	"github.com/delphinus/godo/lib/dozens"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var commandZones = cli.Command{
	Name:    "zones",
	Aliases: []string{"z"},
	Usage:   "List up zones in dozens",
	Action:  doZones,
}

func doZones(c *cli.Context) error {
	zones, err := dozens.GetZones()
	if err != nil {
		return errors.Wrap(err, "error in GetZones")
	}
	if err := json.NewEncoder(os.Stdout).Encode(zones); err != nil {
		return errors.Wrap(err, "error in Encode")
	}
	return nil
}
