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
	components   []elements.Element
	matrix       map[int]map[int]bool
}

var dir string

func NewMiddleware(e elements.Element, d string) *Middleware {
	dir = d
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
	e.Reset()
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
					// Row added, y is row*height+heightoffset
					// x is only offset
					b.rows = b.rows + 1
					// Set y to new row and x to 1
					elem.SetY(b.rows)
					elem.SetX(1)
					// fmt.Println("Row added, X: ", elem.X(), " Y: ", elem.Y())
				} else {
					b.columns = b.columns + 1
					// Column added, x is column*width+offset
					elem.SetY(1)
					elem.SetX(b.columns)
					// fmt.Println("Col added, X: ", elem.X(), " Y: ", elem.Y())
				}
			} else {

				if b.columns > b.rows {
					elem.SetX(b.columns)
					elem.SetY(b.count%b.rows + 1)
				} else {
					elem.SetX(b.count%b.columns + 1)
					elem.SetY(b.rows)

				}
				// fmt.Println("Else, X: ", elem.X(), " Y: ", elem.Y())
			}
			// Add elements here
			b.count = b.count + 1
			b.components = append(b.components, elem)
		}
		// Add border here
	}
}

func (m *Middleware) Render(w http.ResponseWriter, req *http.Request) error {

	b := &BaseMatrix{rows: 1, columns: 1}
	b.scan(m.matrix)

	canvas = svg.New(w)
	canvas.Start(b.width(), b.height())
	for _, v := range b.components {
		c, ok := v.(*elements.Component)
		if ok {
			file := dir + "/" + c.Typ + ".svg"
			posy := (c.Y()*height - height) + b.heightoffset
			posx := (c.X()*width - width) + b.widthoffset
			placepic(posx, posy, file)
		}
	}
	canvas.End()
	m.matrix.Reset()

	return nil
}
