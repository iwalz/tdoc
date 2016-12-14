%{

package parser

import (
  "fmt"
  "github.com/iwalz/tdoc/elements"
)

var program elements.Element
var roots []elements.Element
var depth int

const debug = false

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

  for i, v := range roots {
    if i > 0 {
      v.Added(true)
      roots[i-1].Add(v)
    }
  }
  $$ = roots[0]
  program = $$
  //spew.Dump(program)
}
;
statement_list: statement
{
  if debug {
    fmt.Println("statement_list single", depth)
  }
  if depth == 0 && !$1.IsAdded() {
    $1.Added(true)
    roots[depth].Add($1)
  }
  //spew.Dump(roots[depth])
}
|
statement_list statement
{
  if debug {
    fmt.Println("statement_list multi", depth)
  }
  if $2 != nil && !$2.IsAdded() {
    $2.Added(true)
    roots[depth].Add($2)
    //spew.Dump(roots[depth])
  }
}
;
statement: statement SCOPEIN
{
  if debug {
    fmt.Println("Scope in")
  }
  depth = depth + 1
  roots = append(roots, $1)
}
statement: SCOPEOUT
{
  if debug {
    fmt.Println("Scope out")
  }
  depth = depth - 1
}
;

statement: declaration

declaration: COMPONENT IDENTIFIER
{
  if debug {
    fmt.Println("Component", $1, $2)
  }
  $$ = elements.NewComponent(nil, nil, $1, $2, "")

  if roots == nil {
    roots = make([]elements.Element, 0)
    program = elements.NewMatrix(nil)
    roots = append(roots, program)
  }
}
| COMPONENT IDENTIFIER ALIAS TEXT
{
  if debug {
    fmt.Println("alias")
  }
  $$ = elements.NewComponent(nil, nil, $1, $2, $4)
  if roots == nil {
    roots = make([]elements.Element, 0)
    program = elements.NewMatrix(nil)
    roots = append(roots, program)
  }
}
;

%% /* Start of the program */

func (p *TdocParserImpl) AST() elements.Element {
  roots = nil
  return program
}
