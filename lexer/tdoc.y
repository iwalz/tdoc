%{

package lexer

import "fmt"

%}

%token COMPONENT TEXT ERROR IDENTIFIER ALIAS

%union{
  val string
  pos int
  line int
  token int
}

%%

input:
  declaration
;
declaration:
  declaration ALIAS TEXT
  | COMPONENT IDENTIFIER
  {
    fmt.Println($1.val, $2.val)
  }
;
declaration:
  COMPONENT
  {
    fmt.Println($1.val)
  }
  ;

%% /* Start of the program */
