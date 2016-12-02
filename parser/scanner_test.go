package parser

import (
	"strings"
	"testing"

	"github.com/iwalz/tdoc/elements"
	"github.com/stretchr/testify/assert"
)

func TestSimpleScannerReturn(t *testing.T) {
	s, err := NewScanner(strings.NewReader("cloud foo as bar"))
	m := s.GetMatrix()

	e := elements.NewMatrix()
	c := &elements.Component{
		Identifier: "foo",
		Typ:        "cloud",
		Alias:      "bar",
	}
	e.Add(c)
	assert.Equal(t, nil, err)
	assert.Equal(t, e, m)
}

func TestMultipleScannerReturn(t *testing.T) {
	s, err := NewScanner(strings.NewReader("cloud foo as bar node blubb as baz"))
	m := s.GetMatrix()

	e := elements.NewMatrix()
	c := &elements.Component{
		Identifier: "foo",
		Typ:        "cloud",
		Alias:      "bar",
	}
	n := &elements.Component{
		Identifier: "blubb",
		Typ:        "node",
		Alias:      "baz",
	}
	e.Add(c)
	e.Add(n)
	assert.Equal(t, nil, err)
	assert.Equal(t, e, m)
}

func TestQuotedIdentifierScannerReturn(t *testing.T) {
	s, err := NewScanner(strings.NewReader("cloud 'foo bar blubb' as baz"))
	m := s.GetMatrix()

	e := elements.NewMatrix()
	c := &elements.Component{
		Identifier: "foo bar blubb",
		Typ:        "cloud",
		Alias:      "baz",
	}
	e.Add(c)
	assert.Equal(t, nil, err)
	assert.Equal(t, e, m)
}
