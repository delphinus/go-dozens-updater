package godo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/Songmu/prompter"
	"github.com/delphinus/go-dozens"
	"github.com/pkg/errors"
)

// ConfigFile means config file for godo
var ConfigFile = filepath.Join(os.Getenv("HOME"), ".config", "godo", "godo.json")

// AuthInfo means info for authentication
type AuthInfo struct {
	Key  string `json:"key"`
	User string `json:"user"`
}

// Configs stores config for godo
type Configs struct {
	AuthInfo
	Token     string    `json:"token"`
	IsValid   bool      `json:"isValid"`
	MyIP      string    `json:"myIP"`
	MyIPv6    string    `json:"myIPv6"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// IsExpired will return true if it is expired
func (c Configs) IsExpired() bool {
	return c.ExpiresAt.Before(time.Now())
}

// Config is a loaded config from ConfigFile
var Config Configs

// TokenExpire means expiration of the token
const TokenExpire = time.Duration(20 * time.Hour)

// SetupConfig returns access token for dozens
func SetupConfig() error {
	if Config.Token != "" && Config.IsExpired() {
		return nil
	}

	if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "config file: `%s` does not exist. creating...\n", ConfigFile)
		if err := createConfig(); err != nil {
			return errors.Wrap(err, "error in createConfig")
		}

		return SaveConfig()
	}

	if err := readConfig(); err != nil {
		return errors.Wrap(err, "error in createConfig")
	}

	return SaveConfig()
}

func createConfig() error {
	if !Config.IsValid {
		Config.AuthInfo = inputAuthInfo()
	}

	authorizeResp, err := dozens.GetAuthorize(Config.Key, Config.User)
	if err != nil {
		return errors.Wrap(err, "error in GetAuthorize")
	}

	Config.Token = authorizeResp.AuthToken
	Config.ExpiresAt = time.Now().Add(TokenExpire)

	return nil
}

func readConfig() error {
	txt, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		return errors.Wrap(err, "error in ReadFile")
	}

	Config = Configs{}
	if err := json.Unmarshal(txt, &Config); err != nil {
		return errors.Wrap(err, "error in Unmarshal")
	}

	if Config.Token == "" || Config.IsExpired() {
		return createConfig()
	}

	return nil
}

// SaveConfig will save config to ConfigFile
func SaveConfig() error {
	configDir := filepath.Dir(ConfigFile)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0777); err != nil {
			return errors.Wrap(err, "error in MkDirAll")
		}
	}

	if Config.Token != "" && !Config.IsExpired() {
		Config.IsValid = true
	}

	json, err := json.Marshal(Config)
	if err != nil {
		return errors.Wrap(err, "error in Marshal")
	}

	if err := ioutil.WriteFile(ConfigFile, json, 0666); err != nil {
		return errors.Wrap(err, "error in WriteFile")
	}

	return nil
}

func inputAuthInfo() AuthInfo {
	key := prompter.Prompt("input API Key", "")
	user := prompter.Prompt("input DozensID", "")
	return AuthInfo{key, user}
}
