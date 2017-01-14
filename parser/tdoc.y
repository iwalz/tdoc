%{

package parser

import (
  "fmt"
  "github.com/iwalz/tdoc/elements"
  //"github.com/davecgh/go-spew/spew"
)

var program elements.Element
var roots []elements.Element
var depth int
var registry *elements.Registry

const debug = false

%}

%token <token> SCOPEIN SCOPEOUT
%token <val> COMPONENT TEXT ERROR IDENTIFIER ALIAS RELATION
%type <component> declaration relation_assignment statement
%type <element> statement_list program

%union{
  val string
  pos int
  line int
  token int
  element elements.Element
  component *elements.Component
  relation elements.Relation
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
  $1.Added(true)
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

statement: declaration | relation_assignment

relation_assignment: TEXT RELATION TEXT
{
  rel, _ := elements.NewRelation($2)
  rel.To(elements.Get(registry, $3))
  elements.Get(registry, $1).AddRelation(rel)
}
|
relation_assignment RELATION declaration
{
  if debug {
    fmt.Println("relation_assignment RELATION declaration", $1, $3)
  }
  rel, _ := elements.NewRelation($2)
  rel.To($3)
  $3.Added(true)
  roots[depth].Add($3)
  $1.AddRelation(rel)
  $$ = $3
}
|
declaration RELATION declaration
{
  if debug {
    fmt.Println("declaration RELATION declaration", $1, $3)
  }
  rel, _ := elements.NewRelation($2)
  rel.To($3)
  $1.Added(true)
  roots[depth].Add($1)
  $3.Added(true)
  roots[depth].Add($3)
  $1.AddRelation(rel)
  $$ = $3
}

declaration: COMPONENT IDENTIFIER
{
  if debug {
    fmt.Println("Component", $1, $2)
  }
  $$ = elements.NewComponent($1, $2, "")

  if roots == nil {
    roots = make([]elements.Element, 0)
    program = elements.NewMatrix(nil)
    roots = append(roots, program)
  }

  if registry == nil {
    registry = elements.NewRegistry()
  }
  registry.Add($$)
}
| COMPONENT IDENTIFIER ALIAS TEXT
{
  if debug {
    fmt.Println("alias")
  }
  $$ = elements.NewComponent($1, $2, $4)

  if roots == nil {
    roots = make([]elements.Element, 0)
    program = elements.NewMatrix(nil)
    roots = append(roots, program)
  }

  if registry == nil {
    registry = elements.NewRegistry()
  }
  registry.Add($$)
}
;

%% /* Start of the program */

func (p *TdocParserImpl) AST() elements.Element {
  roots = nil
  registry = nil
  return program
}
