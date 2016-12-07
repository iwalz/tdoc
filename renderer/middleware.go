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

//var components chan elements.Element

type PictureLocation struct {
	x   int
	y   int
	typ string
}

type SVG struct {
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
	Doc    string `xml:",innerxml"`
}

const width = 100
const height = 100

const borderwidth = 25
const borderheight = 25

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

type Renderer interface {
	Render(io.Writer) error
}

type Middleware struct {
	matrix elements.Element
}

// Flatted and SVG compatible AST implementation
type BaseMatrix struct {
	widthoffset  int // width
	heightoffset int // height
	rows         int
	columns      int
	count        int
	posx         int      // Current position on x
	posy         int      // Current position on y
	canvas       *svg.SVG // The root SVG
	components   []*PictureLocation
}

func NewMiddleware(e elements.Element) *Middleware {
	return &Middleware{
		matrix: e,
	}
}

func (b *BaseMatrix) HasFreeSlots() bool {
	if b.rows*b.columns-b.count > 0 {
		return true
	}
	return false
}

func (b BaseMatrix) width() int {
	return b.widthoffset + (b.columns * width)
}

func (b BaseMatrix) height() int {
	return b.heightoffset + (b.rows * height)
}

// Scans recursivly add elements and calculates
// required rows and columns + offset for borders of nested COMPONENTs
func (b *BaseMatrix) scan(e elements.Element) {
	for {
		// break loop if next is nil
		elem := e.Next()
		if elem == nil {
			break
		}

		// e is a nested COMPONENT
		// increase offset for the border
		// and go on looping for this COMPONENT
		if elem.HasChilds() {
			b.heightoffset += borderheight
			b.widthoffset += borderwidth
			b.scan(elem)
		} else {
			if !b.HasFreeSlots() {
				// Resize matrix if no available slots
				if b.columns > b.rows {
					b.rows = b.rows + 1
					b.columns = 1
				} else {
					b.columns = b.columns + 1
				}
			}

			// Since it's the left starting corner, the width and height
			// of itself has to be substracted
			b.posy = (b.heightoffset + (b.rows * height)) - width
			b.posx = (b.widthoffset + (b.columns * width)) - height

			// Add elements here
			b.count = b.count + 1
			typ := elem.(*elements.Component).Typ
			b.components = append(b.components, &PictureLocation{x: b.posx, y: b.posy, typ: typ})
		}
	}
}

func (m *Middleware) Render(w http.ResponseWriter, req *http.Request) error {

	b := &BaseMatrix{rows: 1, columns: 1}
	b.scan(m.matrix)

	canvas = svg.New(w)
	canvas.Start(b.width(), b.height())
	for _, v := range b.components {
		file := "/home/ingo/svg/" + v.typ + ".svg"
		placepic(v.x, v.y, file)
	}
	canvas.End()
	m.matrix.Reset()

	return nil
}
