%{

package parser

import (
  "fmt"
  "github.com/iwalz/tdoc/elements"
  log "github.com/Sirupsen/logrus"
  //"github.com/davecgh/go-spew/spew"
)

var program *elements.Component
var roots []*elements.Component
var depth int
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
  log.Info("program: statement_list")

  $$ = roots[0]
  program = $$
  //spew.Dump(program)
}
;
statement_list: statement
{
  log.Info("statement_list: statement")
  log.Debug("Depth", depth)
  log.Debug($1)

  if depth == 0 && !$1.IsAdded() {
    $1.Added(true)
    roots[depth].Add($1)
  }
  //spew.Dump(roots[depth])
}
|
statement_list statement
{
  log.Info("statement_list statement")
  log.Debug("Depth", depth)
  log.Debug($1)
  log.Debug($2)

  if $2 != nil && !$2.IsAdded() {
    $2.Added(true)
    roots[depth].Add($2)
    //spew.Dump(roots[depth])
  }
}
;

statement: declaration | relation_assignment

relation_assignment: TEXT RELATION TEXT
{
  log.Info("relation_assignment: TEXT RELATION TEXT")
  log.Debug($1)
  log.Debug($2)
  log.Debug($3)

  rel, _ := elements.NewRelation($2)
  rel.To(elements.Get(registry, $3))
  elements.Get(registry, $1).AddRelation(rel)
}
|
TEXT RELATION declaration
{
  log.Info("TEXT RELATION declaration")
  log.Debug($1)
  log.Debug($2)
  log.Debug($3)

  rel, _ := elements.NewRelation($2)
  rel.To($3)
  elements.Get(registry, $1).AddRelation(rel)
  if !$3.IsAdded() {
    $3.Added(true)
    roots[depth].Add($3)
  }
  $$ = $3
}
|
declaration RELATION TEXT
{
  log.Info("declaration RELATION TEXT")
  log.Debug($1)
  log.Debug($2)
  log.Debug($3)

  c := elements.Get(registry, $3)
  rel, _ := elements.NewRelation($2)
  rel.To(c)
  $1.AddRelation(rel)
  if !c.IsAdded() {
    c.Added(true)
    roots[depth].Add(c)
  }
  $$ = c
}
|
relation_assignment RELATION declaration
{
  log.Info("relation_assignment RELATION declaration")
  log.Debug($1)
  log.Debug($2)
  log.Debug($3)

  rel, _ := elements.NewRelation($2)
  rel.To($3)
  if !$3.IsAdded() {
    $3.Added(true)
    roots[depth].Add($3)
  }
  $1.AddRelation(rel)
  $$ = $3
}
|
declaration RELATION declaration
{
  log.Info("declaration RELATION declaration")
  log.Debug($1)
  log.Debug($2)
  log.Debug($3)

  if debug {
    fmt.Println("declaration RELATION declaration", $1, $3)
  }
  rel, _ := elements.NewRelation($2)
  rel.To($3)
  if !$1.IsAdded() {
    $1.Added(true)
    roots[depth].Add($1)
  }

  if !$3.IsAdded() {
    $3.Added(true)
    roots[depth].Add($3)
  }

  $1.AddRelation(rel)
  $$ = $3
}

declaration: COMPONENT IDENTIFIER
{
  log.Info("declaration: COMPONENT IDENTIFIER")
  log.Debug($1)
  log.Debug($2)
  $$ = elements.NewComponent($1, $2, "")

  if roots == nil {
    roots = make([]*elements.Component, 0)
    program = elements.NewComponent("", "", "")
    roots = append(roots, program)
  }

  if registry == nil {
    registry = elements.NewRegistry()
  }
  registry.Add($$)
}
| COMPONENT IDENTIFIER ALIAS TEXT
{
  log.Info("COMPONENT IDENTIFIER ALIAS TEXT")
  log.Debug($1)
  log.Debug($2)
  log.Debug($3)
  log.Debug($4)
  $$ = elements.NewComponent($1, $2, $4)

  if roots == nil {
    roots = make([]*elements.Component, 0)
    program = elements.NewComponent("", "", "")
    roots = append(roots, program)
  }

  if registry == nil {
    registry = elements.NewRegistry()
  }
  registry.Add($$)
}
|
declaration SCOPEIN
{
  log.Info("declaration SCOPEIN")
  log.Debug($1)
  roots[depth].Add($1)
  depth = depth + 1
  $1.Added(true)
  roots = append(roots, $1)
  //roots[depth].Add($1)
}
|
SCOPEOUT
{
  log.Info("SCOPEOUT")
  $$ = roots[depth]
  depth = depth - 1
}
;
;

%% /* Start of the program */

func (p *TdocParserImpl) AST() *elements.Component {
  roots = nil
  registry = nil
  return program
}
