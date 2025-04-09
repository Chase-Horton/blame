package parser

import (
	"testing"

	"github.com/chase-horton/blame/lexer"
)

func TestParser(t *testing.T) {
	// Test cases for the parser
	lexer := lexer.New("x = 5; y = 2; z = x + y;")
	parser := New(lexer)
	ast := parser.Parse()
	if ast == nil {
		t.Fatalf("Expected AST, got nil")
	}
	if len(ast.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(ast.Statements))
	}
}
