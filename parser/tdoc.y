%{

package parser

import (
  "fmt"
  "github.com/iwalz/tdoc/elements"
)

var program elements.Element

const debug = true

%}

%token <token> SCOPEIN SCOPEOUT
%token <val> COMPONENT TEXT ERROR IDENTIFIER ALIAS
%type <element> statement_list statement declaration
%type <element> program

%union{
  val string
  pos int
  line int
  token int
  element elements.Element
}

%%

// Statement declaration, only do add here
program: statement_list
{
  if debug {
    fmt.Println("program")
  }
  $$ = program
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
statement: statement SCOPEIN statement_list SCOPEOUT
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
  if program == nil {
    program = elements.NewMatrix(nil)
  }

  if $$.Root() == nil {
    $$.Parent(program)
  }
}
| COMPONENT IDENTIFIER ALIAS TEXT
{
  if debug {
    fmt.Println("alias")
  }
  $$ = elements.NewComponent(nil, nil, $1, $2, $4)
  if program == nil {
    program = elements.NewMatrix(nil)
  }

  if $$.Root() == nil {
    $$.Parent(program)
  }
}
;

%% /* Start of the program */

func (p *TdocParserImpl) AST() elements.Element {
  return program
}
