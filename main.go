package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/iwalz/tdoc/ast"
	"github.com/iwalz/tdoc/lexer"
)

func main() {
	parser := &lexer.TdocParserImpl{}
	parser.Parse(lexer.NewLexer("cloud foo as bar node test as blubb"))
	foo := parser.AST()
	scan(foo)
	spew.Dump(foo)
}

func scan(node ast.Node) {
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

func print(node ast.Node) {
	switch node.(type) {
	case *ast.ComponentNode:
		fmt.Printf("Component: %+v\n", node.(*ast.ComponentNode).Component)
		fmt.Printf("Identifier: %+v\n", node.(*ast.ComponentNode).Identifier)
	case *ast.AliasNode:
		fmt.Printf("Alias: %+v\n", node.(*ast.AliasNode).Alias)
	default:

	}
}
