package godo

import (
	"github.com/urfave/cli"
)

var commandList = cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "List up entries in dozens",
	Action:  doList,
}

func doList(c *cli.Context) error {
	return nil
}
