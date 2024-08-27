package repl

import (
	"bufio"
	"fmt"
	"io"
	"writing-an-interpreter-in-go-read/lexer"
	"writing-an-interpreter-in-go-read/token"
)

const PROMPT = ">> "

// Start inicia o modo interativo
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
