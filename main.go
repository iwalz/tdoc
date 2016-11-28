package main

import (
	"fmt"

	"github.com/iwalz/tdoc/lexer"
)

func main() {
	parser := &lexer.TdocParserImpl{}
	parser.Parse(lexer.NewLexer("cloud foo"))
	ast := parser.AST()
	for _, value := range ast {
		fmt.Println(value.Identifier)
	}
}
