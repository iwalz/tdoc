package table

import (
	"testing"

	svg "github.com/ajstarks/svgo"
	"github.com/iwalz/tdoc/elements"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddTo(t *testing.T) {
	cl := elements.NewComponentsList("")
	table := NewTable(cl)
	c := elements.NewComponent("node", "foo", "bar")
	cell := NewCell(c, cl)
	table.AddTo(1, 1, cell)
	c1, err := table.GetFrom(1, 1)
	assert.Nil(t, err)

	assert.Equal(t, c, c1.Component())

	c2, err := table.GetFrom(1, 2)
	assert.Equal(t, err, ErrIndexOutOfBounds)
	assert.Nil(t, c2)

	err = table.AddTo(1, 1, cell)
	assert.Equal(t, err, ErrCellNotEmpty)

	c3, err := table.GetFrom(2, 1)
	assert.Equal(t, err, ErrIndexOutOfBounds)
	assert.Nil(t, c3)
}

func TestHigherAddTo(t *testing.T) {
	cl := elements.NewComponentsList("")
	table := NewTable(cl)
	c := elements.NewComponent("node", "foo", "bar")
	cell := NewCell(c, cl)
	table.AddTo(10, 10, cell)
	c1, err := table.GetFrom(10, 10)
	assert.Nil(t, err)
	assert.Equal(t, c, c1.Component())

	assert.Equal(t, 10, len(table.cells[9]))
}

func TestFindFreeSlot(t *testing.T) {
	cl := elements.NewComponentsList("")
	table := NewTable(cl)
	c := elements.NewComponent("node", "foo", "bar")
	cell := NewCell(c, cl)
	table.AddTo(2, 2, cell)
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
	cell := NewCell(c, cl)
	table.AddTo(1, 1, cell)
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

	table.AddTo(1, 2, cell)
	x, y = table.findFreeSlot()
	assert.Equal(t, 2, x)
	assert.Equal(t, 1, y)

	table.AddTo(2, 1, cell)
	x, y = table.findFreeSlot()
	assert.Equal(t, 2, x)
	assert.Equal(t, 2, y)
	table.AddTo(2, 2, cell)

	table.increaseTo(3, 3)

	x, y = table.findFreeSlot()
	assert.Equal(t, 1, x)
	assert.Equal(t, 3, y)
	table.AddTo(1, 3, cell)

	x, y = table.findFreeSlot()
	assert.Equal(t, 2, x)
	assert.Equal(t, 3, y)
	table.AddTo(2, 3, cell)

	x, y = table.findFreeSlot()
	assert.Equal(t, 3, x)
	assert.Equal(t, 1, y)
	table.AddTo(3, 1, cell)

	x, y = table.findFreeSlot()
	assert.Equal(t, 3, x)
	assert.Equal(t, 2, y)
	table.AddTo(3, 2, cell)

	x, y = table.findFreeSlot()
	assert.Equal(t, 3, x)
	assert.Equal(t, 3, y)
	table.AddTo(3, 3, cell)

	x, y = table.findFreeSlot()
	assert.Equal(t, 0, x)
	assert.Equal(t, 0, y)
}

func TestAdd(t *testing.T) {
	cl := elements.NewComponentsList("")
	table := NewTable(cl)
	c := elements.NewComponent("node", "foo", "bar")
	cell := NewCell(c, cl)
	// Add first on row 1 and col 1
	table.Add(cell)
	assert.Equal(t, 1, table.Rows())
	assert.Equal(t, 1, table.Columns())
	e, err := table.GetFrom(1, 1)
	assert.Equal(t, c, e.Component())
	assert.Nil(t, err)

	// Add second on fist row and 2nd col
	table.Add(cell)
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
	table.Add(cell)
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
	table.Add(cell)
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
	c := elements.NewComponent("node", "foo", "bar")
	cell := NewCell(c, cl)
	table := NewTable(cl)
	table.SetX(1)
	table.SetY(1)
	table.SetBorder(BORDER)
	table.SetCaption("Foo")
	table.SetImage("bar")
	table.Add(cell)

	assert.Equal(t, 140, table.Width())
	assert.Equal(t, 140, table.Height())
	assert.Equal(t, 1, table.X())
	assert.Equal(t, 1, table.Y())
}

type CellMock struct {
	mock.Mock
}

func (c *CellMock) Component() *elements.Component {
	args := c.Mock.Called()
	return args.Get(0).(*elements.Component)
}

func (c *CellMock) Height() int {
	return c.Mock.Called().Get(0).(int)
}

func (c *CellMock) Width() int {
	return c.Mock.Called().Get(0).(int)
}

func (c *CellMock) X() int {
	return c.Mock.Called().Get(0).(int)
}

func (c *CellMock) Y() int {
	return c.Mock.Called().Get(0).(int)
}

func (c *CellMock) SetX(x int) {
	c.Mock.Called(x)
}

func (c *CellMock) SetY(y int) {
	c.Mock.Called(y)
}

func (c *CellMock) Render(svg *svg.SVG) error {
	args := c.Mock.Called(svg)
	return args.Error(0)
}

func TestRenderCall(t *testing.T) {
	cl := elements.NewComponentsList("")
	var svg *svg.SVG
	table := NewTable(cl)
	table.increaseTo(5, 5)

	c1 := new(CellMock)
	c1.On("Render", svg).Return(nil)
	c2 := new(CellMock)
	c2.On("Render", svg).Return(nil)
	c3 := new(CellMock)
	c3.On("Render", svg).Return(nil)
	c4 := new(CellMock)
	c4.On("Render", svg).Return(nil)
	c5 := new(CellMock)
	c5.On("Render", svg).Return(nil)
	c6 := new(CellMock)
	c6.On("Render", svg).Return(nil)
	c7 := new(CellMock)
	c7.On("Render", svg).Return(nil)
	c8 := new(CellMock)
	c8.On("Render", svg).Return(nil)
	c9 := new(CellMock)
	c9.On("Render", svg).Return(nil)
	c10 := new(CellMock)
	c10.On("Render", svg).Return(nil)
	c11 := new(CellMock)
	c11.On("Render", svg).Return(nil)
	c12 := new(CellMock)
	c12.On("Render", svg).Return(nil)
	c13 := new(CellMock)
	c13.On("Render", svg).Return(nil)
	c14 := new(CellMock)
	c14.On("Render", svg).Return(nil)
	c15 := new(CellMock)
	c15.On("Render", svg).Return(nil)
	c16 := new(CellMock)
	c16.On("Render", svg).Return(nil)

	table.cells[0][0] = c1
	table.cells[0][1] = c2
	table.cells[1][0] = c3
	table.cells[1][1] = c4
	table.cells[2][0] = c5
	table.cells[2][1] = c6
	table.cells[2][2] = c7
	table.cells[3][0] = c8
	table.cells[3][1] = c9
	table.cells[3][2] = c10
	table.cells[3][3] = c11
	table.cells[4][0] = c12
	table.cells[4][1] = c13
	table.cells[4][2] = c14
	table.cells[4][3] = c15
	table.cells[4][4] = c16

	table.Render(svg)
	c1.AssertNumberOfCalls(t, "Render", 1)
	c2.AssertNumberOfCalls(t, "Render", 1)
	c3.AssertNumberOfCalls(t, "Render", 1)
	c4.AssertNumberOfCalls(t, "Render", 1)
	c5.AssertNumberOfCalls(t, "Render", 1)
	c6.AssertNumberOfCalls(t, "Render", 1)
	c7.AssertNumberOfCalls(t, "Render", 1)
	c8.AssertNumberOfCalls(t, "Render", 1)
	c9.AssertNumberOfCalls(t, "Render", 1)
	c10.AssertNumberOfCalls(t, "Render", 1)
	c11.AssertNumberOfCalls(t, "Render", 1)
	c12.AssertNumberOfCalls(t, "Render", 1)
	c13.AssertNumberOfCalls(t, "Render", 1)
	c14.AssertNumberOfCalls(t, "Render", 1)
	c15.AssertNumberOfCalls(t, "Render", 1)
	c16.AssertNumberOfCalls(t, "Render", 1)
}

func TestWidthAndHeight(t *testing.T) {
	cl := elements.NewComponentsList("")
	c1 := elements.NewComponent("node", "foo", "bar")
	c2 := elements.NewComponent("node", "foo", "bar")
	c3 := elements.NewComponent("node", "foo", "bar")
	cell1 := NewCell(c1, cl)
	cell2 := NewCell(c2, cl)
	cell3 := NewCell(c3, cl)
	t1 := NewTable(cl)
	t1.SetX(1)
	t1.SetY(1)

	t2 := NewTable(cl)
	t2.SetBorder(DASHED_BORDER)
	t2.SetCaption("foo")
	t2.SetImage("bla")
	t2.SetX(1)
	t2.SetY(1)
	t2.Add(cell1)
	t2.Add(cell2)
	t2.Add(cell3)

	assert.Equal(t, 240, t2.Width())
	assert.Equal(t, 240, t2.Height())
	assert.Equal(t, 1, t2.X())
	assert.Equal(t, 1, t2.Y())
}
