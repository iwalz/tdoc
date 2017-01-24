package table

import (
	"testing"

	"github.com/iwalz/tdoc/elements"
	"github.com/stretchr/testify/assert"
)

func TestAddTo(t *testing.T) {
	cl := elements.NewComponentsList("")
	table := NewTable(cl)
	c := elements.NewComponent("node", "foo", "bar")
	table.AddTo(1, 1, c)
	c1, err := table.GetFrom(1, 1)
	assert.Nil(t, err)

	assert.Equal(t, c, c1.Component())

	c2, err := table.GetFrom(1, 2)
	assert.Equal(t, err, ErrIndexOutOfBounds)
	assert.Nil(t, c2)

	err = table.AddTo(1, 1, c)
	assert.Equal(t, err, ErrCellNotEmpty)

	c3, err := table.GetFrom(2, 1)
	assert.Equal(t, err, ErrIndexOutOfBounds)
	assert.Nil(t, c3)
}

func TestHigherAddTo(t *testing.T) {
	cl := elements.NewComponentsList("")
	table := NewTable(cl)
	c := elements.NewComponent("node", "foo", "bar")
	table.AddTo(10, 10, c)
	c1, err := table.GetFrom(10, 10)
	assert.Nil(t, err)
	assert.Equal(t, c, c1.Component())

	assert.Equal(t, 10, len(table.cells[9]))
}

func TestFindFreeSlot(t *testing.T) {
	cl := elements.NewComponentsList("")
	table := NewTable(cl)
	c := elements.NewComponent("node", "foo", "bar")
	table.AddTo(2, 2, c)
	c1, err := table.GetFrom(2, 2)
	assert.Nil(t, err)
	assert.Equal(t, c, c1.Component())

	assert.Equal(t, 2, len(table.cells[1]))
	x, y := table.findFreeSlot()
	assert.Equal(t, 1, x)
	assert.Equal(t, 1, y)
}

func TestMoreDimensionFindFreeSlot(t *testing.T) {
	cl := elements.NewComponentsList("")
	table := NewTable(cl)
	c := elements.NewComponent("node", "foo", "bar")
	table.AddTo(1, 1, c)
	c1, err := table.GetFrom(1, 1)
	assert.Nil(t, err)
	assert.Equal(t, c, c1.Component())

	x, y := table.findFreeSlot()
	assert.Equal(t, 0, x)
	assert.Equal(t, 0, y)

	table.increaseTo(2, 2)
	x, y = table.findFreeSlot()
	assert.Equal(t, 1, x)
	assert.Equal(t, 2, y)

	table.AddTo(1, 2, c)
	x, y = table.findFreeSlot()
	assert.Equal(t, 2, x)
	assert.Equal(t, 1, y)

	table.AddTo(2, 1, c)
	x, y = table.findFreeSlot()
	assert.Equal(t, 2, x)
	assert.Equal(t, 2, y)
	table.AddTo(2, 2, c)

	table.increaseTo(3, 3)

	x, y = table.findFreeSlot()
	assert.Equal(t, 1, x)
	assert.Equal(t, 3, y)
	table.AddTo(1, 3, c)

	x, y = table.findFreeSlot()
	assert.Equal(t, 2, x)
	assert.Equal(t, 3, y)
	table.AddTo(2, 3, c)

	x, y = table.findFreeSlot()
	assert.Equal(t, 3, x)
	assert.Equal(t, 1, y)
	table.AddTo(3, 1, c)

	x, y = table.findFreeSlot()
	assert.Equal(t, 3, x)
	assert.Equal(t, 2, y)
	table.AddTo(3, 2, c)

	x, y = table.findFreeSlot()
	assert.Equal(t, 3, x)
	assert.Equal(t, 3, y)
	table.AddTo(3, 3, c)

	x, y = table.findFreeSlot()
	assert.Equal(t, 0, x)
	assert.Equal(t, 0, y)
}

func TestAdd(t *testing.T) {
	cl := elements.NewComponentsList("")
	table := NewTable(cl)
	c := elements.NewComponent("node", "foo", "bar")

	// Add first on row 1 and col 1
	table.Add(c)
	assert.Equal(t, 1, table.Rows())
	assert.Equal(t, 1, table.Columns())
	e, err := table.GetFrom(1, 1)
	assert.Equal(t, c, e.Component())
	assert.Nil(t, err)

	// Add second on fist row and 2nd col
	table.Add(c)
	assert.Equal(t, 1, table.Rows())
	assert.Equal(t, 2, table.Columns())
	e1, err := table.GetFrom(2, 1)
	assert.Equal(t, c, e1.Component())
	assert.Nil(t, err)

	// None left
	x, y := table.findFreeSlot()
	assert.Equal(t, 0, x)
	assert.Equal(t, 0, y)

	// Third is on the 2nd row and first col
	table.Add(c)
	assert.Equal(t, c, table.cells[0][1].Component())
	e2, err := table.GetFrom(1, 2)
	assert.Equal(t, c, e2.Component())
	assert.Nil(t, err)
	assert.Equal(t, 2, table.Rows())
	assert.Equal(t, 2, table.Columns())

	// 2:2 left
	x, y = table.findFreeSlot()
	assert.Equal(t, 2, x)
	assert.Equal(t, 2, y)
	table.Add(c)
	assert.Equal(t, c, table.cells[1][1].Component())
	e3, err := table.GetFrom(2, 2)
	assert.Equal(t, c, e3.Component())
	assert.Nil(t, err)

	assert.Equal(t, 200, table.Width())
	assert.Equal(t, 200, table.Height())

	assert.Nil(t, table.Component())
}

func TestBorderCalculation(t *testing.T) {
	cl := elements.NewComponentsList("")
	table := NewTable(cl)
	table.SetBorder(BORDER)
	table.SetCaption("Foo")
	table.SetImage("bar")

	assert.Equal(t, 140, table.Width())
	assert.Equal(t, 140, table.Height())
}
