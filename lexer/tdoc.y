%{

package lexer

import "fmt"

%}

%token COMPONENT TEXT ERROR

%union{
  val string
  pos int
  line int
  token int
}

%%

program:
  declaration
;
declaration:
  COMPONENT TEXT
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
