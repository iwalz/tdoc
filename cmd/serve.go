package cmd

import (
	"io/ioutil"
	"log"
	"net/http"

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
		content, err := ioutil.ReadFile(args[0])
		if err != nil {
			return err
		}
		p := &parser.TdocParserImpl{}
		l := parser.NewLexer(string(content), SvgDir)
		p.Parse(l)
		m := renderer.NewMiddleware(p.AST())

		http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			m.Render(w, req)
		}))
		err = http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Fatal("ListenAndServe:", err)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	// Configure port
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "HTTP port to serve output")
}
