package elements

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestDefaultComponentsList(t *testing.T) {
	cl := NewComponentsList("/foo")

	mfs := afero.NewMemMapFs()
	mfs.MkdirAll("/foo", 0655)

	cl.fs = mfs
	cl.Parse()

	assert.True(t, cl.Exists("node"))
	assert.True(t, cl.Exists("actor"))
	assert.True(t, cl.Exists("cloud"))
	assert.False(t, cl.Exists("foo"))
}

func TestComponentsList(t *testing.T) {
	cl := NewComponentsList("/foo")

	mfs := afero.NewMemMapFs()
	mfs.MkdirAll("/foo", 0655)
	mfs.Create("/foo/foo.svg")
	mfs.Create("/foo/bar.svg")
	mfs.Create("/foo/blubb.svg")

	cl.fs = mfs
	cl.Parse()

	assert.True(t, cl.Exists("node"))
	assert.True(t, cl.Exists("actor"))
	assert.True(t, cl.Exists("cloud"))
	assert.True(t, cl.Exists("foo"))
	assert.True(t, cl.Exists("bar"))
	assert.True(t, cl.Exists("blubb"))
	assert.False(t, cl.Exists("baz"))
}

func TestThemedComponentsList(t *testing.T) {
	cl := NewComponentsList("/foo")

	mfs := afero.NewMemMapFs()
	mfs.MkdirAll("/foo", 0655)
	mfs.Create("/foo/foo.svg")
	mfs.Create("/foo/bar.svg")
	mfs.Create("/foo/blubb.svg")
	mfs.Create("/foo/aws/lambda.svg")
	mfs.Create("/foo/aws/ec2_instance.svg")
	mfs.Create("/foo/aws/sns.svg")

	cl.fs = mfs
	cl.Parse()

	assert.True(t, cl.Exists("node"))
	assert.True(t, cl.Exists("actor"))
	assert.True(t, cl.Exists("cloud"))
	assert.True(t, cl.Exists("foo"))
	assert.True(t, cl.Exists("bar"))
	assert.True(t, cl.Exists("blubb"))
	assert.True(t, cl.Exists("aws_lambda"))
	assert.True(t, cl.Exists("aws_ec2_instance"))
	assert.True(t, cl.Exists("aws_sns"))
	assert.False(t, cl.Exists("baz"))
	assert.False(t, cl.Exists("aws_foo"))
}
