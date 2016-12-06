%{

package parser

import (
  "fmt"
  "github.com/iwalz/tdoc/elements"
)

var Program elements.Element

const debug = false

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

// Statement declaration, only do add here
program: statement_list
{
  if debug {
    fmt.Println("program")
  }
}
;
statement_list: statement
{
  if debug {
    fmt.Println("statement_list single")
  }
  $1.Root().Add($1)
}
|
statement_list statement
{
  if debug {
    fmt.Println("statement_list multi")
  }
  $2.Parent($1.Root())
  $2.Root().Add($2)
}
;
statement: statement SCOPEIN statement SCOPEOUT
{
  if debug {
    fmt.Println("Declaration scope")
  }
  $3.Parent($1)
  $3.Root().Add($3)
}
;

statement: declaration

declaration: COMPONENT IDENTIFIER
{
  if debug {
    fmt.Println("Component", $1, $2)
  }
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
  if debug {
    fmt.Println("alias")
  }
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
