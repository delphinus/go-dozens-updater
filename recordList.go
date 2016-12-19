package godo

import (
	"encoding/json"
	"os"

	"github.com/delphinus/godo/lib/dozens"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var commandRecordList = cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "List up records in the zone",
	Action:  doRecordList,
}

func doRecordList(c *cli.Context) error {
	zoneName := c.Args().Get(0)
	if zoneName == "" {
		return errors.New("Please specify zone name to list")
	}

	record, err := dozens.RecordList(Config.Token, zoneName)
	if err != nil {
		return errors.Wrap(err, "error in RecordList")
	}

	if err := json.NewEncoder(os.Stdout).Encode(record); err != nil {
		return errors.Wrap(err, "error in Encode")
	}

	return nil
}
