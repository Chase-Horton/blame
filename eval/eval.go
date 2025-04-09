package eval

import (
	"strconv"

	"github.com/chase-horton/blame/ast"
)

var emitter = &CEmitter{}

type CEmitter struct {
	cfile string
}

func (e *CEmitter) Emit(s string) {
	e.cfile += s
}

// create evaluator that emits C code
func EvalProgram(p *ast.Program) string {
	// Emit C code for the program
	emitter.Emit("#include <stdio.h>\n")
	emitter.Emit("#include <stdlib.h>\n")
	emitter.Emit("#include <string.h>\n")
	//func main
	emitter.Emit("int main() {\n")
	for _, stmt := range p.Statements {
		EvalStatement(stmt)
	}
	emitter.Emit("return 0;\n")
	emitter.Emit("}\n")
	return emitter.cfile
}
func EvalStatement(stmt ast.Statement) {
	switch s := stmt.(type) {
	case *ast.AssignmentStatement:
		emitter.Emit(s.Identifier.Value + " = ")
		EvalExpression(s.Expression)
	default:
		panic("unknown statement type")
	}
}
func EvalBlock(b *ast.Block) {
	emitter.Emit("{\n")
	for _, stmt := range b.Statements {
		EvalStatement(stmt)
	}
	emitter.Emit("}\n")
}
func EvalExpression(e *ast.Expression) {
	emitter.Emit("/* Expression */\n")
	emitter.Emit("/* Left term */\n")
	EvalSignedTerm(e.Left)
	for _, term := range e.Right {
		emitter.Emit("/* Right term */\n")
		EvalSignedTerm(term)
	}
}
func EvalSignedTerm(st *ast.SignedTerm) {
	emitter.Emit("/* Signed term */\n")
	if st.Sign != "" {
		emitter.Emit(st.Sign + " ")
	}
	EvalTerm(st.Term)
	emitter.Emit("/* End signed term */\n")
}
func EvalTerm(t *ast.Term) {
	EvalFactor(t.Left)
	for i, op := range t.Op {
		emitter.Emit(op + " ")
		EvalFactor(t.Right[i])
	}
}
func EvalFactor(f ast.Factor) {
	switch f := f.(type) {
	case *ast.Identifier:
		emitter.Emit(f.Value + " ")
	case *ast.NumberLiteral:
		emitter.Emit(strconv.Itoa(f.Value) + " ")
	case *ast.StringLiteral:
		emitter.Emit(f.Value + " ")
	case *ast.ParenthesesExpression:
		emitter.Emit("(")
		EvalExpression(f.Expression)
		emitter.Emit(")")
	default:
		panic("unknown factor type")
	}
}
func EvalCondition(c *ast.Condition) {
	emitter.Emit("/* Condition */\n")
	EvalExpression(c.Left)
	emitter.Emit(c.Op + " ")
	EvalExpression(c.Right)
	emitter.Emit("/* End condition */\n")
}
