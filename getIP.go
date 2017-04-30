package godo

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rdegges/go-ipify"
)

var ipv6URL = "https://jsonip.com"

type ipv6Resp struct {
	IP            string `json:"ip"`
	About         string `json:"about"`
	Pro           string `json:"Pro!"`
	RejectFascism string `json:"reject-fascism"`
}

func getIP(ipv6 bool) (string, error) {
	if !ipv6 {
		return ipify.GetIp()
	}

	resp, err := http.Get(ipv6URL)
	if err != nil {
		return "", errors.Wrap(err, "error in Get")
	}
	defer func() { _ = resp.Body.Close() }()

	var r ipv6Resp
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return "", errors.Wrap(err, "error in Decode")
	}

	return r.IP, nil
}
