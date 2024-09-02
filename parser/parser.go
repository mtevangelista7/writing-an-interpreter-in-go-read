package parser

// Statements (Declarações)
// Um statement é uma instrução que realiza uma ação.
// Ele pode alterar o estado do programa ou produzir um efeito colateral.
// Statements não necessariamente retornam um valor.
// x = 5 /// if (x > 0) /// while (x > 0) /// def foo()

// Expressions (Expressões)
// Uma expressão é qualquer parte de um código que pode ser avaliada para produzir um valor.
// Expressões podem ser usadas dentro de statements.
// Sempre retornam um valor
// 2 + 3 /// foo(2) /// obj.property /// x > 0

import (
	"fmt"
	"writing-an-interpreter-in-go-read/ast"
	"writing-an-interpreter-in-go-read/lexer"
	"writing-an-interpreter-in-go-read/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	// Faz a leitura de dois tokens, pois dessa forma os dois iniciaram com valor
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

// peekError adiciona um erro personalizado para os casos onde o token é diferente do esperado
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	// passa o proximo para a leitura do token atual
	p.curToken = p.peekToken

	// seta o proximo token
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// caso não ainda não tenha chego no final
	for p.curToken.Type != token.EOF {

		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

// parseStatement retorna um statement completo baseado no tipo do token lido
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// parseLetStatement retorna uma expressão para o caso de declaração de varável (LET)
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// primeiro verificamos se o primeiro item é um identificador
	// ou seja um nome de variavel
	// let teste
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// pega o nome
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// depois verificamos se o proximo sinal é de igual
	// let teste =
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: aqui falta fazer o parse do resto da expressão, ou seja
	// o que fica entre o sinal de igual e o ponto e virgula
	// aqui seguimos até encontrar um ponto e virgula que finaliza a expressão
	// let teste = .... ;
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// curTokenIs verifica se um curTokenIs é do tipo esperado
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// curTokenIs verifica se um peekTokenIs é do tipo esperado
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek verifica se o proximo token é de acordo com o esperado
// caso sim, avança mais um token
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
