package table

import (
	"fmt"

	svg "github.com/ajstarks/svgo"
	"github.com/iwalz/tdoc/elements"
	"github.com/iwalz/tdoc/image"
	"github.com/spf13/afero"
)

type SVG struct {
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
	Doc    string `xml:",innerxml"`
}

// Cell dimension
var Dimension = 120

// Border dimension
var Border = 40

// Renders the pic at its location
func (c *cell) Render(svg *svg.SVG) error {

	filename := c.cl.GetFilenameByType(c.Component())
	if filename == "" {
		return nil
	}
	f, err := c.fs.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	c.rewrite.SetX(c.X() + 10)
	c.rewrite.SetY(c.Y())
	c.rewrite.SetWidth(100)
	c.rewrite.SetHeight(100)
	c.rewrite.SetName(c.Component().Typ)

	err = c.rewrite.Place(svg, f)
	if err != nil {
		return err
	}
	svg.Group("", fmt.Sprintf(`transform="translate(%d,%d)"`, c.X(), c.Y()+100))

	svg.Text(60, 10, c.Component().Identifier, `text-anchor="middle" alignment-baseline="central"`)
	svg.Gend()

	if Wireframe {
		// Renders the clipPath wireframe
		svg.Rect(c.X(), c.Y(), Dimension, Dimension, wireoptions)
	}
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
	cl        elements.ComponentsList
	fs        afero.Fs
	rewrite   image.Rewriter
}

func (c *cell) SetRewriter(r image.Rewriter) {
	c.rewrite = r
}

// Correctly initialize a cell
func NewCell(c *elements.Component, cl elements.ComponentsList) *cell {
	fs := afero.NewOsFs()
	return &cell{component: c, width: Dimension, height: Dimension, rowspan: 1, colspan: 1, cl: cl, fs: fs}
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
