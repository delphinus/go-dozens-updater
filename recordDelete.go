package godo

import (
	"encoding/json"
	"os"

	"github.com/delphinus/godo/lib/dozens"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var commandRecordDelete = cli.Command{
	Name:    "delete",
	Aliases: []string{"d"},
	Usage:   "Delete the specified record",
	Action:  doRecordDelete,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "id, i",
			Usage: "The record ID to delete",
		},
	},
}

func doRecordDelete(c *cli.Context) error {
	recordID := c.String("id")
	if recordID == "" {
		return errors.New("Please specify record ID")
	}

	record, err := dozens.RecordDelete(Config.Token, recordID)
	if err != nil {
		return errors.Wrap(err, "errorin RecordDelete")
	}

	if err := json.NewEncoder(os.Stdout).Encode(record); err != nil {
		return errors.Wrap(err, "error in Encode")
	}
	return nil
}
