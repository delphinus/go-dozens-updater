package dozens

import (
	"io"
	"net/http"

	"github.com/delphinus/godo/lib/dozens/endpoint"
	"github.com/pkg/errors"
)

// MakeGet returns request for dozens
func MakeGet(token string, p endpoint.Endpoint) (*http.Request, error) {
	req, err := http.NewRequest("GET", p.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error in NewRequest")
	}

	if err := setHeader(token, req); err != nil {
		return nil, errors.Wrap(err, "error in setHeader")
	}

	return req, nil
}

// MakePost returns request for dozens
func MakePost(token string, p endpoint.Endpoint, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", p.String(), body)
	if err != nil {
		return nil, errors.Wrap(err, "error in NewRequest")
	}

	if err := setHeader(token, req); err != nil {
		return nil, errors.Wrap(err, "error in setHeader")
	}

	return req, nil
}

func setHeader(token string, req *http.Request) error {
	req.Header.Set("X-Auth-Token", token)
	req.Header.Set("Content-Type", "application/json")

	return nil
}
