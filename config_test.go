package godo

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"
	"time"
)

func TestSetupConfigHaveValidConfig(t *testing.T) {
	Config = Configs{
		Token:     "hoge",
		ExpiredAt: time.Now().Add(-time.Duration(1) * time.Minute),
	}

	if err := SetupConfig(); err != nil {
		t.Errorf("error ocurred: %v", err)
	}
}

func TestSetupConfigCreateConfig(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("", "")
	originalConfigFile := ConfigFile
	ConfigFile = path.Join(tmpDir, "config")
	defer func() { ConfigFile = originalConfigFile }()

	Config = Configs{
		IsValid: true,
	}

	err := SetupConfig()
	expected := "error in createConfig"
	if err == nil || strings.Index(err.Error(), expected) != 0 {
		t.Errorf("error differs: %v", err)
	}
}
