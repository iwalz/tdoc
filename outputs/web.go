package outputs

import (
	"errors"
	"html/template"
	"net/http"
	"strings"

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
	store      Store
}

type Store struct {
	Files map[string][]string // Registered files, key is the directory
	Path  string
}

func NewWeb(fs afero.Fs, ext string, dir string) *Web {
	web := &Web{fs: fs, ext: ext, dir: dir}
	web.parser = &parser.TdocParserImpl{}
	web.cl = elements.NewComponentsList(dir)
	web.cl.Parse()

	s := Store{}
	s.Files = make(map[string][]string, 0)

	web.store = s

	return web
}

func (w *Web) HandleDir(d string) error {
	return nil
}

func getDirectoryAndFilename(file string) (string, string) {
	index := strings.LastIndex(file, "/")
	if index == -1 {
		return "/", file
	}
	d := file[:index]
	f := file[index+1:]

	return d, f
}

func (w *Web) HandleFile(file string) error {
	d, f := getDirectoryAndFilename(file)

	w.store.Files[d] = append(w.store.Files[d], f)
	return nil
}

func (web *Web) WebHandler(path string, w http.ResponseWriter) error {
	p := path[1:]

	// Check if it's an asset
	if p == "reset.css" || p == "styles.css" || p == "file.svg" || p == "folder.svg" {
		web.renderAssets(p, w)
	}

	if _, ok := web.store.Files[p]; ok || p == "" {
		web.renderDir(p, w)
		return nil
	}

	web.renderFile(p, w)

	return nil
}

func (web *Web) renderAssets(path string, w http.ResponseWriter) error {
	header := w.Header()
	if strings.HasSuffix(path, ".css") {
		header.Set("Content-Type", "text/css")
	}

	if strings.HasSuffix(path, ".svg") {
		header.Set("Content-Type", "image/svg+xml")
	}

	d, err := Asset("templates/" + path)
	if err != nil {
		return err
	}
	w.Write(d)

	return nil
}

func stripPath(prefix, path string) string {
	return strings.Replace(path, prefix+"/", "", 1)
}

func startsWith(prefix, path string) bool {
	if prefix == "" && path == "" {
		return false
	}

	if strings.HasPrefix("/"+path, prefix+"/") {
		return true
	}

	return false
}

func (web *Web) renderDir(path string, w http.ResponseWriter) error {
	header := w.Header()
	header.Set("Content-Type", "text/html")

	web.store.Path = path

	d, err := Asset("templates/tdoc.html")
	if err != nil {
		return err
	}

	fmap := template.FuncMap{
		"startsWith": startsWith,
		"stripPath":  stripPath,
	}
	t := template.New("index").Funcs(fmap)
	t, err = t.Parse(string(d))
	if err != nil {
		return err
	}
	t.Execute(w, web.store)

	return nil
}

func (web *Web) renderFile(path string, w http.ResponseWriter) error {
	header := w.Header()
	header.Set("Content-Type", "image/svg+xml")

	content, err := afero.ReadFile(web.fs, path)
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
