package parser

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleComponent(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer("cloud foo"))
	ast := p.AST()
	assert.Equal(t, "*parser.ProgramNode", reflect.TypeOf(ast).String())
	d := ast.Front()
	assert.Equal(t, "*parser.DefaultNode", reflect.TypeOf(d).String())
	c := d.Front()
	assert.Equal(t, "*parser.ComponentNode", reflect.TypeOf(c).String())
	assert.Equal(t, "foo", c.(*ComponentNode).Identifier)
	assert.Equal(t, "cloud", c.(*ComponentNode).Component)
}

func TestAliasComponent(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer("cloud foo as bar"))
	ast := p.AST()
	assert.Equal(t, "*parser.ProgramNode", reflect.TypeOf(ast).String())
	d := ast.Front()
	assert.Equal(t, "*parser.DefaultNode", reflect.TypeOf(d).String())
	a := d.Front()
	assert.Equal(t, "*parser.AliasNode", reflect.TypeOf(a).String())
	assert.Equal(t, "bar", a.(*AliasNode).Alias)
	c := a.Front()
	assert.Equal(t, "foo", c.(*ComponentNode).Identifier)
	assert.Equal(t, "cloud", c.(*ComponentNode).Component)
}
