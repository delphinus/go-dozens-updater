package godo

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/delphinus/go-dozens"
	"github.com/delphinus/go-dozens/endpoint"
	"github.com/jarcoal/httpmock"
)

func TestSetupConfigHaveValidConfig(t *testing.T) {
	Config = Configs{
		Token:     "hoge",
		ExpiresAt: time.Now().Add(-time.Duration(1) * time.Minute),
	}

	if err := SetupConfig(); err != nil {
		t.Errorf("error ocurred: %v", err)
	}
}

func makeTmpConfig(ctx context.Context, txt string) (string, error) {
	tmp, err := ioutil.TempFile("", "")
	if err != nil {
		return "", err
	}

	go func() {
		<-ctx.Done()
		_ = tmp.Close()
	}()

	if txt != "" {
		_, err := tmp.WriteString(txt)
		if err != nil {
			return "", err
		}
	}

	return tmp.Name(), nil
}

func makeMockDozens(ctx context.Context, method, url string, responder httpmock.Responder) {
	httpmock.Activate()

	go func() {
		<-ctx.Done()
		httpmock.DeactivateAndReset()
	}()

	httpmock.RegisterResponder(method, url, responder)
}

func TestSetupConfigCreateConfig(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	f, err := makeTmpConfig(ctx, "")
	if err != nil {
		t.Errorf("error to create tmp config: %v", err)
	}
	original := ConfigFile
	ConfigFile = f
	defer func() {
		cancel()
		ConfigFile = original
	}()

	Config = Configs{
		IsValid: true,
	}

	expected := "error in createConfig"
	if err := SetupConfig(); err == nil || strings.Index(err.Error(), expected) != 0 {
		t.Errorf("error differs: %v", err)
	}
}

func TestSetupConfigReadConfigValidly(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	f, err := makeTmpConfig(ctx, fmt.Sprintf(`{
		"key": "hoge",
		"user": "fuga",
		"token": "hoge",
		"expiresAt": "%s"
	}`, time.Now().Add(time.Duration(1)*time.Minute).Format(time.RFC3339)))

	if err != nil {
		t.Errorf("error to create tmp config: %v", err)
	}
	original := ConfigFile
	ConfigFile = f
	defer func() {
		cancel()
		ConfigFile = original
	}()

	if err := SetupConfig(); err != nil {
		t.Errorf("error occurred: %v", err)
	}
}

func TestSetupConfigReadConfigCreateConfigAndFail(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	f, err := makeTmpConfig(ctx, fmt.Sprintf(`{
		"key": "hoge",
		"user": "fuga",
		"token": "hoge",
		"expiresAt": "%s"
	}`, time.Now().Add(-time.Duration(1)*time.Minute).Format(time.RFC3339)))

	if err != nil {
		t.Errorf("error to create tmp config: %v", err)
	}
	original := ConfigFile
	ConfigFile = f
	defer func() {
		cancel()
		ConfigFile = original
	}()

	url := endpoint.Authorize().String()
	responder, _ := httpmock.NewJsonResponder(http.StatusOK, &dozens.AuthorizeResponse{"hoge"})
	makeMockDozens(ctx, "GET", url, responder)

	if err := SetupConfig(); err != nil {
		t.Errorf("error occuured: %v", err)
	}
}
