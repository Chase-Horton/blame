package parser

import (
	"strconv"

	"github.com/chase-horton/blame/ast"

	"github.com/chase-horton/blame/lexer"
)

type Parser struct {
	lexer            *lexer.Lexer
	errors           []string
	currentToken     lexer.Token
	prevToken        lexer.Token
	identifierLookup map[string]map[string]*ast.Identifier
	currentScope     string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:            l,
		currentScope:     "global",
		identifierLookup: make(map[string]map[string]*ast.Identifier),
	}
	p.nextToken()
	return p
}
func (p *Parser) lookupIdent(ident string) *ast.Identifier {
	if _, ok := p.identifierLookup[p.currentScope]; !ok {
		p.identifierLookup[p.currentScope] = make(map[string]*ast.Identifier)
	}
	if _, ok := p.identifierLookup[p.currentScope][ident]; !ok {
		p.identifierLookup[p.currentScope][ident] = &ast.Identifier{Value: ident}
	}
	return p.identifierLookup[p.currentScope][ident]
}
func (p *Parser) error(msg string) {
	p.errors = append(p.errors, msg)
}
func (p *Parser) nextToken() {
	p.prevToken = p.currentToken
	p.currentToken = p.lexer.NextToken()
}
func (p *Parser) accept(t lexer.TokenType) bool {
	if p.currentToken.Type == t {
		p.nextToken()
		return true
	}
	return false
}
func (p *Parser) expect(t lexer.TokenType) bool {
	if p.currentToken.Type != t {
		p.error("expected token: " + t.String() + ", got: " + p.currentToken.Type.String())
		return false
	}
	p.nextToken()
	return true
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	for p.currentToken.Type != lexer.TokenEOF {
		program.Statements = append(program.Statements, p.statement())
	}
	return program
}
func (p *Parser) block() *ast.Block {
	block := ast.Block{}
	p.expect(lexer.TokenOpenBrace)
	for p.currentToken.Type != lexer.TokenCloseBrace {
		block.Statements = append(block.Statements, p.statement())
	}
	p.expect(lexer.TokenCloseBrace)
	return &block
}
func (p *Parser) statement() ast.Statement {
	if p.accept(lexer.TokenIdentifier) {
		S := &ast.AssignmentStatement{Identifier: p.lookupIdent(p.prevToken.Literal)}
		if p.accept(lexer.TokenAssign) {
			S.Value = p.expression()
		} else {
			p.error("expected assignment")
		}
		return S
	} else if p.accept(lexer.TokenKeywordIf) {
		iff := ast.IfStatement{}
		iff.Condition = p.condition()
		iff.Then = p.block()
		if p.accept(lexer.TokenKeywordElse) {
			iff.Else = p.block()
		}
		return &iff
	} else if p.accept(lexer.TokenKeywordWhile) {
		C := p.condition()
		B := p.block()
		return &ast.WhileStatement{Condition: C, Block: B}
	} else if p.accept(lexer.TokenKeywordDo) {
		B := p.block()
		if p.accept(lexer.TokenKeywordWhile) {
			C := p.condition()
			return &ast.DoWhileStatement{Block: B, Condition: C}
		} else {
			p.error("expected while after do")
		}
	} else if p.accept(lexer.TokenOpenBrace) {
		B := p.block()
		return &ast.BlockStatement{Block: B}
	}
	return nil
}
func (p *Parser) condition() *ast.Condition {
	C := ast.Condition{}
	C.Left = p.expression()
	//expect operand
	if p.accept(lexer.TokenEqual) || p.accept(lexer.TokenNotEqual) ||
		p.accept(lexer.TokenLessThan) || p.accept(lexer.TokenGreaterThan) ||
		p.accept(lexer.TokenLessThanEqual) || p.accept(lexer.TokenGreaterThanEqual) {
		C.Op = p.prevToken.Literal
		C.Right = p.expression()
	} else {
		p.error("expected condition operator")
	}
	return &C
}
func (p *Parser) factor() ast.Factor {
	if p.accept(lexer.TokenIdentifier) {
		return &ast.Identifier{Value: p.prevToken.Literal}
	} else if p.accept(lexer.TokenNumber) {
		num, err := strconv.Atoi(p.prevToken.Literal)
		if err != nil {
			p.error("invalid number literal")
			return nil
		}
		return &ast.NumberLiteral{Value: num}
	} else if p.accept(lexer.TokenOpenParen) {
		return p.parenExpression()
	} else {
		p.error("expected factor")
	}
	return nil
}
func (p *Parser) term() *ast.Term {
	T := ast.Term{}
	T.Left = p.factor()
	for p.accept(lexer.TokenMultiply) || p.accept(lexer.TokenDivide) {
		T.Op = append(T.Op, p.prevToken.Literal)
		T.Right = append(T.Right, p.factor())
	}
	return &T
}
func (p *Parser) signedTerm() *ast.SignedTerm {
	if p.accept(lexer.TokenPlus) || p.accept(lexer.TokenMinus) {
		return &ast.SignedTerm{Sign: p.prevToken.Literal, Term: p.term()}
	}
	return nil
}
func (p *Parser) expression() *ast.Expression {
	E := ast.Expression{}
	if p.accept(lexer.TokenPlus) || p.accept(lexer.TokenMinus) {
		E.Left = &ast.SignedTerm{Sign: p.prevToken.Literal}
	}
	E.Left.Term = p.term()
	for p.accept(lexer.TokenPlus) || p.accept(lexer.TokenMinus) {
		E.Op = append(E.Op, p.prevToken.Literal)
		E.Right = append(E.Right, p.signedTerm())
	}
	return &E
}
func (p *Parser) parenExpression() *ast.ParenthesesExpression {
	p.expect(lexer.TokenOpenParen)
	expr := p.expression()
	p.expect(lexer.TokenCloseParen)
	return &ast.ParenthesesExpression{Expression: expr}
}
