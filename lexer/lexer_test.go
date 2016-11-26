package lexer

import "testing"

func TestSimpleTextNextToken(t *testing.T) {
	input := "bar foo"

	tests := []struct {
		extectedType    tokenType
		expectedLiteral string
	}{
		{tokenText, "bar"},
		{tokenText, "foo"},
	}

	l := NewLexer("foo", input)
	for i, tt := range tests {
		tok := l.nextToken()
		if tok.typ != tt.extectedType {
			t.Fatalf("test[%d] - wrong type, expected=%q, got=%q", i, tt.extectedType, tok.typ)
		}
	}
}

func TestEmptyInput(t *testing.T) {
	input := ``

	l := NewLexer("foo", input)
	tok := l.nextToken()

	if tok.typ != tokenEOF {
		t.Fatalf("Empty input should return tokenEOF expected=%q, got=%q", tokenEOF, tok.typ)
	}

}

func TestComplexTextNextToken(t *testing.T) {
	input := `foo bar blubb
  baz
    quoo
    la
          le`

	tests := []struct {
		extectedType    tokenType
		expectedLiteral string
	}{
		{tokenText, "foo"},
		{tokenText, "bar"},
		{tokenText, "blubb"},
		{tokenText, "baz"},
		{tokenText, "quoo"},
		{tokenText, "la"},
		{tokenText, "le"},
	}

	l := NewLexer("foo", input)
	for i, tt := range tests {
		tok := l.nextToken()
		if tok.typ != tt.extectedType {
			t.Fatalf("test[%d] - wrong type, expected=%q, got=%q", i, tt.extectedType, tok.typ)
		}
	}
}

func TestSimpleComponentNextToken(t *testing.T) {
	input := `cloud actor node`

	tests := []struct {
		extectedType    tokenType
		expectedLiteral string
	}{
		{tokenCloud, "cloud"},
		{tokenActor, "actor"},
		{tokenNode, "node"},
	}

	l := NewLexer("foo", input)
	for i, tt := range tests {
		tok := l.nextToken()
		if tok.typ != tt.extectedType {
			t.Fatalf("test[%d] - wrong type, expected=%q, got=%q", i, tt.extectedType, tok.typ)
		}
	}
}

func TestSimpleMixNextToken(t *testing.T) {
	input := `cloud foo actor bar node duck`

	tests := []struct {
		extectedType    tokenType
		expectedLiteral string
	}{
		{tokenCloud, "cloud"},
		{tokenText, "foo"},
		{tokenActor, "actor"},
		{tokenText, "bar"},
		{tokenNode, "node"},
		{tokenText, "duck"},
	}

	l := NewLexer("foo", input)
	for i, tt := range tests {
		tok := l.nextToken()
		if tok.typ != tt.extectedType {
			t.Fatalf("test[%d] - wrong type, expected=%q, got=%q", i, tt.extectedType, tok.typ)
		}
		if tok.val != tt.expectedLiteral {
			t.Fatalf("test[%d] - wrong value, expected=%q, got=%q", i, tt.expectedLiteral, tok.val)
		}
	}
}

func TestUnicodeMixNextToken(t *testing.T) {
	input := `cloud ✓ actor ✓ node`

	tests := []struct {
		extectedType    tokenType
		expectedLiteral string
	}{
		{tokenCloud, "cloud"},
		{tokenText, "✓"},
		{tokenActor, "actor"},
		{tokenText, "✓"},
		{tokenNode, "node"},
	}

	l := NewLexer("foo", input)
	for i, tt := range tests {
		tok := l.nextToken()
		if tok.typ != tt.extectedType {
			t.Fatalf("test[%d] - wrong type, expected=%q, got=%q", i, tt.extectedType, tok.typ)
		}
		if tok.val != tt.expectedLiteral {
			t.Fatalf("test[%d] - wrong value, expected=%q, got=%q", i, tt.expectedLiteral, tok.val)
		}
	}
}
