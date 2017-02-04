package renderer

import (
	svg "github.com/ajstarks/svgo"
	"github.com/iwalz/tdoc/elements"
	"github.com/iwalz/tdoc/image"
	"github.com/iwalz/tdoc/table"
)

type MWare interface {
	Scan(elements.Element, elements.ComponentsList) table.TableAbstract
	Render(*svg.SVG) error
}

type Middleware struct {
	root    elements.Element
	cl      elements.ComponentsList
	table   table.TableAbstract
	rewrite image.Rewriter
}

func NewMiddleware(r elements.Element, cl elements.ComponentsList) *Middleware {
	rewrite := image.NewRewrite()
	return &Middleware{
		root:    r,
		cl:      cl,
		rewrite: rewrite,
	}
}

// Scans recursivly
func (m *Middleware) Scan(e elements.Element, cl elements.ComponentsList) table.TableAbstract {
	e.Reset()
	t := table.NewTable(cl)
	t.SetRewriter(m.rewrite)
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
			st.(*table.Table).SetRewriter(m.rewrite)
			t.Add(st)
		} else {
			// Add element
			cell := table.NewCell(c, cl)
			cell.SetRewriter(m.rewrite)
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
