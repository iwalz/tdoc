package table

import (
	"errors"

	"github.com/iwalz/tdoc/elements"
)

// interface for table and cell (a cell can contain a table)
type tableAbstract interface {
	Component() *elements.Component
	Height() int
	Width() int
	X() int
	Y() int
}

// Errors
var ErrCellNotEmpty = errors.New("Cell not empty")
var ErrIndexOutOfBounds = errors.New("Index out of bounds")

// Table representation
type Table struct {
	cells [][]tableAbstract
	x     int
	y     int
}

// Satisfy tableAbstract interface
func (t *Table) Component() *elements.Component {
	return nil
}

// Get width
func (t *Table) Width() int {
	return t.Columns() * 100
}

// Get height
func (t *Table) Height() int {
	return t.Rows() * 100
}

// Get rows
func (t *Table) Rows() int {
	return len(t.cells[0])
}

// Get columns
func (t *Table) Columns() int {
	return len(t.cells)
}

// Returns and initializes and empty table
func NewTable() *Table {
	cells := make([][]tableAbstract, 1)
	t := &Table{cells: cells}
	return t
}

// Increase the internal data structure to x:y
func (t *Table) increaseTo(x int, y int) {
	// Make sure first dimension works
	for i := 0; i < x; i++ {
		if len(t.cells) < x {
			var rows []tableAbstract
			// Set second dimension
			for index := 0; index < y; index++ {
				rows = append(rows, nil)
			}
			t.cells = append(t.cells, rows)
		}

		// Do this only if length doesn't match
		if len(t.cells[i]) < y {
			for f := 0; f < y; f++ {
				// And only for elements that doesn't match
				if len(t.cells[i]) < y {
					t.cells[i] = append(t.cells[i], nil)
				}
			}
		}
	}
}

// Add finds the next free slot and adds a component there.
// Increases the table if no free slot is available
func (t *Table) Add(c *elements.Component) {
	rowCount := t.Rows()
	columnCount := t.Columns()
	x, y := t.findFreeSlot()

	// add col if row and col count is equal
	if (x == 0 && y == 0) && t.Rows() == t.Columns() {
		rowCount = t.Rows()
		columnCount = t.Columns() + 1
	}

	if (x == 0 && y == 0) && t.Rows() < t.Columns() {
		rowCount = t.Rows() + 1
		columnCount = t.Columns()
	}

	if (x == 0 && y == 0) && (t.Rows() != rowCount || t.Columns() != columnCount) {
		t.increaseTo(columnCount, rowCount)
		x, y = t.findFreeSlot()
	}

	t.AddTo(x, y, c)
}

// Explicit add to x:y
func (t *Table) AddTo(x int, y int, c *elements.Component) error {
	t.increaseTo(x, y)
	// check if cell is already used
	r := t.cells[x-1][y-1]
	if r != nil {
		return ErrCellNotEmpty
	}
	// table starts at 1:1, slice at 0:0
	t.cells[x-1][y-1] = NewCell(c)

	return nil
}

// Retries an element from pos x:y
func (t *Table) GetFrom(x int, y int) (tableAbstract, error) {
	if len(t.cells) < x {
		return nil, ErrIndexOutOfBounds
	}

	if len(t.cells[x-1]) < y {
		return nil, ErrIndexOutOfBounds
	}

	return t.cells[x-1][y-1], nil
}

// Identifies next free slot in Table
func (t *Table) findFreeSlot() (int, int) {
	//fmt.Println("Find free slot")
	for x, vx := range t.cells {
		for y, vy := range vx {
			//spew.Dump(vy)
			if vy == nil {
				return x + 1, y + 1
			}
		}
	}

	return 0, 0
}
