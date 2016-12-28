package godo

import (
	"encoding/json"
	"os"

	"github.com/delphinus/godo/lib/dozens"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var commandRecordUpdate = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Usage:   "Update the specified record",
	Action:  doRecordUpdate,
	Flags: []cli.Flag{
		cli.UintFlag{
			Name:  "prio, p",
			Usage: "Priority to set",
		},
		cli.StringFlag{
			Name:  "content, c",
			Usage: "Content to set",
		},
		cli.StringFlag{
			Name:  "ttl, t",
			Usage: "TTL to set",
		},
	},
}

func doRecordUpdate(c *cli.Context) error {
	if c.NArg() == 0 {
		return errors.New("Please specify record ID")
	}
	recordID := c.Args().Get(0)

	body := dozens.RecordUpdateBody{
		Prio:    c.Uint("prio"),
		Content: c.String("content"),
		TTL:     c.String("ttl"),
	}

	record, err := dozens.RecordUpdate(Config.Token, recordID, body)
	if err != nil {
		return errors.Wrap(err, "errorin RecordUpdate")
	}

	if err := json.NewEncoder(os.Stdout).Encode(record); err != nil {
		return errors.Wrap(err, "error in Encode")
	}
	return nil
}
