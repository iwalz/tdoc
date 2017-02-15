package outputs

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/mock"
)

type OutputMock struct {
	mock.Mock
}

func (o *OutputMock) HandleDir(d string) error {
	args := o.Mock.Called(d)
	return args.Error(0)
}

func (o *OutputMock) HandleFile(d string) error {
	args := o.Mock.Called(d)
	return args.Error(0)
}

func TestExecutor(t *testing.T) {
	mfs := afero.NewMemMapFs()
	mfs.MkdirAll("/foo/bar/blubb/", 0644)
	mfs.Create("/foo/bar/blubb/1.tdoc")
	mfs.Create("2.tdoc")
	e := NewExecutor(mfs, "tdoc")

	om := new(OutputMock)
	om.On("HandleDir", "/foo/bar/blubb").Return(nil)
	om.On("HandleDir", "/foo/bar").Return(nil)
	om.On("HandleDir", "/foo").Return(nil)
	om.On("HandleFile", "/foo/bar/blubb/1.tdoc").Return(nil)
	om.On("HandleFile", "2.tdoc").Return(nil)

	e.Exec(om, []string{"/foo", "2.tdoc"})
	om.AssertNumberOfCalls(t, "HandleFile", 2)
}
