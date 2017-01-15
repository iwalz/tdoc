package table

import (
	"errors"

	"github.com/iwalz/tdoc/elements"
)

type tableAbstract interface {
}

var ErrCellNotEmpty = errors.New("Cell not empty")
var ErrIndexOutOfBounds = errors.New("Index out of bounds")

type Table struct {
	rows    int
	columns int
	cells   [][]tableAbstract
}

func (t *Table) AddRow() {
	t.rows = t.rows + 1
}

func (t *Table) AddColumn() {
	t.columns = t.columns + 1
}

func (t *Table) Rows() int {
	return t.rows
}

func (t *Table) Columns() int {
	return t.columns
}

func NewTable() *Table {
	cells := make([][]tableAbstract, 0)
	t := &Table{rows: 1, columns: 1, cells: cells}
	return t
}

func (t *Table) increaseTo(x int, y int) {
	if len(t.cells) < x {
		for i := len(t.cells); i < x; i++ {
			t.cells = append(t.cells, make([]tableAbstract, 1))
		}
	}
	t.rows = x

	if len(t.cells[x-1]) < y {
		for i := len(t.cells[x-1]); i < y; i++ {
			t.cells[x-1] = append(t.cells[x-1], nil)
		}
	}
	t.columns = y
}

func (t *Table) AddTo(x int, y int, c *elements.Component) error {
	t.increaseTo(x, y)
	// check if cell is already used
	r := t.cells[x-1][y-1]
	if r != nil {
		return ErrCellNotEmpty
	}
	// table starts at 1:1, slice at 0:0
	t.cells[x-1][y-1] = c

	return nil
}

func (t *Table) GetFrom(x int, y int) (tableAbstract, error) {
	if len(t.cells) < x {
		return nil, ErrIndexOutOfBounds
	}

	if len(t.cells[x-1]) < y {
		return nil, ErrIndexOutOfBounds
	}

	return t.cells[x-1][y-1], nil
}

func (t *Table) findFreeSlot() (int, int) {
	for x, vx := range t.cells {
		for y, vy := range vx {
			if vy == nil {
				return x + 1, y + 1
			}
		}
	}

	return 0, 0
}
