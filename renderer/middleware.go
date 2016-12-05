package renderer

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/iwalz/tdoc/elements"
)

type SVG struct {
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
	Doc    string `xml:",innerxml"`
}

var (
	byrow                                          bool
	startx, starty, count, gutter, gwidth, gheight int
	canvas                                         *svg.SVG
)

// placepic puts a SVG file at a location
func placepic(x, y int, filename string) (int, int) {
	var (
		s             SVG
		width, height int
		wunit, hunit  string
	)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 0, 0
	}
	defer f.Close()
	if err := xml.NewDecoder(f).Decode(&s); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse (%v)\n", err)
		return 0, 0
	}

	// read the width and height, including any units
	// if there are errors use 10 as a default
	nw, _ := fmt.Sscanf(s.Width, "%d%s", &width, &wunit)
	if nw < 1 {
		width = 10
	}
	nh, _ := fmt.Sscanf(s.Height, "%d%s", &height, &hunit)
	if nh < 1 {
		height = 10
	}
	canvas.Group(`clip-path="url(#pic)"`, fmt.Sprintf(`transform="translate(%d,%d)"`, x, y))
	canvas.ClipPath(`id="pic"`)
	canvas.Rect(0, 0, width, height)
	canvas.ClipEnd()
	io.WriteString(canvas.Writer, s.Doc)
	canvas.Gend()
	return width, height
}

// compose places files row or column-wise
func compose(x, y, n int, rflag bool, files []string) {
	px := x
	py := y
	var pw, ph int
	for i, f := range files {
		if i > 0 && i%n == 0 {
			if rflag {
				px = x
				py += gutter + ph
			} else {
				px += gutter + pw
				py = y
			}
		}
		pw, ph = placepic(px, py, f)
		fmt.Println("Widht", pw)
		fmt.Println("Height", ph)
		if rflag {
			px += gutter + pw
		} else {
			py += gutter + ph
		}
	}
}

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
	startx = 0
	starty = 0
	byrow = true
	count = 3
	gutter = 100
	gwidth = 10240
	gheight = 7680

	canvas = svg.New(w)
	canvas.Start(gwidth, gheight)
	compose(startx, starty, count, byrow, []string{"pics/cloud.svg", "pics/cloud.svg"})
	canvas.End()
	/*
		w.Header().Set("Content-Type", "image/svg+xml")
		s := svg.New(w)
		s.Start(500, 500)
		s.Circle(250, 250, 125, "fill:none;stroke:black")
		s.End()
	*/
	return nil
}
