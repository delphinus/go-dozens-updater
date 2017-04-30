package godo

import (
	"encoding/json"
	"net/http"

	"github.com/koron/go-dproxy"
	"github.com/pkg/errors"
	"github.com/rdegges/go-ipify"
)

var ipv6URL = "https://jsonip.com"

func getIP(ipv6 bool) (string, error) {
	if !ipv6 {
		return ipify.GetIp()
	}

	resp, err := http.Get(ipv6URL)
	if err != nil {
		return "", errors.Wrap(err, "error in Get")
	}
	defer func() { _ = resp.Body.Close() }()

	var v interface{}
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return "", errors.Wrap(err, "error in Decode")
	}

	return dproxy.New(v).M("ip").String()
}
