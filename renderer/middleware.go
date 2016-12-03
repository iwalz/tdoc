package renderer

import (
	"io"
	"net/http"

	svg "github.com/ajstarks/svgo"
	"github.com/iwalz/tdoc/elements"
)

type Renderer interface {
	Render(io.Writer) error
}

type Middleware struct {
	matrix elements.Element
}

func NewMiddleware(e elements.Element) *Middleware {
	return &Middleware{
		matrix: e,
	}
}

func (m *Middleware) Render(w http.ResponseWriter, req *http.Request) error {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 125, "fill:none;stroke:black")
	s.End()

	return nil
}
