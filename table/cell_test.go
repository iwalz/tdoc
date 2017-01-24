package table

import (
	"testing"

	"github.com/iwalz/tdoc/elements"
	"github.com/stretchr/testify/assert"
)

func TestCellInitialize(t *testing.T) {
	c := elements.NewComponent("foo", "bar", "blubb")
	cl := elements.NewComponentsList("")
	cell := NewCell(c, cl)
	assert.Equal(t, c, cell.component)

	assert.Equal(t, 100, cell.Height())
	assert.Equal(t, 100, cell.Width())

	cell.SetHeight(150)
	cell.SetWidth(150)

	assert.Equal(t, 150, cell.Height())
	assert.Equal(t, 150, cell.Width())
}

func TestColAndRowspan(t *testing.T) {
	c := elements.NewComponent("foo", "bar", "blubb")
	cl := elements.NewComponentsList("")
	cell := NewCell(c, cl)
	assert.Equal(t, c, cell.component)
	cell.Rowspan(1)
	cell.Colspan(1)
	cell.x = 10
	cell.y = 10
	assert.Equal(t, 200, cell.Width())
	assert.Equal(t, 200, cell.Height())
	assert.Equal(t, 10, cell.X())
	assert.Equal(t, 10, cell.Y())
}
