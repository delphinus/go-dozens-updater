package godo

import (
	"fmt"

	"github.com/delphinus/go-dozens"
	"github.com/pkg/errors"
	"github.com/rdegges/go-ipify"
	"github.com/urfave/cli"
)

func renew(c *cli.Context) error {
	zone := c.String("zone")
	domain := c.String("domain")
	showIPOnly := c.Bool("show-ip-only")
	if zone == "" || domain == "" {
		return errors.New("Please specify zone & domain")
	}
	if showIPOnly {
		return showIP(c, zone, domain)
	}

	ip, err := ipify.GetIp()
	if err != nil {
		return errors.Wrap(err, "error in GetIp")
	}

	if Config.MyIP == ip {
		fmt.Fprintf(c.App.Writer, "IP Address has not changed: %s. exiting...\n", ip)
		return nil
	}
	fmt.Fprintf(c.App.Writer, "IP Address has changed: %s. Try to update dozens entry...\n", ip)

	recordID, oldIP, err := searchRecord(zone, domain)
	if err != nil {
		return errors.Wrap(err, "error in searchRecord")
	}

	if oldIP == ip {
		fmt.Fprintln(c.App.Writer, "IP Address has already registered. Do nothing.")
	} else {
		body := dozens.RecordUpdateBody{Content: ip}
		if _, err := dozens.RecordUpdate(Config.Token, recordID, body); err != nil {
			return errors.Wrap(err, "error in RecordUpdate")
		}
		fmt.Fprintln(c.App.Writer, "IP Adress has renewed successfully.")
	}

	Config.MyIP = ip
	if err := SaveConfig(); err != nil {
		return errors.Wrap(err, "error in SetupConfig")
	}
	return nil
}

func showIP(c *cli.Context, zone, domain string) error {
	ip, err := ipify.GetIp()
	if err != nil {
		return errors.Wrap(err, "error in GetIp")
	}

	recordID, oldIP, err := searchRecord(zone, domain)
	if err != nil {
		return errors.Wrap(err, "error in searchRecord")
	}

	fmt.Fprintf(c.App.Writer, "My IP Address: %s\n", ip)
	fmt.Fprintf(c.App.Writer, "Info on dozens. recordID: %s, content: %s\n", recordID, oldIP)
	return nil
}

func searchRecord(zone, domain string) (string, string, error) {
	record, err := dozens.RecordList(Config.Token, zone)
	if err != nil {
		return "", "", errors.Wrap(err, "error in RecordList")
	}

	recordID, oldIP := searchDomain(record, domain)
	if recordID == "" {
		return "", "", errors.Errorf("the domain '%s' does not found on the zone '%s'", domain, zone)
	}

	return recordID, oldIP, nil
}

func searchDomain(record dozens.RecordResponse, domain string) (string, string) {
	for _, r := range record.Record {
		if r.Name == domain {
			return r.ID, r.Content
		}
	}
	return "", ""
}
