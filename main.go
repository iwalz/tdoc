package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/iwalz/tdoc/cmd"
)

func main() {
	v := cmd.RootCmd.Flag("verbose")
	d := cmd.RootCmd.Flag("debug")
	log.SetLevel(log.WarnLevel)

	if v.Changed {
		log.SetLevel(log.InfoLevel)
	}
	if d.Changed {
		log.SetLevel(log.DebugLevel)
	}

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}
