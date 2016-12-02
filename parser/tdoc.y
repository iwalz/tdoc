%{

package parser

var Program Node

%}

%token SCOPEIN SCOPEOUT
%token <val> COMPONENT TEXT ERROR IDENTIFIER ALIAS
%type <node> program statement_list statement

%union{
  val string
  pos int
  line int
  token int
  node Node
}

%%

program: statement_list
{
  //fmt.Println("Program")
  $$ = NewProgramNode($1)
  Program = $$
  //fmt.Printf("Return: %+v\n", $$)
  //fmt.Printf("First: %+v\n", $1)
}
statement_list: statement
{
  //fmt.Println("statement_list")

  //$$ = NewDefaultNode($1)
  //fmt.Printf("Return: %+v\n", $$)
  //fmt.Printf("First: %+v\n", $1)
}
| statement statement_list
{
  //fmt.Println("statement_list alt")
  $$ = NewListNode($1, $2)
  //fmt.Printf("Return: %+v\n", $$)
  //fmt.Printf("First: %+v\n", $1)
  //fmt.Printf("Second: %+v\n", $2)
}
;
statement: COMPONENT IDENTIFIER
{
  //fmt.Println("statement")
  $$ = NewComponentNode(nil, nil, $1, $2)
  //fmt.Printf("Return: %+v\n", $$)
  //fmt.Printf("First: %+v\n", $1)
  //fmt.Printf("Second: %+v\n", $2)
}
;
statement: statement ALIAS TEXT
{
  //fmt.Println("alias")
  $$ = NewAliasNode($1, $3)
  //fmt.Printf("Return: %+v\n", $$)
  //fmt.Printf("First: %+v\n", $1)
  //fmt.Printf("Second: %+v\n", $2)
  //fmt.Printf("Third: %+v\n", $3)
}
;
statement: statement SCOPEIN statement SCOPEOUT
{
  if _, ok := $1.(*AliasNode); ok {
    $1.Front().AppendChild($3)
  } else {
    $1.AppendChild($3)
  }
}
;

%% /* Start of the program */

func (p *TdocParserImpl) AST() Node {
  return Program
}
