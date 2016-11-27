package main

import (
	"github.com/iwalz/tdoc/lexer"
)

func main() {
	lexer.TdocParse(lexer.NewLexer("cloud foo"))
}
