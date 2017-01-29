package outputs

import (
	"io/ioutil"
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/iwalz/tdoc/elements"
	"github.com/iwalz/tdoc/parser"
	"github.com/iwalz/tdoc/renderer"
	"github.com/spf13/afero"
)

type File struct {
	Output
	fs         afero.Fs                // File system mock
	ext        string                  // File extension
	dir        string                  // SVG directory
	cl         elements.ComponentsList // component list
	parser     parser.Tdoc             // YACC Parser
	lexer      parser.TdocLexer        // Lexer
	middleware renderer.MWare          // Middleware
}

func NewFile(fs afero.Fs, ext string, dir string) *File {
	file := &File{fs: fs, ext: ext, dir: dir}
	file.parser = &parser.TdocParserImpl{}
	file.cl = elements.NewComponentsList(dir)
	file.cl.Parse()
	return file
}

func (f *File) HandleDir(d string) error {
	// Nothing to do
	return nil
}

func (f *File) HandleFile(file string) error {
	// New filename
	newFilename := strings.Replace(file, "."+f.ext, ".svg", 1)

	// Input file
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	f.lexer = parser.NewLexer(string(content), f.cl)
	f.parser.Parse(f.lexer)
	ast := f.parser.AST()

	f.middleware = renderer.NewMiddleware(ast, f.cl)
	nfile, err := f.fs.Create(newFilename)
	svg := svg.New(nfile)
	table := f.middleware.Scan(ast, f.cl)
	svg.Start(table.Width(), table.Height())
	f.middleware.Render(svg)
	svg.End()

	return nil
}
