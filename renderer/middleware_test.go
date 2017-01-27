package renderer

import (
	"bytes"
	"testing"

	svg "github.com/ajstarks/svgo"
	"github.com/iwalz/tdoc/elements"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ComponentsListMock struct {
	mock.Mock
}

func (c *ComponentsListMock) GetFilenameByType(component *elements.Component) string {
	args := c.Mock.Called(component)
	return args.Get(0).(string)
}

func (c *ComponentsListMock) Parse() error {
	return c.Mock.Called().Error(0)
}

func (c *ComponentsListMock) Exists(s string) bool {
	return c.Mock.Called(s).Bool(0)
}

type TableAbstractMock struct {
	mock.Mock
}

func (t *TableAbstractMock) Component() *elements.Component {
	args := t.Mock.Called()
	return args.Get(0).(*elements.Component)
}

func (t *TableAbstractMock) Height() int {
	return t.Mock.Called().Get(0).(int)
}

func (t *TableAbstractMock) Width() int {
	return t.Mock.Called().Get(0).(int)
}

func (t *TableAbstractMock) X() int {
	return t.Mock.Called().Get(0).(int)
}

func (t *TableAbstractMock) Y() int {
	return t.Mock.Called().Get(0).(int)
}

func (t *TableAbstractMock) SetX(x int) {
	t.Mock.Called(x)
}

func (t *TableAbstractMock) SetY(y int) {
	t.Mock.Called(y)
}

func (t *TableAbstractMock) Render(svg *svg.SVG) error {
	args := t.Mock.Called(svg)
	return args.Error(0)
}

func TestRender(t *testing.T) {
	c := elements.NewComponent("node", "foo", "bar")
	c1 := elements.NewComponent("node", "foo", "bar")
	c.Add(c1)
	cl := new(ComponentsListMock)
	cl.On("GetFilenameByType", c).Return("/foo/bar.svg")

	b := new(bytes.Buffer)
	svg := svg.New(b)
	table := new(TableAbstractMock)
	table.On("Render", svg).Return(nil)
	table.On("Width").Return(100)
	table.On("Height").Return(100)

	m := NewMiddleware(c, cl)
	m.table = table
	err := m.Render(svg)
	assert.Nil(t, err)
	table.AssertNumberOfCalls(t, "Render", 1)
}

func TestScan(t *testing.T) {
	root := elements.NewComponent("", "", "")
	c1 := elements.NewComponent("node", "foo", "bar")
	c2 := elements.NewComponent("node", "foo", "bar")
	c3 := elements.NewComponent("node", "foo", "bar")
	c4 := elements.NewComponent("node", "foo", "bar")
	c5 := elements.NewComponent("node", "foo", "bar")
	c4.Add(c5)
	root.Add(c1)
	root.Add(c2)
	root.Add(c3)
	root.Add(c4)
	cl := new(ComponentsListMock)
	cl.On("GetFilenameByType", c1).Return("/foo/bar.svg")
	cl.On("GetFilenameByType", c2).Return("/foo/bar.svg")
	cl.On("GetFilenameByType", c3).Return("/foo/bar.svg")
	cl.On("GetFilenameByType", c4).Return("/foo/bar.svg")

	b := new(bytes.Buffer)
	svg := svg.New(b)
	mtable := new(TableAbstractMock)
	mtable.On("Render", svg).Return(nil)
	mtable.On("Width").Return(100)
	mtable.On("Height").Return(100)

	m := NewMiddleware(root, cl)
	table := m.Scan(root, cl)
	err := m.Render(svg)
	assert.Nil(t, err)

	assert.Equal(t, 280, table.Width())
	assert.Equal(t, 280, table.Height())
}
