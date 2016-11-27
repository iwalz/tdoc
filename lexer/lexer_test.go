package lexer

import "testing"

func TestSimpleTextNextToken(t *testing.T) {
	input := "bar foo"

	tests := []struct {
		extectedType    int
		expectedLiteral string
	}{
		{TEXT, "bar"},
		{TEXT, "foo"},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		if tok != tt.extectedType {
			t.Fatalf("test[%d] - wrong type, expected=%q, got=%q", i, tt.extectedType, tok)
		}
	}
}

func TestEmptyInput(t *testing.T) {
	input := ``

	l := NewLexer(input)
	lval := &TdocSymType{}
	tok := l.Lex(lval)

	if tok != 0 {
		t.Fatalf("Empty input should return 0 expected=%q, got=%q", 0, tok)
	}

}

func TestComplexTextNextToken(t *testing.T) {
	input := `foo bar blubb
  baz
    quoo
    la
          le`

	tests := []struct {
		extectedType    int
		expectedLiteral string
	}{
		{TEXT, "foo"},
		{TEXT, "bar"},
		{TEXT, "blubb"},
		{TEXT, "baz"},
		{TEXT, "quoo"},
		{TEXT, "la"},
		{TEXT, "le"},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		if tok != tt.extectedType {
			t.Fatalf("test[%d] - wrong type, expected=%q, got=%q", i, tt.extectedType, tok)
		}
	}
}

func TestSimpleComponentNextToken(t *testing.T) {
	input := `cloud actor node`

	tests := []struct {
		extectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "cloud"},
		{ERROR, "actor"},
		{COMPONENT, "node"},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		if tok != tt.extectedType {
			t.Fatalf("test[%d] - wrong type, expected=%q, got=%q", i, tt.extectedType, tok)
		}
	}
}

func TestSimpleMixNextToken(t *testing.T) {
	input := `cloud foo actor bar node duck`

	tests := []struct {
		extectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "cloud"},
		{IDENTIFIER, "foo"},
		{COMPONENT, "actor"},
		{IDENTIFIER, "bar"},
		{COMPONENT, "node"},
		{IDENTIFIER, "duck"},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		if tok != tt.extectedType {
			t.Fatalf("test[%d] - wrong type, expected=%q, got=%q", i, tt.extectedType, tok)
		}
		if lval.val != tt.expectedLiteral {
			t.Fatalf("test[%d] - wrong value, expected=%q, got=%q", i, tt.expectedLiteral, lval.val)
		}
	}
}

func TestSingleQuoteIdentifier(t *testing.T) {
	input := `actor 'test foo'`

	tests := []struct {
		extectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "actor"},
		{IDENTIFIER, "test foo"},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		if tok != tt.extectedType {
			t.Fatalf("test[%d] - wrong type, expected=%q, got=%q", i, tt.extectedType, tok)
		}
		if lval.val != tt.expectedLiteral {
			t.Fatalf("test[%d] - wrong value, expected=%q, got=%q", i, tt.expectedLiteral, lval.val)
		}
	}
}

func TestDoubleQuoteIdentifier(t *testing.T) {
	input := `actor "test
	foo"`

	tests := []struct {
		extectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "actor"},
		{IDENTIFIER, "test foo"},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		if tok != tt.extectedType {
			t.Fatalf("test[%d] - wrong type, expected=%q, got=%q", i, tt.extectedType, tok)
		}
		if lval.val != tt.expectedLiteral {
			t.Fatalf("test[%d] - wrong value, expected=%q, got=%q", i, tt.expectedLiteral, lval.val)
		}
	}
}

func TestUnicodeMixNextToken(t *testing.T) {
	input := `cloud ✓ actor ✓ node`

	tests := []struct {
		extectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "cloud"},
		{IDENTIFIER, "✓"},
		{COMPONENT, "actor"},
		{IDENTIFIER, "✓"},
		{COMPONENT, "node"},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		if tok != tt.extectedType {
			t.Fatalf("test[%d] - wrong type, expected=%q, got=%q", i, tt.extectedType, tok)
		}
		if lval.val != tt.expectedLiteral {
			t.Fatalf("test[%d] - wrong value, expected=%q, got=%q", i, tt.expectedLiteral, lval.val)
		}
	}
}
