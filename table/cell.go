package table

import (
	"encoding/xml"
	"fmt"
	"io"
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

// Renders the pic at its location
func (c *cell) Render(svg *svg.SVG) error {
	var (
		s             SVG
		width, height int
		wunit, hunit  string
	)
	filename := c.cl.GetFilenameByType(c.Component())
	if filename == "" {
		return nil
	}
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := xml.NewDecoder(f).Decode(&s); err != nil {
		return err
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
	svg.Group(`clip-path="url(#pic)"`, fmt.Sprintf(`transform="translate(%d,%d)"`, c.X(), c.Y()))
	svg.ClipPath(`id="pic"`)
	svg.Rect(0, 0, width, height)
	svg.ClipEnd()
	io.WriteString(svg.Writer, s.Doc)
	svg.Gend()
	return nil
}

// Represents a cell in a table
type cell struct {
	component *elements.Component
	width     int
	height    int
	x         int
	y         int
	rowspan   int
	colspan   int
	cl        *elements.ComponentsList
}

// Correctly initialize a cell
func NewCell(c *elements.Component, cl *elements.ComponentsList) *cell {
	return &cell{component: c, width: 100, height: 100, rowspan: 1, colspan: 1, cl: cl}
}

// Set with
func (c *cell) SetWidth(w int) {
	c.width = w
}

// Set height
func (c *cell) SetHeight(h int) {
	c.height = h
}

func (c *cell) SetX(x int) {
	c.x = x
}

func (c *cell) SetY(y int) {
	c.y = y
}

// Get width (respects colspan)
func (c *cell) Width() int {
	return c.width * c.colspan
}

// Get height (respects rowspan)
func (c *cell) Height() int {
	return c.height * c.rowspan
}

// Number of rows
func (c *cell) Rowspan(r int) {
	c.rowspan = c.rowspan + r
}

// Number of columns
func (c *cell) Colspan(col int) {
	c.colspan = c.colspan + col
}

// Get relativ y coords
func (c *cell) Y() int {
	return c.y
}

// Get relativ x coords
func (c *cell) X() int {
	return c.x
}

// Returns the containing component
func (c *cell) Component() *elements.Component {
	return c.component
}
