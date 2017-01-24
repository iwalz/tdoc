package renderer

import (
	svg "github.com/ajstarks/svgo"
	"github.com/iwalz/tdoc/elements"
	"github.com/iwalz/tdoc/table"
)

type Middleware struct {
	root  elements.Element
	cl    elements.ComponentsList
	table table.TableAbstract
}

func NewMiddleware(r elements.Element, cl elements.ComponentsList) *Middleware {
	return &Middleware{
		root: r,
		cl:   cl,
	}
}

// Scans recursivly
func (m *Middleware) Scan(e elements.Element, cl elements.ComponentsList) table.TableAbstract {
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
			m.Scan(c, cl)
		} else {
			// Add element
			t.Add(c)
		}
	}
	m.table = t

	return t
}

func (m *Middleware) Render(svg *svg.SVG) error {

	svg.Start(m.table.Width(), m.table.Height())
	m.table.Render(svg)
	svg.End()
	m.root.Reset()

	return nil
}
