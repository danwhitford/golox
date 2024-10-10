package compiler

import (
	"fmt"
	"strconv"

	"github.com/danwhitford/golox/chunk"
	"github.com/danwhitford/golox/scanner"
	"github.com/danwhitford/golox/value"
)

type Compiler struct {
	Scnr *scanner.Scanner
	CurrentChunk chunk.Chunk
	token scanner.Token
}

func Init(source string) *Compiler {
	return &Compiler{
		Scnr: scanner.NewScanner(source),
	}
}

func (compiler *Compiler) Compile(source string) chunk.Chunk {

	for {
		compiler.token = compiler.Scnr.ScanToken()
		
		if compiler.token.Type == scanner.TOKEN_EOF {
			compiler.CurrentChunk.WriteCode(chunk.OP_RETURN, compiler.token.Line)
			break
		}

		compiler.expression()
	}

	return compiler.CurrentChunk
}

func (compiler *Compiler) expression() {
	switch compiler.token.Type {
	case scanner.TOKEN_NUMBER:
		{
			f, err := strconv.ParseFloat(compiler.token.Lexeme, 64)
			if err != nil {
				panic(fmt.Sprintf("float parse error. %v", err))
			}
			compiler.CurrentChunk.WriteConstant(value.Value(f), compiler.token.Line)
		}
	}
}
