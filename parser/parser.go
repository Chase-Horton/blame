package parser

import (
	"github.com/chase-horton/blame/lexer"
)

type Parser struct {
	lexer        *lexer.Lexer
	errors       []string
	currentToken lexer.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer: l,
	}
	p.nextToken()
	return p
}
func (p *Parser) error(msg string) {
	p.errors = append(p.errors, msg)
}
func (p *Parser) nextToken() {
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
func (p *Parser) block() {
	p.expect(lexer.TokenOpenBrace)
	for p.currentToken.Type != lexer.TokenCloseBrace {
		p.statement()
	}
	p.expect(lexer.TokenCloseBrace)
}
func (p *Parser) statement() {
	if p.accept(lexer.TokenIdentifier) {
		if p.accept(lexer.TokenAssign) {
			p.expression()
		} else {
			p.error("expected assignment")
		}
	}
}
func (p *Parser) factor() {
	if p.accept(lexer.TokenIdentifier) {
	} else if p.accept(lexer.TokenNumber) {
	} else if p.accept(lexer.TokenOpenParen) {
		p.expression()
	} else {
		p.error("expected factor")
	}
}
func (p *Parser) term() {
	p.factor()
	for p.accept(lexer.TokenMultiply) || p.accept(lexer.TokenDivide) {
		p.factor()
	}
}
func (p *Parser) expression() {
	if p.accept(lexer.TokenPlus) || p.accept(lexer.TokenMinus) {
		p.nextToken()
	}
	p.term()
	for p.accept(lexer.TokenPlus) || p.accept(lexer.TokenMinus) {
		p.term()
	}
}
