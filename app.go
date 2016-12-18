package godo

import (
	"github.com/urfave/cli"
)

// NewApp will return godo App
func NewApp() *cli.App {
	app := cli.NewApp()
	return app
}
