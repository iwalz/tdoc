package cmd

import (
	"errors"
	"os"

	"github.com/dnephin/cobra"
)

func CheckFile(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("File missing")
	}

	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		return errors.New("File not found " + args[0])
	}

	return nil
}
