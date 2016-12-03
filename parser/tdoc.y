%{

package parser

import (
  "github.com/iwalz/tdoc/elements"
)

var Program elements.Element
var depth int
var root []elements.Element
var comp []elements.Element

%}

%token SCOPEIN SCOPEOUT
%token <val> COMPONENT TEXT ERROR IDENTIFIER ALIAS
%type <element> program statement_list statement

%union{
  val string
  pos int
  line int
  token int
  element elements.Element
}

%%

program: statement_list
{
  if _, ok := $1.(*elements.Matrix); ok {
    Program = $1
  }
  for _, v := range comp {
    Program.Add(v)
  }
  comp = make([]elements.Element, 0)
}
statement_list: statement
{
  $$ = elements.NewMatrix(nil)
}
| statement_list statement
{

}
;
statement: COMPONENT IDENTIFIER
{
  $$ = elements.NewComponent(nil, nil, $1, $2)
  comp = append(comp, $$)
}
;
statement: statement ALIAS TEXT
{
  if c, ok := $1.(*elements.Component); ok {
    c.Alias = $3
  }
}
;
statement: statement SCOPEIN statement SCOPEOUT
{
  $1.Add($3)
}
;

%% /* Start of the program */

func (p *TdocParserImpl) AST() elements.Element {
  return Program
}
