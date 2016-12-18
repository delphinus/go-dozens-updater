package godo

import (
	"encoding/json"
	"os"

	"github.com/delphinus/godo/lib/dozens"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var commandZone = cli.Command{
	Name:    "zone",
	Aliases: []string{"z"},
	Usage:   "List up zones in dozens",
	Action:  doZone,
}

func doZone(c *cli.Context) error {
	zone, err := dozens.GetZone()
	if err != nil {
		return errors.Wrap(err, "error in GetZone")
	}
	if err := json.NewEncoder(os.Stdout).Encode(zone); err != nil {
		return errors.Wrap(err, "error in Encode")
	}
	return nil
}
