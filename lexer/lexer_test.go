package lexer

import "testing"

func testPrintUtil(t *testing.T, L *Lexer) {
	t.Logf("Lexing:\n%v", string(L.input))
	L.position = 0
	L.nextPosition = 0
	L.readChar()
	for token := L.NextToken(); true; {
		t.Logf("Token: %v", token)
		if token.Type == TokenEOF {
			break
		}
		token = L.NextToken()
	}
}
func TestLexer(t *testing.T) {
	testCases := []struct {
		input    string
		expected []Token
	}{
		{"x = 5;", []Token{
			{Type: TokenIdentifier, Literal: "x"},
			{Type: TokenAssign, Literal: "="},
			{Type: TokenNumber, Literal: "5"},
			{Type: TokenSep, Literal: ";"},
			{Type: TokenEOF, Literal: ""},
		}},
		{"x - y + z;", []Token{
			{Type: TokenIdentifier, Literal: "x"},
			{Type: TokenMinus, Literal: "-"},
			{Type: TokenIdentifier, Literal: "y"},
			{Type: TokenPlus, Literal: "+"},
			{Type: TokenIdentifier, Literal: "z"},
			{Type: TokenSep, Literal: ";"},
			{Type: TokenEOF, Literal: ""},
		}},
		{"if (x < y) { x = 5; }", []Token{
			{Type: TokenKeywordIf, Literal: "if"},
			{Type: TokenOpenParen, Literal: "("},
			{Type: TokenIdentifier, Literal: "x"},
			{Type: TokenLessThan, Literal: "<"},
			{Type: TokenIdentifier, Literal: "y"},
			{Type: TokenCloseParen, Literal: ")"},
			{Type: TokenOpenBrace, Literal: "{"},
			{Type: TokenIdentifier, Literal: "x"},
			{Type: TokenAssign, Literal: "="},
			{Type: TokenNumber, Literal: "5"},
			{Type: TokenSep, Literal: ";"},
			{Type: TokenCloseBrace, Literal: "}"},
			{Type: TokenEOF, Literal: ""},
		}},
	}
	for _, testCase := range testCases {
		L := New(testCase.input)
		for i, expected := range testCase.expected {
			token := L.NextToken()
			if token.Type != expected.Type {
				t.Fatalf("Test case %d: Expected token type %v, got %v", i, expected.Type, token.Type)
			}
			if token.Literal != expected.Literal {
				t.Fatalf("Test case %d: Expected token literal %q, got %q", i, expected.Literal, token.Literal)
			}
		}
		testPrintUtil(t, L)
	}
}
