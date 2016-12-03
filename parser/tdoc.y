%{

package parser

import (
  "fmt"
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
%type <element> program statement_list statement

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
  //Program.Add($1)
}
statement_list: statement
{
  Program.Add($1)
}
| statement_list statement
{
  $1.Root().Add($2)
}
;
statement: COMPONENT IDENTIFIER
{
  if Program == nil {
    Program = elements.NewMatrix(nil)
    root = append(root, Program)
  }
  fmt.Println($1, $2)
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
}
;


%% /* Start of the program */

func (p *TdocParserImpl) AST() elements.Element {
  ret := Program
  Program = nil
  return ret
}
