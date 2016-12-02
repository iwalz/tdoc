package main

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/iwalz/tdoc/parser"
)

func main() {
	content := "cloud foo as bar node blubb as bazt"
	scanner, err := parser.NewScanner(strings.NewReader(content))
	if err != nil {
		panic(err)
	}
	spew.Dump(scanner.GetMatrix())
}
