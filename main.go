package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/iwalz/tdoc/parser"
)

func main() {
	p := &parser.TdocParserImpl{}
	p.Parse(parser.NewLexer("cloud foo"))
	foo := p.AST()
	scan(foo)
	spew.Dump(foo)
}

func scan(node parser.Node) {
	if node == nil {
		return
	}

	for e := node.Front(); e != nil; e = node.Next() {
		print(e)
	}

	for e := node.Front(); e != nil; e = node.Next() {
		scan(e)
	}
}

func print(node parser.Node) {
	switch node.(type) {
	case *parser.ComponentNode:
		fmt.Printf("Component: %+v\n", node.(*parser.ComponentNode).Component)
		fmt.Printf("Identifier: %+v\n", node.(*parser.ComponentNode).Identifier)
	case *parser.AliasNode:
		fmt.Printf("Alias: %+v\n", node.(*parser.AliasNode).Alias)
	default:

	}
}
