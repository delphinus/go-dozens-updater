package godo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/delphinus/go-dozens"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var commandRecordCreate = cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Usage:   "Create a new record",
	Action:  doRecordCreate,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "domain, d",
			Usage: "Name of the domain (required)",
		},
		cli.StringFlag{
			Name:  "name, n",
			Usage: "Name of the record to create (required)",
		},
		cli.StringFlag{
			Name:  "type, t",
			Usage: "Type of the record: A, AAAA, CNAME, MX, TXT, SRV, ALIAS (required)",
		},
		cli.IntFlag{
			Name:  "prio, p",
			Usage: "Priority of the record. (required if type is MX)",
		},
		cli.StringFlag{
			Name:  "content, c",
			Usage: "Content of the record (required)",
		},
		cli.StringFlag{
			Name:  "ttl, T",
			Usage: "TTL of the record",
		},
	},
}

var availableType = map[string]bool{
	"A":     true,
	"AAAA":  true,
	"CNAME": true,
	"MX":    true,
	"TXT":   true,
	"SRV":   true,
	"ALIAS": true,
}

func doRecordCreate(c *cli.Context) error {
	domain := c.String("domain")
	if domain == "" {
		return errors.New("Please specify domain of the record")
	}

	name := c.String("name")
	rType := c.String("type")
	if rType == "" {
		return errors.New("Please specify type of the record")
	}
	if _, ok := availableType[rType]; !ok {
		return errors.New("Please specify a valid type")
	}

	rawPrio := c.Int("prio")
	if rawPrio == 0 && rType == "MX" {
		return errors.New("Please specify a valid prio")
	}
	prio := fmt.Sprintf("%d", rawPrio)
	if prio == "0" {
		prio = ""
	}

	content := c.String("content")
	if content == "" {
		return errors.New("Please specify content of the record")
	}

	body := dozens.RecordCreateBody{
		Domain:  domain,
		Name:    name,
		Type:    rType,
		Prio:    prio,
		Content: content,
		TTL:     c.String("ttl"),
	}

	record, err := dozens.RecordCreate(Config.Token, body)
	if err != nil {
		return errors.Wrap(err, "error in RecordCreate")
	}

	if err := json.NewEncoder(os.Stdout).Encode(record); err != nil {
		return errors.Wrap(err, "error in Encode")
	}

	return nil
}
