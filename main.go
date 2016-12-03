package main

import (
	"os"

	"github.com/iwalz/tdoc/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
