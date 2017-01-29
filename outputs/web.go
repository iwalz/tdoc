package outputs

import (
	"errors"
	"io/ioutil"
	"net/http"

	svg "github.com/ajstarks/svgo"
	"github.com/iwalz/tdoc/elements"
	"github.com/iwalz/tdoc/parser"
	"github.com/iwalz/tdoc/renderer"
	"github.com/spf13/afero"
)

type Web struct {
	Output
	fs         afero.Fs                // File system mock
	ext        string                  // File extension
	dir        string                  // SVG directory
	cl         elements.ComponentsList // component list
	parser     parser.Tdoc             // YACC Parser
	lexer      parser.TdocLexer        // Lexer
	middleware renderer.MWare          // Middleware
	dirs       map[string]bool         // Registered directories
	files      map[string]bool         // Registered files
}

func NewWeb(fs afero.Fs, ext string, dir string) *Web {
	web := &Web{fs: fs, ext: ext, dir: dir}
	web.parser = &parser.TdocParserImpl{}
	web.cl = elements.NewComponentsList(dir)
	web.cl.Parse()

	web.dirs = make(map[string]bool, 0)
	web.files = make(map[string]bool, 0)

	return web
}

func (w *Web) HandleDir(d string) error {
	w.dirs[d] = true
	return nil
}

func (w *Web) HandleFile(file string) error {
	w.files[file] = true

	return nil
}

func (web *Web) WebHandler(path string, w http.ResponseWriter) error {
	p := path[1:]
	if ok, _ := web.files[p]; ok {
		web.renderFile(p, w)
		return nil
	}

	if ok, _ := web.dirs[p]; ok {
		//web.renderDir(p, w)
		return nil
	}

	return nil
}

func (web *Web) renderFile(path string, w http.ResponseWriter) error {
	// Strip trailing slash
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.New("Could not open file " + path)
	}

	web.lexer = parser.NewLexer(string(content), web.cl)
	web.parser.Parse(web.lexer)
	ast := web.parser.AST()

	web.middleware = renderer.NewMiddleware(ast, web.cl)
	svg := svg.New(w)
	table := web.middleware.Scan(ast, web.cl)
	svg.Start(table.Width(), table.Height())
	web.middleware.Render(svg)
	svg.End()

	return nil
}
