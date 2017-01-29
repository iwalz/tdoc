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
	fs  afero.Fs // File system mock
	ext string   // File extension
	dir string   // SVG directory
}

func NewFile(fs afero.Fs, ext string, dir string) *File {
	return &File{fs: fs, ext: ext, dir: dir}
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

	cl := elements.NewComponentsList(f.dir)
	cl.Parse()
	p := &parser.TdocParserImpl{}
	l := parser.NewLexer(string(content), cl)
	p.Parse(l)
	ast := p.AST()

	m := renderer.NewMiddleware(ast, cl)
	nfile, err := f.fs.Create(newFilename)
	svg := svg.New(nfile)
	table := m.Scan(ast, cl)
	svg.Start(table.Width(), table.Height())
	m.Render(svg)
	svg.End()

	return nil
}
