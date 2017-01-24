package renderer

import (
	"io"
	"net/http"

	svg "github.com/ajstarks/svgo"
	"github.com/iwalz/tdoc/elements"
	"github.com/iwalz/tdoc/table"
)

type Middleware struct {
	root  elements.Element
	cl    elements.ComponentsList
	table *table.Table
}

func NewMiddleware(r elements.Element, cl elements.ComponentsList) *Middleware {
	return &Middleware{
		root: r,
		cl:   cl,
	}
}

// Scans recursivly
func scan(e elements.Element, cl elements.ComponentsList) *table.Table {
	e.Reset()
	t := table.NewTable(cl)
	for {
		// break loop if next is nil
		elem := e.Next()
		if elem == nil {
			break
		}
		c := elem.(*elements.Component)

		if elem.HasChilds() {
			scan(c, cl)
		} else {
			// Add element
			t.Add(c)
		}
	}

	return t
}

func (m *Middleware) Render(w io.Writer, req *http.Request) error {

	t := scan(m.root, m.cl)
	canvas := svg.New(w)
	canvas.Start(t.Width(), t.Height())
	t.Render(canvas)
	canvas.End()
	m.root.Reset()

	return nil
}
