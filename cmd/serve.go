package cmd

import (
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/dnephin/cobra"
	"github.com/iwalz/tdoc/parser"
	"github.com/iwalz/tdoc/renderer"
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

		http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			content, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Error("Could not open file ", args[0])
			}
			p := &parser.TdocParserImpl{}
			l := parser.NewLexer(string(content), SvgDir)
			p.Parse(l)
			m := renderer.NewMiddleware(p.AST())

			m.Render(w, req)
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
