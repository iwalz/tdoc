package renderer

import (
	"io"
	"net/http"

	svg "github.com/ajstarks/svgo"
	"github.com/iwalz/tdoc/elements"
	"github.com/iwalz/tdoc/table"
)

type Middleware struct {
	root  *elements.Component
	cl    *elements.ComponentsList
	table *table.Table
}

func NewMiddleware(r *elements.Component, cl *elements.ComponentsList) *Middleware {
	return &Middleware{
		root: r,
		cl:   cl,
	}
}

// Scans recursivly add elements and calculates
// required rows and columns + offset for borders of nested COMPONENTs
func scan(e *elements.Component, cl *elements.ComponentsList) *table.Table {
	e.Reset()
	t := table.NewTable(cl)
	for {
		// break loop if next is nil
		elem := e.Next()
		if elem == nil {
			break
		}
		c := elem.(*elements.Component)

		// e is a nested COMPONENT
		// increase offset for the border
		// and go on looping for this COMPONENT
		if elem.HasChilds() {

			scan(c, cl)
		} else {
			// Add element
			t.Add(c)
		}
		// Add border here
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
