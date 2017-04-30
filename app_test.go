package godo

import (
	"io/ioutil"
	"net/http"
	"path"
	"testing"
	"time"

	dozens "github.com/delphinus/go-dozens"
	"github.com/delphinus/go-dozens/endpoint"
	"github.com/jarcoal/httpmock"
)

func TestNewApp(t *testing.T) {
	_ = NewApp()
}

func TestNewAppBefore(t *testing.T) {
	app := NewApp()
	_ = app.Run([]string{"godo", "zone"})
}

func TestNewAppBeforeSetupConfig(t *testing.T) {
	app := NewApp()

	d, _ := ioutil.TempDir("", "")
	original := ConfigFile
	ConfigFile = path.Join(d, "config")
	defer func() { ConfigFile = original }()

	Config = Configs{
		IsValid:   true,
		Token:     "hoge",
		ExpiresAt: time.Now().Add(-time.Duration(1) * time.Minute),
	}

	httpmock.Activate()
	{
		url := endpoint.Authorize().String()
		responder, _ := httpmock.NewJsonResponder(http.StatusOK, &dozens.AuthorizeResponse{"hoge"})
		httpmock.RegisterResponder("GET", url, responder)
	}
	{
		url := endpoint.ZoneList().String()
		responder, _ := httpmock.NewJsonResponder(http.StatusOK, &dozens.ZoneResponse{})
		httpmock.RegisterResponder("GET", url, responder)
	}
	defer httpmock.DeactivateAndReset()

	_ = app.Run([]string{"godo", "zone", "list"})
}
