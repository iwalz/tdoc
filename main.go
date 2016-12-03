package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/iwalz/tdoc/parser"
)

func main() {
	content := "cloud foo as bar node blubb as bazt"
	p := &parser.TdocParserImpl{}
	l := parser.NewLexer(content)
	p.Parse(l)
	spew.Dump(p.AST())
}
