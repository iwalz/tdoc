%{

package lexer

import (
    "fmt"
    "github.com/iwalz/tdoc/component"
)

var Components []*component.Component

%}

%token <val> COMPONENT TEXT ERROR IDENTIFIER ALIAS
%type <comp> declaration

%union{
  val string
  pos int
  line int
  token int
  comp *component.Component
}

%%

input:
  declaration
;
declaration:
  declaration ALIAS TEXT
  | COMPONENT IDENTIFIER
  {
    Components = append(Components, &component.Component{Typ: $1, Identifier: $2})
  }
;
declaration:
  COMPONENT
  {
    fmt.Println($1)
  }
  ;

%% /* Start of the program */

func (p *TdocParserImpl) AST() []*component.Component {
  return Components
}
