package ast

type Node interface {
}
type Statement interface {
	StmtType() string
}
type Program struct {
	Statements []Statement
}
type Block struct {
	Statements []Statement
}
type Expression struct {
	Left  *SignedTerm
	Right []*SignedTerm
	Op    []string
}
type ParenthesesExpression struct {
	Expression *Expression
}
type Condition struct {
	Left  *Expression
	Right *Expression
	Op    string
}
type SignedTerm struct {
	Sign string
	Term *Term
}
type Term struct {
	Left  Factor
	Right []Factor
	Op    []string
}
type Factor interface {
	Factor()
}
type Identifier struct {
	Value string
}
type NumberLiteral struct {
	Value int
}
type StringLiteral struct {
	Value string
}

func (*Identifier) Factor() {
	panic("implement me")
}
func (*NumberLiteral) Factor() {
	panic("implement me")
}
func (*StringLiteral) Factor() {
	panic("implement me")
}
func (*ParenthesesExpression) Factor() {
	panic("implement me")
}

// stmts
type IfStatement struct {
	Condition *Condition
	Then      *Block
	Else      *Block
}
type WhileStatement struct {
	Condition *Condition
	Block     *Block
}
type DoWhileStatement struct {
	Condition *Condition
	Block     *Block
}
type BlockStatement struct {
	Block *Block
}
type AssignmentStatement struct {
	Identifier *Identifier
	Expression *Expression
}

func (*AssignmentStatement) StmtType() string {
	return "AssignmentStatement"
}
func (*IfStatement) StmtType() string {
	return "IfStatement"
}
func (*WhileStatement) StmtType() string {
	return "WhileStatement"
}
func (*DoWhileStatement) StmtType() string {
	return "DoWhileStatement"
}
func (*BlockStatement) StmtType() string {
	return "BlockStatement"
}
