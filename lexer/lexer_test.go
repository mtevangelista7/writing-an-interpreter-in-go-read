package lexer

import (
	"testing"
	"writing-an-interpreter-in-go-read/token"
)

func TestNextToken(t *testing.T) {
	// apenas para testar isso poderia ser
	// um read file, lendo de um arquivo de verdade a linguagem
	// porém no momento não é preciso
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
	x + y;
	};
	let result = add(five, ten);
	`

	// slice de struct com os tipos e seu valor literal
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{

		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LEFTPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RIGHTPAREN, ")"},
		{token.LEFTBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LEFTPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RIGHTPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

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
