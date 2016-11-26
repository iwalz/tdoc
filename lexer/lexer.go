package lexer

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const eof = -1

type Pos int
type stateFn func(*Lexer) stateFn
type tokenType int

type Lexer struct {
	name   string // used for error reports
	input  string // input to scan
	start  Pos    // start position of this item
	pos    Pos    // current position of the input
	width  Pos    // width of last rune read
	line   int
	tokens chan token // channel of scanned items
	state  stateFn    // The state function
}

type token struct {
	typ  tokenType // The type
	pos  Pos       // and position
	val  string    // The value
	line int       // Line
}

const (
	tokenError tokenType = iota
	tokenText
	tokenCloud
	tokenNode
	tokenActor
	tokenEOF
)

func (t token) String() string {
	switch t.typ {
	case tokenEOF:
		return "EOF"
	case tokenError:
		return t.val
	}
	if len(t.val) > 10 {
		return fmt.Sprintf("%.10q...", t.val)
	}

	return fmt.Sprintf("%q", t.val)
}

func NewLexer(name, input string) *Lexer {
	l := &Lexer{
		name:   name,
		input:  input,
		state:  lexText,
		tokens: make(chan token, 2),
	}

	return l
}

// Sends a new item to the channel
func (l *Lexer) emit(t tokenType) {
	l.tokens <- token{t, l.start, l.input[l.start:l.pos], l.line}
	l.start = l.pos
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.tokens <- token{
		tokenError,
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

// Check if next rune is acceptable
func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *Lexer) isDelimiter() bool {
	c := l.peek()

	if c == ' ' || c == '\n' || c == eof || c == '\t' {

		return true
	}

	return false
}

var components = map[string]tokenType{
	"cloud": tokenCloud,
	"node":  tokenNode,
	"actor": tokenActor,
}

func (l *Lexer) isComponent() bool {
	if _, ok := components[l.input[l.start:l.pos]]; ok {
		return true
	}

	return false
}

func lexComponent(l *Lexer) stateFn {
	token, _ := components[l.input[l.start:l.pos]]

	l.emit(token)
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

func lexText(l *Lexer) stateFn {
	l.stripWhitespaces()
	for {
		if l.isDelimiter() {
			// Component found, name next
			if l.isComponent() {
				return lexComponent
			}
			if l.pos > l.start {
				l.emit(tokenText)
				return lexText
			}
		}
		if c := l.next(); c == eof {
			break
		}
	}

	// Correctly reached EOF
	l.emit(tokenEOF)
	return nil
}

func (l *Lexer) nextToken() token {
	for {
		select {
		case token := <-l.tokens:
			return token
		default:
			l.state = l.state(l)
		}
	}
}

func (l *Lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}
