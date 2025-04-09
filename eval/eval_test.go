package eval

import (
	"fmt"
	"testing"

	"github.com/chase-horton/blame/parser"

	"github.com/chase-horton/blame/lexer"
)

func TestEvalProgram(t *testing.T) {
	l := lexer.New("x = 5; y = 2; z = x + y;")
	p := parser.New(l)
	ast := p.Parse()
	if ast == nil {
		t.Fatalf("Expected AST, got nil")
	}
	fmt.Println(EvalProgram(ast))

}
