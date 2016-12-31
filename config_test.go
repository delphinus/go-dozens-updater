package godo

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"
	"time"
)

func TestSetupConfigHaveValidConfig(t *testing.T) {
	originalToken := Config.Token
	Config.Token = "hoge"
	defer func() { Config.Token = originalToken }()

	originalExpiredAt := Config.ExpiredAt
	Config.ExpiredAt = time.Now().Add(-time.Duration(1) * time.Minute)
	defer func() { Config.ExpiredAt = originalExpiredAt }()

	if err := SetupConfig(); err != nil {
		t.Errorf("error ocurred: %v", err)
	}
}

func TestSetupConfigCreateConfig(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("", "")
	originalConfigFile := ConfigFile
	ConfigFile = path.Join(tmpDir, "config")
	defer func() { ConfigFile = originalConfigFile }()

	originalIsValid := Config.IsValid
	Config.IsValid = true
	defer func() { Config.IsValid = originalIsValid }()

	err := SetupConfig()
	expected := "error in createConfig"
	if err == nil || strings.Index(err.Error(), expected) != 0 {
		t.Errorf("error differs: %v", err)
	}
}
