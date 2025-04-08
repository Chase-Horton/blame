package lexer

import "os"

type TokenType string

func (t *TokenType) String() string {
	return string(*t)
}

const (
	// Special tokens
	TokenEOF        TokenType = "EOF"
	TokenIllegal    TokenType = "ILLEGAL"
	TokenWhitespace TokenType = "WHITESPACE"
	// Identifiers
	TokenIdentifier TokenType = "IDENTIFIER"
	// Keywords
	TokenKeywordIf    TokenType = "IF"
	TokenKeywordElse  TokenType = "ELSE"
	TokenKeywordWhile TokenType = "WHILE"
	TokenKeywordDo    TokenType = "DO"
	// Operators
	TokenPlus     TokenType = "+"
	TokenMinus    TokenType = "-"
	TokenMultiply TokenType = "*"
	TokenDivide   TokenType = "/"
	//Comparators
	TokenAssign           TokenType = "EQUALS"
	TokenEqual            TokenType = "CHECK_EQUALS"
	TokenNotEqual         TokenType = "NOT_EQUAL"
	TokenLessThan         TokenType = "<"
	TokenGreaterThan      TokenType = "<"
	TokenLessThanEqual    TokenType = "<="
	TokenGreaterThanEqual TokenType = ">="
	TokenAnd              TokenType = "AND"
	TokenOr               TokenType = "OR"
	TokenNot              TokenType = "NOT"
	//Separators
	TokenComma      TokenType = ","
	TokenOpenParen  TokenType = "("
	TokenCloseParen TokenType = ")"
	TokenOpenBrace  TokenType = "{"
	TokenCloseBrace TokenType = "}"
	TokenSep        TokenType = "SEPARATOR"
	// Literals
	TokenString TokenType = "STRING"
	TokenNumber TokenType = "NUMBER"
)

type Token struct {
	Type    TokenType
	Literal string
}

func newToken(t TokenType, literal string) Token {
	return Token{
		Type:    t,
		Literal: literal,
	}
}

type Lexer struct {
	input        []rune
	position     int
	nextPosition int
	currentChar  rune
}

func New(in string) *Lexer {
	l := &Lexer{
		input: []rune(in),
	}
	l.readChar()
	return l
}
func NewFromFile(filePath string) (*Lexer, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	l := &Lexer{
		input: []rune(string(data)),
	}
	l.readChar()
	return l, nil
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = rune(l.input[l.nextPosition])
	}
	l.position = l.nextPosition
	l.nextPosition++
}
func (l *Lexer) peek() rune {
	if l.nextPosition >= len(l.input) {
		return 0
	}
	return rune(l.input[l.nextPosition])
}

func isEnglishLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}
func (l *Lexer) identifyToken() Token {
	start := l.position
	for isEnglishLetter(l.currentChar) {
		l.readChar()
	}
	identifier := string(l.input[start:l.position])
	switch identifier {
	case "if":
		return newToken(TokenKeywordIf, identifier)
	case "else":
		return newToken(TokenKeywordElse, identifier)
	case "while":
		return newToken(TokenKeywordWhile, identifier)
	case "do":
		return newToken(TokenKeywordDo, identifier)
	default:
		return newToken(TokenIdentifier, identifier)
	}
}
func (l *Lexer) identifyNumber() Token {
	start := l.position
	for isDigit(l.currentChar) {
		l.readChar()
	}
	return newToken(TokenNumber, string(l.input[start:l.position]))

}

func (l *Lexer) NextToken() Token {
	for isWhitespace(l.currentChar) {
		l.readChar()
	}
	switch l.currentChar {
	case 0:
		return newToken(TokenEOF, "")
	case '+':
		l.readChar()
		return newToken(TokenPlus, "+")
	case '-':
		l.readChar()
		return newToken(TokenMinus, "-")
	case '*':
		l.readChar()
		return newToken(TokenMultiply, "*")
	case '/':
		l.readChar()
		return newToken(TokenDivide, "/")
	case '=':
		if l.peek() == '=' {
			l.readChar()
			l.readChar()
			return newToken(TokenEqual, "==")
		}
		l.readChar()
		return newToken(TokenAssign, "=")
	case '\n', ';':
		l.readChar()
		return newToken(TokenSep, "\\n")
	case ',':
		l.readChar()
		return newToken(TokenComma, ",")
	case '(':
		l.readChar()
		return newToken(TokenOpenParen, "(")
	case ')':
		l.readChar()
		return newToken(TokenCloseParen, ")")
	case '{':
		l.readChar()
		return newToken(TokenOpenBrace, "{")
	case '}':
		l.readChar()
		return newToken(TokenCloseBrace, "}")
	default:
		if isEnglishLetter(l.currentChar) {
			return l.identifyToken()
		} else if isDigit(l.currentChar) {
			return l.identifyNumber()
		} else {
			return newToken(TokenIllegal, string(l.currentChar))
		}
	}
}
