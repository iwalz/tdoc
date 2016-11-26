%{

package main

import (
  "fmt"
)

%}

%token COMPONENT TEXT

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

func main() {
  TdocParse(NewLexer("foo", "cloud foo"))
}
