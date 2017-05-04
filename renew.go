package godo

import (
	"fmt"

	"github.com/delphinus/go-dozens"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func renew(c *cli.Context) error {
	zone := c.String("zone")
	domain := c.String("domain")
	showIPOnly := c.Bool("show-ip-only")
	ipv6 := c.Bool("ipv6")
	mail := c.Bool("mail")
	if zone == "" || domain == "" {
		return errors.New("Please specify zone & domain")
	}
	if showIPOnly {
		return showIP(c, zone, domain, ipv6)
	}

	ip, err := getIP(ipv6)
	if err != nil {
		return errors.Wrap(err, "error in GetIp")
	}

	if (ipv6 && Config.MyIPv6 == ip) || (!ipv6 && Config.MyIP == ip) {
		fmt.Fprintf(c.App.Writer, "IP Address has not changed: %s. exiting...\n", ip)
		return nil
	}
	fmt.Fprintf(c.App.Writer, "IP Address has changed: %s. Try to update dozens entry...\n", ip)

	recordID, oldIP, err := searchRecord(zone, domain, ipv6)
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
		fmt.Fprintln(c.App.Writer, "IP Adress has been renewed successfully.")

		if mail {
			var m string
			if ipv6 {
				m = "IPv6"
			} else {
				m = "IPv4"
			}
			content := fmt.Sprintf(`%s Address has been renewd successfully

old: %s
new: %s`, m, oldIP, ip)
			if err := sendMail([]byte(content)); err != nil {
				return errors.Wrap(err, "error in mail")
			}
		}
	}

	if ipv6 {
		Config.MyIPv6 = ip
	} else {
		Config.MyIP = ip
	}
	if err := SaveConfig(); err != nil {
		return errors.Wrap(err, "error in SetupConfig")
	}
	return nil
}

func showIP(c *cli.Context, zone, domain string, ipv6 bool) error {
	ip, err := getIP(ipv6)
	if err != nil {
		return errors.Wrap(err, "error in GetIp")
	}

	recordID, oldIP, err := searchRecord(zone, domain, ipv6)
	if err != nil {
		return errors.Wrap(err, "error in searchRecord")
	}

	fmt.Fprintf(c.App.Writer, "My IP Address: %s\n", ip)
	fmt.Fprintf(c.App.Writer, "Info on dozens. recordID: %s, content: %s\n", recordID, oldIP)
	return nil
}

func searchRecord(zone, domain string, ipv6 bool) (string, string, error) {
	record, err := dozens.RecordList(Config.Token, zone)
	if err != nil {
		return "", "", errors.Wrap(err, "error in RecordList")
	}

	recordID, oldIP := searchDomain(record, domain, ipv6)
	if recordID == "" {
		return "", "", errors.Errorf("the domain '%s' does not found on the zone '%s'", domain, zone)
	}

	return recordID, oldIP, nil
}

func searchDomain(record dozens.RecordResponse, domain string, ipv6 bool) (string, string) {
	t := ""
	if ipv6 {
		t = "AAAA"
	} else {
		t = "A"
	}
	for _, r := range record.Record {
		if r.Type == t && r.Name == domain {
			return r.ID, r.Content
		}
	}
	return "", ""
}
