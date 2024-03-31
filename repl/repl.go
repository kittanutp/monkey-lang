package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey-lang/lexer"
	"monkey-lang/token"
)

const PROMT = ">> "

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)
	for {
		fmt.Fprint(output, PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(output, "%+v\n", tok)
		}
	}
}
