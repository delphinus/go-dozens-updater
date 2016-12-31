package godo

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"
)

func TestSetupConfigCreateConfigError(t *testing.T) {
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
