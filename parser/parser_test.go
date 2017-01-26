package parser

import (
	"reflect"
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/iwalz/tdoc/elements"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetLevel(log.WarnLevel)
}

func TestSimpleComponent(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer("cloud abc", elements.NewComponentsList("")))
	ast := p.AST()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())
	c := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "abc", c.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c.(*elements.Component).Typ)
}

func TestManyComponents(t *testing.T) {
	p := &TdocParserImpl{}

	p.Parse(NewLexer("cloud foo1 as bar1 cloud foo2 as bar2 cloud foo3 as bar3 cloud foo4 as bar4", elements.NewComponentsList("")))
	ast := p.AST()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())
	c1 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c1).String())
	assert.Equal(t, "foo1", c1.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c1.(*elements.Component).Typ)
	assert.Equal(t, "bar1", c1.(*elements.Component).Alias)

	c2 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c2).String())
	assert.Equal(t, "foo2", c2.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c2.(*elements.Component).Typ)
	assert.Equal(t, "bar2", c2.(*elements.Component).Alias)

	c3 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c3).String())
	assert.Equal(t, "foo3", c3.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c3.(*elements.Component).Typ)
	assert.Equal(t, "bar3", c3.(*elements.Component).Alias)

	c4 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c4).String())
	assert.Equal(t, "foo4", c4.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c4.(*elements.Component).Typ)
	assert.Equal(t, "bar4", c4.(*elements.Component).Alias)
}

func TestAliasComponent(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer("cloud foo as bar", elements.NewComponentsList("")))
	ast := p.AST()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())
	c := ast.Next()
	assert.Equal(t, "foo", c.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c.(*elements.Component).Typ)
	assert.Equal(t, "bar", c.(*elements.Component).Alias)
}

func TestRecursiveAliasComponent(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer("cloud foo as bar node blubb as baz", elements.NewComponentsList("")))
	ast := p.AST()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())

	c := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "cloud", c.(*elements.Component).Typ)
	assert.Equal(t, "foo", c.(*elements.Component).Identifier)
	assert.Equal(t, "bar", c.(*elements.Component).Alias)

	c1 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c1).String())
	assert.Equal(t, "node", c1.(*elements.Component).Typ)
	assert.Equal(t, "blubb", c1.(*elements.Component).Identifier)
	assert.Equal(t, "baz", c1.(*elements.Component).Alias)
}

func TestNewlineScopedComponent(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer(`cloud foo {
		actor blubb
		actor baz
	}
	`, elements.NewComponentsList("")))
	ast := p.AST()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())
	c := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "foo", c.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c.(*elements.Component).Typ)
	c1 := c.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c1).String())
	assert.Equal(t, "blubb", c1.(*elements.Component).Identifier)
	assert.Equal(t, "actor", c1.(*elements.Component).Typ)
	c2 := c.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c2).String())
	assert.Equal(t, "baz", c2.(*elements.Component).Identifier)
	assert.Equal(t, "actor", c2.(*elements.Component).Typ)
}

func TestScopedComponent(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer("cloud foo as bar { actor blubb as baz }", elements.NewComponentsList("")))
	ast := p.AST()

	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())
	c := ast.Next()
	c.Reset()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "foo", c.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c.(*elements.Component).Typ)
	c1 := c.Next()
	c.Reset()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c1).String())
	assert.Equal(t, "blubb", c1.(*elements.Component).Identifier)
	assert.Equal(t, "actor", c1.(*elements.Component).Typ)
}

func TestAliasScopedComponent(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer("cloud foo as bar { actor blubb as baz }", elements.NewComponentsList("")))
	ast := p.AST()

	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())
	c := ast.Next()
	c.Reset()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "foo", c.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c.(*elements.Component).Typ)
	assert.Equal(t, "bar", c.(*elements.Component).Alias)

	c1 := c.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c1).String())
	assert.Equal(t, "blubb", c1.(*elements.Component).Identifier)
	assert.Equal(t, "actor", c1.(*elements.Component).Typ)
	assert.Equal(t, "baz", c1.(*elements.Component).Alias)
}

func TestMultiNestedComponent(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer("cloud foo as bar{   actor blubb as baz    {node foo as quo}}", elements.NewComponentsList("")))
	ast := p.AST()

	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())
	c := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "foo", c.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c.(*elements.Component).Typ)
	assert.Equal(t, "bar", c.(*elements.Component).Alias)

	c1 := c.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c1).String())
	assert.Equal(t, "blubb", c1.(*elements.Component).Identifier)
	assert.Equal(t, "actor", c1.(*elements.Component).Typ)
	assert.Equal(t, "baz", c1.(*elements.Component).Alias)

	c2 := c1.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c2).String())
	assert.Equal(t, "foo", c2.(*elements.Component).Identifier)
	assert.Equal(t, "node", c2.(*elements.Component).Typ)
	assert.Equal(t, "quo", c2.(*elements.Component).Alias)
}

func TestComplexMultiNestedComponent(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer(`
		cloud foo as bar1
		cloud foo as bar2 {
			cloud blubb as baz1
			actor blubb as baz2 {
				node foo as quo
				node blubb as quo1
				node blubb as quo2
			}
		}`, elements.NewComponentsList("")))
	ast := p.AST()

	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())
	c := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "foo", c.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c.(*elements.Component).Typ)
	assert.Equal(t, "bar1", c.(*elements.Component).Alias)

	c1 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c1).String())
	assert.Equal(t, "foo", c1.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c1.(*elements.Component).Typ)
	assert.Equal(t, "bar2", c1.(*elements.Component).Alias)

	c2 := c1.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c2).String())
	assert.Equal(t, "blubb", c2.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c2.(*elements.Component).Typ)
	assert.Equal(t, "baz1", c2.(*elements.Component).Alias)

	c3 := c1.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c3).String())
	assert.Equal(t, "blubb", c3.(*elements.Component).Identifier)
	assert.Equal(t, "actor", c3.(*elements.Component).Typ)
	assert.Equal(t, "baz2", c3.(*elements.Component).Alias)

	c4 := c3.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c4).String())
	assert.Equal(t, "foo", c4.(*elements.Component).Identifier)
	assert.Equal(t, "node", c4.(*elements.Component).Typ)
	assert.Equal(t, "quo", c4.(*elements.Component).Alias)

	c5 := c3.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c5).String())
	assert.Equal(t, "blubb", c5.(*elements.Component).Identifier)
	assert.Equal(t, "node", c5.(*elements.Component).Typ)
	assert.Equal(t, "quo1", c5.(*elements.Component).Alias)

	c6 := c3.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c6).String())
	assert.Equal(t, "blubb", c6.(*elements.Component).Identifier)
	assert.Equal(t, "node", c6.(*elements.Component).Typ)
	assert.Equal(t, "quo2", c6.(*elements.Component).Alias)
}

func TestSimpleRelation(t *testing.T) {
	p := &TdocParserImpl{}

	p.Parse(NewLexer(`cloud foo
		cloud bar
		foo --> bar`, elements.NewComponentsList("")))
	ast := p.AST()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())
	c := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "foo", c.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c.(*elements.Component).Typ)
	relations := c.Relations()
	r := relations[0]
	assert.NotNil(t, r)

	c1 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "bar", c1.(*elements.Component).Identifier)
	assert.Equal(t, "cloud", c1.(*elements.Component).Typ)
}

func TestSimpleRelationAlias(t *testing.T) {
	p := &TdocParserImpl{}

	p.Parse(NewLexer(`cloud foo as a1
		cloud bar as a2
		a1 --> a2`, elements.NewComponentsList("")))
	ast := p.AST()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())
	c := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "a1", c.(*elements.Component).Alias)
	assert.Equal(t, "cloud", c.(*elements.Component).Typ)
	relations := c.Relations()
	r := relations[0]
	assert.NotNil(t, r)

	c1 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "a2", c1.(*elements.Component).Alias)
	assert.Equal(t, "cloud", c1.(*elements.Component).Typ)
}

func TestSimpleRelationDeclaration(t *testing.T) {
	p := &TdocParserImpl{}

	p.Parse(NewLexer(`cloud foo as a1 --> cloud bar as a2`, elements.NewComponentsList("")))
	ast := p.AST()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())
	c := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "a1", c.(*elements.Component).Alias)
	assert.Equal(t, "cloud", c.(*elements.Component).Typ)
	relations := c.Relations()
	r := relations[0]
	assert.NotNil(t, r)

	c1 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "a2", c1.(*elements.Component).Alias)
	assert.Equal(t, "cloud", c1.(*elements.Component).Typ)

	assert.Equal(t, c1, r.Element())
}

func TestMultiRelationDeclaration(t *testing.T) {
	p := &TdocParserImpl{}

	p.Parse(NewLexer(`cloud foo as a1 --> cloud bar as a2 --> actor baz as a3 --> node blubb as a4`, elements.NewComponentsList("")))
	ast := p.AST()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())
	c := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "a1", c.(*elements.Component).Alias)
	assert.Equal(t, "cloud", c.(*elements.Component).Typ)
	relations := c.Relations()
	r := relations[0]
	assert.NotNil(t, r)

	c1 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c1).String())
	assert.Equal(t, "a2", c1.(*elements.Component).Alias)
	assert.Equal(t, "cloud", c1.(*elements.Component).Typ)

	assert.Equal(t, c1, r.Element())

	relations = c1.Relations()
	r = relations[0]
	assert.NotNil(t, r)

	c2 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c2).String())
	assert.Equal(t, "a3", c2.(*elements.Component).Alias)
	assert.Equal(t, "actor", c2.(*elements.Component).Typ)

	assert.Equal(t, c2, r.Element())
	relations = c2.Relations()
	r = relations[0]
	assert.NotNil(t, r)

	c3 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c3).String())
	assert.Equal(t, "a4", c3.(*elements.Component).Alias)
	assert.Equal(t, "node", c3.(*elements.Component).Typ)

	assert.Equal(t, c3, r.Element())

	c4 := ast.Next()
	assert.Nil(t, c4)
}

func TestAliasRelationDeclaration(t *testing.T) {
	p := &TdocParserImpl{}

	p.Parse(NewLexer(`cloud foo as a1
		a1 --> cloud bar as a2`, elements.NewComponentsList("")))
	ast := p.AST()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())
	c := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "a1", c.(*elements.Component).Alias)
	assert.Equal(t, "cloud", c.(*elements.Component).Typ)
	relations := c.Relations()
	r := relations[0]
	assert.NotNil(t, r)

	c1 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c1).String())
	assert.Equal(t, "a2", c1.(*elements.Component).Alias)
	assert.Equal(t, "cloud", c1.(*elements.Component).Typ)

	assert.Equal(t, c1, r.Element())
}

func TestComplexAliasRelationDeclaration(t *testing.T) {
	p := &TdocParserImpl{}

	p.Parse(NewLexer(`cloud foo as a1
		a1 --> cloud bar as a2 {
			actor bar as a3
		} --> node baz`, elements.NewComponentsList("")))
	ast := p.AST()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(ast).String())

	c := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c).String())
	assert.Equal(t, "a1", c.(*elements.Component).Alias)
	assert.Equal(t, "cloud", c.(*elements.Component).Typ)
	relations := c.Relations()
	r := relations[0]
	assert.NotNil(t, r)

	c1 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c1).String())
	assert.Equal(t, "a2", c1.(*elements.Component).Alias)
	assert.Equal(t, "cloud", c1.(*elements.Component).Typ)

	assert.Equal(t, c1, r.Element())
	relations = c1.Relations()
	r = relations[0]
	assert.NotNil(t, r)

	c2 := ast.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(c2).String())
	assert.Equal(t, "baz", c2.(*elements.Component).Identifier)
	assert.Equal(t, "node", c2.(*elements.Component).Typ)
	assert.Equal(t, c2, r.Element())

	a := c1.Next()
	assert.Equal(t, "*elements.Component", reflect.TypeOf(a).String())
	assert.Equal(t, "a3", a.(*elements.Component).Alias)
	assert.Equal(t, "actor", a.(*elements.Component).Typ)

	a2 := c2.Next()
	assert.Nil(t, a2)

	c3 := ast.Next()
	assert.Nil(t, c3)

}
