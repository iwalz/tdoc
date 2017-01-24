package table

import (
	"bytes"
	"testing"

	svg "github.com/ajstarks/svgo"
	"github.com/iwalz/tdoc/elements"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestRender(t *testing.T) {
	c := elements.NewComponent("foo", "identifier", "alias")
	cl := new(ComponentsListMock)
	cl.On("GetFilenameByType", c).Return("/foo/bar.svg")

	mfs := afero.NewMemMapFs()
	mfs.MkdirAll("/foo", 0655)
	mfs.Create("/foo/foo.svg")
	mfs.Create("/foo/bar.svg")
	afero.WriteFile(mfs, "/foo/bar.svg", []byte(`<?xml version="1.0"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN"
  "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">

<svg xmlns="http://www.w3.org/2000/svg" version="1.1"
      width="120" height="120" viewBox="0 0 236 120">
  <rect x="14" y="23" width="200" height="7" fill="lime"
      stroke="black" stroke-width="1" />
</svg>`), 0644)

	cell := NewCell(c, cl)
	cell.fs = mfs

	b := new(bytes.Buffer)
	svg := svg.New(b)
	err := cell.Render(svg)
	assert.Nil(t, err)
}

func TestErrors(t *testing.T) {
	c := elements.NewComponent("foo", "identifier", "alias")
	cl := new(ComponentsListMock)
	cl.On("GetFilenameByType", c).Return("/foo/bar.svg")

	mfs := afero.NewMemMapFs()
	mfs.MkdirAll("/foo", 0655)
	mfs.Create("/foo/bar.svg")
	afero.WriteFile(mfs, "/foo/bar.svg", []byte(``), 0644)

	cell := NewCell(c, cl)
	cell.fs = mfs

	b := new(bytes.Buffer)
	s := svg.New(b)
	err := cell.Render(s)
	assert.Error(t, err)

	mfs = afero.NewMemMapFs()
	mfs.MkdirAll("/foo", 0655)
	cell.fs = mfs

	b = new(bytes.Buffer)
	s = svg.New(b)
	err = cell.Render(s)
	assert.Error(t, err)

	c = elements.NewComponent("foo", "identifier", "alias")
	cl = new(ComponentsListMock)
	cl.On("GetFilenameByType", c).Return("")

	mfs = afero.NewMemMapFs()
	mfs.MkdirAll("/foo", 0655)
	mfs.Create("/foo/bar.svg")

	cell = NewCell(c, cl)
	cell.fs = mfs
	cell.Render(s)
}
