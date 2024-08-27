package lexer

import (
	"testing"
	"writing-an-interpreter-in-go-read/token"
)

func TestNextToken(t *testing.T) {
	// apenas para testar isso poderia ser
	// um read file, lendo de um arquivo de verdade a linguagem
	// porém no momento não é preciso
	input := `
	`

	// slice de struct com os tipos e seu valor literal
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{}

	// inicia um novo lexer com a input da sintaxe monkey
	l := New(input)

	for index, tt := range tests {
		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				index, tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				index, tt.expectedLiteral, token.Literal)
		}
	}
}
