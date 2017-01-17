%{

package parser

import (
  "fmt"
  "github.com/iwalz/tdoc/elements"
  "github.com/davecgh/go-spew/spew"
)

var program *elements.Component
var roots [][]*elements.Component
var depth int
var depths []int
var registry *elements.Registry

const debug = false

%}

%token <token> SCOPEIN SCOPEOUT
%token <val> COMPONENT TEXT ERROR IDENTIFIER ALIAS RELATION
%type <component> declaration relation_assignment statement statement_list program

%union{
  val string
  pos int
  line int
  token int
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

  $$ = roots[0][0]
  program = $$
  //spew.Dump(program)
}
;
statement_list: statement
{
  if debug {
    fmt.Println("statement_list single", depth, $1)
  }
  if depth == 0 && !$1.IsAdded() {
    $1.Added(true)
    sub_depth := depths[depth]
    roots[depth][sub_depth].Add($1)
  }
  //spew.Dump(roots[depth])
}
|
statement_list statement
{
  if debug {
    fmt.Println("statement_list multi", depth, $1, $2)
  }
  if $2 != nil && !$2.IsAdded() {
    $2.Added(true)
    sub_depth := depths[depth]
    roots[depth][sub_depth].Add($2)
    //spew.Dump(roots[depth])
  }
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
TEXT RELATION declaration
{
  if debug {
    fmt.Println("TEXT RELATION declaration", $1, $2, $3)
  }
  rel, _ := elements.NewRelation($2)
  rel.To($3)
  elements.Get(registry, $1).AddRelation(rel)
  if !$3.IsAdded() {
    $3.Added(true)
    sub_depth := depths[depth]
    roots[depth][sub_depth].Add($3)
  }
  $$ = $3
}
|
declaration RELATION TEXT
{
  c := elements.Get(registry, $3)
  rel, _ := elements.NewRelation($2)
  rel.To(c)
  $1.AddRelation(rel)
  if !c.IsAdded() {
    c.Added(true)
    sub_depth := depths[depth]
    roots[depth][sub_depth].Add(c)
  }
  $$ = c
}
|
relation_assignment RELATION declaration
{
  if debug {
    fmt.Println("relation_assignment RELATION declaration", $1, $3)
  }
  rel, _ := elements.NewRelation($2)
  rel.To($3)
  if !$3.IsAdded() {
    $3.Added(true)
    sub_depth := depths[depth]
    roots[depth][sub_depth].Add($3)
  }
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
  if !$1.IsAdded() {
    $1.Added(true)
    sub_depth := depths[depth]
    roots[depth][sub_depth].Add($1)
  }

  if !$3.IsAdded() {
    $3.Added(true)
    sub_depth := depths[depth]
    roots[depth][sub_depth].Add($3)
  }

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
    program = elements.NewComponent("", "", "")
    roots = make([][]*elements.Component, 0)
    sub := make([]*elements.Component, 0)
    sub = append(sub, program)
    roots = append(roots, sub)
  }

  if depths == nil {
    depths = make([]int, 1)
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
    program = elements.NewComponent("", "", "")
    roots = make([][]*elements.Component, 0)
    sub := make([]*elements.Component, 0)
    sub = append(sub, program)
    roots = append(roots, sub)
  }

  if depths == nil {
    depths = make([]int, 1)
  }

  if registry == nil {
    registry = elements.NewRegistry()
  }
  registry.Add($$)
}
|
declaration SCOPEIN
{
  if debug {
    fmt.Println("Scope in")
  }
  sub := make([]*elements.Component, 0)
  sub = append(sub, $1)
  roots = append(roots, sub)

  sub_depth := depths[depth]
  fmt.Println("Scope in")
  depths = append(depths, 0)
  fmt.Println("Depth: ", depth)
  fmt.Println("Index: ", depths[depth])
  roots[depth] = append(roots[depth], $1)
  roots[depth][sub_depth].Add($1)

  $1.Added(true)
  spew.Dump(roots)

  depths[depth] = depths[depth] + 1
  depth = depth + 1

  //roots[depth].Add($1)
}
|
SCOPEOUT
{
  if debug {
    fmt.Println("Scope out")
  }
  fmt.Println("Depth: ", depth)
  fmt.Println("Index: ", depths[depth])
  $$ = roots[depth][depths[depth]]
  depth = depth - 1
}
;
;

%% /* Start of the program */

func (p *TdocParserImpl) AST() *elements.Component {
  roots = nil
  registry = nil
  depths = nil
  return program
}
