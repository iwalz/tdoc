%{

package parser

import (
  "github.com/iwalz/tdoc/elements"
)

var Program elements.Element
var depth int
var root []elements.Element
var comp []elements.Element

const yydebug=1

%}

%token <val> SCOPEIN SCOPEOUT
%token <val> COMPONENT TEXT ERROR IDENTIFIER ALIAS
%type <element> statement_list statement
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
  Program.Add($1)
}

statement_list: statement
{
  //fmt.Printf("statement")
  //if $1.Root() != Program {
    $1.Root().Add($1)
  //}
}
|
statement_list statement
{
  //fmt.Println("statement_list")
  $1.Root().Add($2)
}
;
statement: COMPONENT IDENTIFIER
{
  if Program == nil {
    Program = elements.NewMatrix(nil)
    root = append(root, Program)
  }
  $$ = elements.NewComponent(nil, nil, $1, $2)
  if $$.Root() == nil {
    $$.Parent(Program)
  }
}
| statement ALIAS TEXT
{
  if c, ok := $1.(*elements.Component); ok {
    c.Alias = $3
  }
}
| statement SCOPEIN statement_list SCOPEOUT
{
  $3.Parent($1)
  $$ = $3
}
;


%% /* Start of the program */

func (p *TdocParserImpl) AST() elements.Element {
  ret := Program
  Program = nil
  return ret
}
