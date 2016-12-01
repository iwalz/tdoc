%{

package lexer

import (
  "github.com/iwalz/tdoc/ast"
)

var Program ast.Node

%}

%token <val> COMPONENT TEXT ERROR IDENTIFIER ALIAS
%type <node> program statement_list statement

%union{
  val string
  pos int
  line int
  token int
  node ast.Node
}

%%

program: statement_list
{
  //fmt.Println("Program")
  $$ = ast.NewProgramNode($1)
  Program = $$
  //fmt.Printf("Return: %+v\n", $$)
  //fmt.Printf("First: %+v\n", $1)
}
statement_list: statement
{
  //fmt.Println("statement_list")

  $$ = ast.NewDefaultNode($1)
  //fmt.Printf("Return: %+v\n", $$)
  //fmt.Printf("First: %+v\n", $1)
}
| statement statement_list
{
  //fmt.Println("statement_list alt")
  $$ = ast.NewListNode($1, $2)
  //fmt.Printf("Return: %+v\n", $$)
  //fmt.Printf("First: %+v\n", $1)
  //fmt.Printf("Second: %+v\n", $2)
}
;
statement: COMPONENT IDENTIFIER
{
  //fmt.Println("statement")
  $$ = ast.NewComponentNode(nil, nil, $1, $2)
  //fmt.Printf("Return: %+v\n", $$)
  //fmt.Printf("First: %+v\n", $1)
  //fmt.Printf("Second: %+v\n", $2)
}
;
statement: statement ALIAS TEXT
{
  //fmt.Println("alias")
  $$ = ast.NewAliasNode($1, $3)
  //fmt.Printf("Return: %+v\n", $$)
  //fmt.Printf("First: %+v\n", $1)
  //fmt.Printf("Second: %+v\n", $2)
  //fmt.Printf("Third: %+v\n", $3)
}
;

%% /* Start of the program */

func (p *TdocParserImpl) AST() ast.Node {
  return Program
}
