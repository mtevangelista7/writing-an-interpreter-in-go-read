package ast

import "writing-an-interpreter-in-go-read/token"

type Node interface {
	TokenLiteral() string // debug only
}

type Statement interface {
	// "herda" a interface node, ou seja, para implementar essa interface
	// é preciso também, implementar Node
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// representa o nó raiz da ast
// a linguagem monkey é basicamente uma série de statements
// e para criar esse é o nó inial
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// let
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// é o x, em let x = 10
type Identifier struct {
	Token token.Token // token type IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// return
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
