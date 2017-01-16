package table

import "github.com/iwalz/tdoc/elements"

// Represents a cell in a table
type cell struct {
	component *elements.Component
	width     int
	height    int
	x         int
	y         int
	rowspan   int
	colspan   int
}

// Correctly initialize a cell
func NewCell(c *elements.Component) *cell {
	return &cell{component: c, width: 100, height: 100, rowspan: 1, colspan: 1}
}

// Set with
func (c *cell) SetWidth(w int) {
	c.width = w
}

// Set height
func (c *cell) SetHeight(h int) {
	c.height = h
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
