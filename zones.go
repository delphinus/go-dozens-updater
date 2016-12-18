package godo

import (
	"github.com/urfave/cli"
)

var commandZones = cli.Command{
	Name:    "zones",
	Aliases: []string{"z"},
	Usage:   "List up zones in dozens",
	Action:  doZones,
}

func doZones(c *cli.Context) error {
	return nil
}
