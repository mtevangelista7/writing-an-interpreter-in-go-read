package token

// definimos como string pois dessa forma podemos salvar qualquer tipo de token
type TokenType string

type Token struct {
	Type    TokenType
	Literal string // armazenar o valor literal do token
}

const (
	ILLEGAL = "ILLEGAL" // token não reconhecido
	EOF     = "EOF"     // final do arquivo
	// Identificadores + literais
	IDENT = "IDENT" // variaveis, nomes de funções etc
	INT   = "INT"
	// Operadores
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="
	// Delimitadores
	COMMA      = ","
	SEMICOLON  = ";"
	LEFTPAREN  = "("
	RIGHTPAREN = ")"
	LEFTBRACE  = "{"
	RIGHTBRACE = "}"
	// Palavras reservadas
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent verifica o tipo de um identificador
func LookupIdent(ident string) TokenType {
	// verifica se existe no map de keywords
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
