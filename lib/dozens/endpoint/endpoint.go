package endpoint

import (
	"net/url"
	"strings"
)

var baseURL *url.URL

func init() {
	url, err := url.Parse("http://dozens.jp/api")
	if err != nil {
		panic(err)
	}
	baseURL = url
}

// Endpoint means the path struct
type Endpoint struct {
	Base  *url.URL
	Chunk string
}

// NewEndpoint returns Endpoint struct
func NewEndpoint(chunk string) Endpoint {
	return Endpoint{baseURL, chunk}
}

func (p Endpoint) String() string {
	u, err := url.Parse(p.Base.String())
	if err != nil {
		panic(err)
	}
	u.Path = strings.Join([]string{u.Path, p.Chunk}, "/")
	return u.String()
}

// Authorize means `http://dozens.jp/api/authorize.json`
func Authorize() Endpoint {
	return NewEndpoint("authorize.json")
}

// Zone means `http://dozens.jp/api/zone.json`
func Zone() Endpoint {
	return NewEndpoint("zone.json")
}

// Create means `http://dozens.jp/api/zone/create.json`
func Create() Endpoint {
	return NewEndpoint("zone/create.json")
}
