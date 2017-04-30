package main

import (
	"os"

	"github.com/delphinus/godo"
)

func main() {
	_ = godo.NewApp().Run(os.Args)
}
