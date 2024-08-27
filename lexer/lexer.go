package lexer

import (
	"writing-an-interpreter-in-go-read/token"
)

type Lexer struct {
	input        string // input da linguagem
	position     int    // posição atual (char atual) que já foi lida
	readPosition int    // posição seguinte ao char atual a ser lido
	ch           byte   // char atual
}

// New, inicia um lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar, atualiza as variaveis de posição do char
func (l *Lexer) readChar() {
	// verifica se a proxima posição é o final do input
	// caso sim, podemos falar que o char é 0 (NULL na tabela ASCII)
	// caso não, ch recebe a proxima posição
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken, recebe um lexer e devolve o proximo token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LEFTPAREN, l.ch)
	case ')':
		tok = newToken(token.RIGHTPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LEFTBRACE, l.ch)
	case '}':
		tok = newToken(token.RIGHTBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// newToken, retorna um novo token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// isLetter verifica se o ch é uma letra
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// readIdentifier verifica se um lexer é um identificador
// verifica do começo até um ch que não seja uma letra
// corta esse conteudo e devolve como identificador
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace ignora qualquer tipo de espaço em branco e segue para o prox char
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// isDigit verifica se o char é um número
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// readNumber analisa o lexer e devolve o número
func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}
