package godo

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"
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

	originalConfigFile := ConfigFile
	ConfigFile = tmp.Name()

	go func() {
		<-ctx.Done()
		_ = tmp.Close()
		ConfigFile = originalConfigFile
	}()

	if txt != "" {
		_, err := tmp.WriteString(txt)
		if err != nil {
			return "", err
		}
	}

	return tmp.Name(), nil
}

func TestSetupConfigCreateConfig(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	_, err := makeTmpConfig(ctx, "")
	if err != nil {
		t.Errorf("error to create tmp config: %v", err)
	}
	defer cancel()

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
	_, err := makeTmpConfig(ctx, fmt.Sprintf(`{
		"key": "hoge",
		"user": "fuga",
		"token": "hoge",
		"expiresAt": "%s"
	}`, time.Now().Add(time.Duration(1)*time.Minute).Format(time.RFC3339)))
	defer cancel()

	if err != nil {
		t.Errorf("error to create tmp config: %v", err)
	}
	defer cancel()

	if err := SetupConfig(); err != nil {
		t.Errorf("error occurred: %v", err)
	}
}
