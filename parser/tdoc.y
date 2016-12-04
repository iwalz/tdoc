%{

package parser

import (
  "fmt"
  "github.com/iwalz/tdoc/elements"
)

var Program elements.Element

%}

%token <val> SCOPEIN SCOPEOUT
%token <val> COMPONENT TEXT ERROR IDENTIFIER ALIAS
%type <element> statement_list statement declaration
%type <val> program

%union{
  val string
  pos int
  line int
  token int
  depth int
  element elements.Element
}

%%

program: statement_list
{
  fmt.Println("program")
  //$1.Root().Add($1)
}
;
statement_list: statement
{
  fmt.Println("statement_list")
  $1.Root().Add($1)
}
|
statement_list statement
{
  fmt.Println("statement_list")
  //if $2.Root() == nil {
  $2.Parent($1.Root())
  $2.Root().Add($2)
  //}
}
|
statement SCOPEIN statement SCOPEOUT
{
  fmt.Println("Declaration scope")
  $1.Root().Add($1)
  $3.Parent($1)
  $3.Root().Add($3)
}
;

statement: declaration
{
  //fmt.Println("Statement")
  //$1.Root().Add($1)
}



declaration: COMPONENT IDENTIFIER
{
  //fmt.Println("Component", $1, $2)
  $$ = elements.NewComponent(nil, nil, $1, $2, "")
  if Program == nil {
    Program = elements.NewMatrix(nil)
  }

  if $$.Root() == nil {
    $$.Parent(Program)
  }
}
| COMPONENT IDENTIFIER ALIAS TEXT
{
  //fmt.Println("alias")
  $$ = elements.NewComponent(nil, nil, $1, $2, $4)
  if Program == nil {
    Program = elements.NewMatrix(nil)
  }

  if $$.Root() == nil {
    $$.Parent(Program)
  }
}
;


%% /* Start of the program */

func (p *TdocParserImpl) AST() elements.Element {
  ret := Program
  Program = nil
  return ret
}
