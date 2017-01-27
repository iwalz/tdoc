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
			st := m.Scan(c, cl)
			st.(*table.Table).SetBorder(table.DASHED_BORDER)
			st.(*table.Table).SetImage(cl.GetFilenameByType(c))
			st.(*table.Table).SetCaption(c.Identifier)
			t.Add(st)
		} else {
			// Add element
			cell := table.NewCell(c, cl)
			t.Add(cell)
		}
	}
	m.table = t

	return t
}

func (m *Middleware) Render(svg *svg.SVG) error {

	m.table.Render(svg)
	m.root.Reset()

	return nil
}
