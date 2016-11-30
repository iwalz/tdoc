package ast

import (
	"github.com/iwalz/tdoc/lexer"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type ComponentStatement struct {
	Token      lexer.TokenType
	Component  string
	Identifier string
	Alias      string
}

func (cs *ComponentStatement) statementNode() {}
func (cs *ComponentStatement) TokenLiteral() lexer.TokenType {
	return cs.Token
}
