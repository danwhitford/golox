package compiler

import (
	"fmt"

	"github.com/danwhitford/golox/scanner"
)

func Compile(source string) {
	scner := scanner.NewScanner(source)

	line := -1
	for {
		token := scner.ScanToken()
		if token.Line != line {
			fmt.Printf("%4d", token.Line)
			line = token.Line
		} else {
			fmt.Print("    | ")
		}
		fmt.Printf(" %v '%s'\n", token.Type, token.Lexeme)

		if token.Type == scanner.TOKEN_EOF {
			break
		}
	}
}
