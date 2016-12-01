package parser

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleParser(t *testing.T) {
	p := &TdocParserImpl{}
	p.Parse(NewLexer("cloud foo"))
	foo := p.AST()
	assert.Equal(t, "*parser.ProgramNode", reflect.TypeOf(foo).String())
}
