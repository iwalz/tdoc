package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode/utf8"
)

const eof = -1

type Pos int
type stateFn func(*Lexer) stateFn
type TokenType int

type Lexer struct {
	input  string // input to scan
	start  Pos    // start position of this item
	pos    Pos    // current position of the input
	width  Pos    // width of last rune read
	line   int
	tokens chan token // channel of scanned items
	state  stateFn    // The state function
}

type token struct {
	typ  TokenType // The type
	pos  Pos       // and position
	val  string    // The value
	line int       // Line
}

func (t token) String() string {
	switch t.typ {
	case 0:
		return "EOF"
	case ERROR:
		return t.val
	}
	if len(t.val) > 10 {
		return fmt.Sprintf("%.10q...", t.val)
	}

	return fmt.Sprintf("%q", t.val)
}

func (l *Lexer) Error(s string) {
	fmt.Println("syntax error: ", s, l.pos)
}

func NewLexer(input, source string) *Lexer {
	l := &Lexer{
		input:  input,
		state:  lexText,
		tokens: make(chan token, 2),
	}

	if source != "" {
		files, err := ioutil.ReadDir(source)
		if err != nil {
			log.Println("Error while reading components", err)
		}

		for _, v := range files {
			if strings.HasSuffix(v.Name(), ".svg") {
				name := strings.Replace(v.Name(), ".svg", "", 1)
				components[name] = COMPONENT
			}
		}
	}

	return l
}

// Sends a new item to the channel
func (l *Lexer) emit(t TokenType) {
	l.tokens <- token{t, l.start, l.input[l.start:l.pos], l.line}
	l.start = l.pos
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.tokens <- token{
		ERROR,
		l.start,
		fmt.Sprintf(format, args...),
		l.line,
	}
	return nil
}

func (p Pos) Position() Pos {
	return p
}

// Get the next rune
func (l *Lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = Pos(0)
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
	if r == '\n' {
		l.line++
	}
	return r
}

// Wherever you are - ignore it
func (l *Lexer) ignore() {
	l.start = l.pos
}

// Went 1 char too far? Use this
func (l *Lexer) backup() {
	l.pos -= l.width
}

// Peek for the next char
func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// Check if it indicates a new token
// lexIdentifier doesn't respect this
func (l *Lexer) isDelimiter() bool {
	c := l.peek()

	if c == ' ' || c == '\n' || c == eof || c == '\t' || c == '{' || c == '}' {

		return true
	}

	return false
}

var components = map[string]TokenType{
	"cloud": COMPONENT,
	"node":  COMPONENT,
	"actor": COMPONENT,
}

var keywords = map[string]TokenType{
	"as": ALIAS,
	"{":  SCOPEIN,
	"}":  SCOPEOUT,
}

// Check for components
func (l *Lexer) isComponent() bool {
	if _, ok := components[l.input[l.start:l.pos]]; ok {
		return true
	}

	return false
}

// Check for keywords
func (l *Lexer) isKeyword() bool {
	if _, ok := keywords[l.input[l.start:l.pos]]; ok {
		return true
	}

	return false
}

// we already know its a keyword
func lexKeyword(l *Lexer) stateFn {
	if int(keywords[l.input[l.start:l.pos]]) == ALIAS {
		l.emit(keywords[l.input[l.start:l.pos]])
		return lexAlias
	}

	l.emit(keywords[l.input[l.start:l.pos]])
	return lexText
}

// Push a component off the heap
func lexComponent(l *Lexer) stateFn {
	// After component, you always need an identifier
	l.emit(COMPONENT)
	return lexIdentifier
}

// Alias is not allowed to be quoted
// Error check and continues with lexText
func lexAlias(l *Lexer) stateFn {
	l.stripWhitespaces()
	firstChar := l.peek()

	if firstChar == '"' || firstChar == '\'' {
		l.errorf("Aliases are not allowed to be quoted")
	}

	return lexText
}

// Lexes identifiers, matches foo as well as "foo bar" and 'bar
// baz' (including new line)
func lexIdentifier(l *Lexer) stateFn {
	l.stripWhitespaces()
	firstChar := l.peek()
	skipEscape := false
	if firstChar == '"' || firstChar == '\'' {
		skipEscape = true
		// Ignore first char
		l.next()
		l.ignore()
	}
	for {
		c := l.next()
		if skipEscape && c == firstChar {
			l.backup()
			l.emit(IDENTIFIER)
			l.next()
			l.ignore()
			break
		}

		if l.isComponent() {
			// Emit error if component follows
			l.errorf("A component can't be next to a component - need an identifier first")
			break
		}

		if !skipEscape && l.isDelimiter() {
			l.emit(IDENTIFIER)
			break
		}
	}

	return lexText
}

// Loops until next char isn't a whitespace and ignores them
func (l *Lexer) stripWhitespaces() {
	for {
		c := l.next()
		if c != ' ' && c != '\t' && c != '\n' {
			l.backup()
			break
		}
	}
	l.ignore()
}

// lex text - entrypoint if you don't know what type of text follows
func lexText(l *Lexer) stateFn {
	l.stripWhitespaces()
	for {
		if l.isDelimiter() {
			// Component found, identifier next
			if l.isComponent() {
				return lexComponent
			}
			if l.isKeyword() {
				return lexKeyword
			}
			if l.pos > l.start {
				l.emit(TEXT)
				return lexText
			}
		}
		if c := l.next(); c == eof {
			break
		}
	}

	// Correctly reached EOF
	// EOF is 0 for yacc
	l.emit(0)
	return nil
}

func (l *Lexer) Lex(lval *TdocSymType) int {

	for {
		select {
		case token := <-l.tokens:
			lval.line = token.line
			lval.pos = int(token.pos)
			lval.token = int(token.typ)
			lval.val = token.val
			return int(token.typ)
		default:
			l.state = l.state(l)
		}
	}
}
