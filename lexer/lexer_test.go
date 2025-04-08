package lexer

import "testing"

func testPrintUtil(t *testing.T, L *Lexer) {
	t.Logf("Lexing:\n%v", string(L.input))
	for token := L.NextToken(); token.Type != TokenEOF; {
		t.Logf("Token: %v", token)
		token = L.NextToken()
	}
}
func TestLexer(t *testing.T) {
	L := New("x = 5")
	testPrintUtil(t, L)
	L, err := NewFromFile("./test_files/test1.blame")
	if err != nil {
		t.Fatalf("Error opening file: %v", err)
	}
	testPrintUtil(t, L)
}
