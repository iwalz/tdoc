package table

import (
	"testing"

	"github.com/iwalz/tdoc/elements"
	"github.com/stretchr/testify/assert"
)

func TestRowAndColumn(t *testing.T) {
	table := NewTable()
	table.AddColumn()
	table.AddRow()

	assert.Equal(t, 2, table.Columns())
	assert.Equal(t, 2, table.Rows())
}

func TestAddTo(t *testing.T) {
	table := NewTable()
	c := elements.NewComponent("node", "foo", "bar")
	table.AddTo(1, 1, c)
	c1, err := table.GetFrom(1, 1)
	assert.Nil(t, err)

	assert.Equal(t, c, c1)

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
	table := NewTable()
	c := elements.NewComponent("node", "foo", "bar")
	table.AddTo(10, 10, c)
	c1, err := table.GetFrom(10, 10)
	assert.Nil(t, err)
	assert.Equal(t, c, c1)

	assert.Equal(t, 10, len(table.cells[9]))
}
