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
	c := ast.Front()
	assert.Equal(t, "*parser.ComponentNode", reflect.TypeOf(c).String())
	assert.Equal(t, "foo", c.(*ComponentNode).Identifier)
	assert.Equal(t, "cloud", c.(*ComponentNode).Component)
}

func TestAliasComponent(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer("cloud foo as bar"))
	ast := p.AST()
	assert.Equal(t, "*parser.ProgramNode", reflect.TypeOf(ast).String())
	a := ast.Front()
	assert.Equal(t, "*parser.AliasNode", reflect.TypeOf(a).String())
	assert.Equal(t, "bar", a.(*AliasNode).Alias)
	c := a.Front()
	assert.Equal(t, "foo", c.(*ComponentNode).Identifier)
	assert.Equal(t, "cloud", c.(*ComponentNode).Component)
}

func TestRecursiveAliasComponent(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer("cloud foo as bar node blubb as baz"))
	ast := p.AST()
	assert.Equal(t, "*parser.ProgramNode", reflect.TypeOf(ast).String())
	l := ast.Front()
	assert.Equal(t, "*parser.ListNode", reflect.TypeOf(l).String())
	a := l.Front()
	assert.Equal(t, "*parser.AliasNode", reflect.TypeOf(a).String())
	assert.Equal(t, "bar", a.(*AliasNode).Alias)
	c := a.Front()
	assert.Equal(t, "*parser.ComponentNode", reflect.TypeOf(c).String())
	assert.Equal(t, "cloud", c.(*ComponentNode).Component)
	assert.Equal(t, "foo", c.(*ComponentNode).Identifier)

	a1 := l.Next()
	assert.Equal(t, "*parser.AliasNode", reflect.TypeOf(a1).String())
	assert.Equal(t, "baz", a1.(*AliasNode).Alias)
	c1 := a1.Front()
	assert.Equal(t, "*parser.ComponentNode", reflect.TypeOf(c1).String())
	assert.Equal(t, "node", c1.(*ComponentNode).Component)
	assert.Equal(t, "blubb", c1.(*ComponentNode).Identifier)

	assert.Equal(t, nil, l.Next())
}

func TestScopedComponent(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer("cloud foo { actor blubb }"))
	ast := p.AST()
	assert.Equal(t, "*parser.ProgramNode", reflect.TypeOf(ast).String())
	c := ast.Front()
	assert.Equal(t, "*parser.ComponentNode", reflect.TypeOf(c).String())
	assert.Equal(t, "foo", c.(*ComponentNode).Identifier)
	assert.Equal(t, "cloud", c.(*ComponentNode).Component)
	c1 := c.Front()
	assert.Equal(t, "*parser.ComponentNode", reflect.TypeOf(c1).String())
	assert.Equal(t, "blubb", c1.(*ComponentNode).Identifier)
	assert.Equal(t, "actor", c1.(*ComponentNode).Component)
}

func TestAliasScopedComponent(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer("cloud foo as bar { actor blubb as baz }"))
	ast := p.AST()
	assert.Equal(t, "*parser.ProgramNode", reflect.TypeOf(ast).String())
	a := ast.Front()
	assert.Equal(t, "*parser.AliasNode", reflect.TypeOf(a).String())
	assert.Equal(t, "bar", a.(*AliasNode).Alias)
	c := a.Front()
	assert.Equal(t, "*parser.ComponentNode", reflect.TypeOf(c).String())
	assert.Equal(t, "foo", c.(*ComponentNode).Identifier)
	assert.Equal(t, "cloud", c.(*ComponentNode).Component)

	a1 := c.Front()
	assert.Equal(t, "*parser.AliasNode", reflect.TypeOf(a1).String())
	assert.Equal(t, "baz", a1.(*AliasNode).Alias)
	c1 := a1.Front()
	assert.Equal(t, "*parser.ComponentNode", reflect.TypeOf(c1).String())
	assert.Equal(t, "blubb", c1.(*ComponentNode).Identifier)
	assert.Equal(t, "actor", c1.(*ComponentNode).Component)
}
