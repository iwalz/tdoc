package cmd

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/dnephin/cobra"
	"github.com/iwalz/tdoc/outputs"
	"github.com/spf13/afero"
)

var port string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve input.tdoc",
	Short: "Serve output via webserver",
	Long:  `Already long enough`,
	Args:  CheckFile,
	RunE: func(cmd *cobra.Command, args []string) error {
		if verbose {
			log.SetLevel(log.InfoLevel)
		}
		if debug {
			log.SetLevel(log.DebugLevel)
		}

		fs := afero.NewOsFs()
		web := outputs.NewWeb(fs, extension, SvgDir)
		executor := outputs.NewExecutor(fs, extension)
		executor.Exec(web, args)

		http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			web.WebHandler(req.URL.Path, w)
		}))

		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Error("Failed to bind port:", err)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	// Configure port
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "HTTP port to serve output")
}
