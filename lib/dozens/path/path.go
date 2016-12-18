package path

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

// Path means the path struct
type Path struct {
	Base  *url.URL
	Chunk string
}

// NewPath returns Path struct
func NewPath(chunk string) Path {
	return Path{baseURL, chunk}
}

func (p Path) String() string {
	u, err := url.Parse(p.Base.String())
	if err != nil {
		panic(err)
	}
	u.Path = strings.Join([]string{u.Path, p.Chunk}, "/")
	return u.String()
}

// Zone means `http://dozens.jp/api/zone.json`
func Zone() Path {
	return NewPath("zone.json")
}
