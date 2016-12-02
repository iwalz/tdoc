package parser

import (
	"errors"
	"io"
	"io/ioutil"

	"github.com/iwalz/tdoc/elements"
)

var ReadError = errors.New("Can't read io.Reader data")

type Scanner struct {
	reader io.Reader
	lexer  *Lexer
	parser *TdocParserImpl
	matrix *elements.Matrix
	stack  []elements.Stackable
	index  int
}

func NewScanner(r io.Reader) (*Scanner, error) {
	p := &TdocParserImpl{}
	s, e := ioutil.ReadAll(r)
	if e != nil {
		return nil, ReadError
	}
	l := NewLexer(string(s))
	m := elements.NewMatrix()
	stack := make([]elements.Stackable, 0)
	stack = append(stack, elements.Stackable(m))

	return &Scanner{
		reader: r,
		parser: p,
		lexer:  l,
		matrix: m,
		stack:  stack,
	}, nil
}

func (s Scanner) GetMatrix() *elements.Matrix {
	s.parser.Parse(s.lexer)
	return s.Scan(s.parser.AST())
}

func (s *Scanner) Scan(node Node) *elements.Matrix {
	if node == nil {
		return nil
	}

	for e := node.Front(); e != nil; e = node.Next() {
		s.convert(e, node)
	}

	for e := node.Front(); e != nil; e = node.Next() {
		s.Scan(e)
	}

	return s.matrix
}

func (s *Scanner) convert(curr, prev Node) {
	switch curr.(type) {
	case *ComponentNode:
		// Is alias defined?
		alias := ""
		if a, ok := prev.(*AliasNode); ok {
			alias = a.Alias
		}
		c := &elements.Component{
			Typ:        curr.(*ComponentNode).Component,
			Identifier: curr.(*ComponentNode).Identifier,
			Alias:      alias,
		}
		s.matrix.Add(c)
	default:
		// hmmmm
	}
}
