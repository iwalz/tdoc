package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/dnephin/cobra"
)

var port string

// themeCmd represents the theme command
var themeCmd = &cobra.Command{
	Use:   "theme",
	Short: "Manage tdoc themes",
	Long:  `Already long enough`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if verbose {
			log.SetLevel(log.InfoLevel)
		}
		if debug {
			log.SetLevel(log.DebugLevel)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(themeCmd)
}
