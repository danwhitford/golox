package compiler

import (
	"fmt"
	"strconv"

	"github.com/danwhitford/golox/chunk"
	"github.com/danwhitford/golox/scanner"
	"github.com/danwhitford/golox/value"
)

type Compiler struct {
	Scnr         *scanner.Scanner
	CurrentChunk chunk.Chunk
	token        scanner.Token
}

type parseRule struct {
	prefix func()
	infix  func()
}

func Init(source string) *Compiler {
	return &Compiler{
		Scnr: scanner.NewScanner(source),
	}
}

func (compiler *Compiler) getRule(t scanner.TokenType) parseRule {
	switch t {
	case scanner.TOKEN_MINUS:
		return parseRule{compiler.unary, compiler.binary}
	case scanner.TOKEN_NUMBER:
		return parseRule{compiler.number, nil}
	case scanner.TOKEN_PLUS, scanner.TOKEN_STAR, scanner.TOKEN_SLASH:		
		return parseRule{nil, compiler.binary}
	}
	panic("don't know rule for " + t.String())
}

func (compiler *Compiler) unary() {
	compiler.token = compiler.Scnr.ScanToken()
	compiler.expression()
	compiler.CurrentChunk.WriteCode(chunk.OP_NEGATE, compiler.token.Line)
}

func (compiler *Compiler) binary() {
	infixer := compiler.token
	compiler.token = compiler.Scnr.ScanToken()
	compiler.expression()
	switch infixer.Type {
	case scanner.TOKEN_PLUS:
		compiler.CurrentChunk.WriteCode(chunk.OP_ADD, compiler.token.Line)
	case scanner.TOKEN_MINUS:
		compiler.CurrentChunk.WriteCode(chunk.OP_SUB, compiler.token.Line)
	case scanner.TOKEN_STAR:
		compiler.CurrentChunk.WriteCode(chunk.OP_MULT, compiler.token.Line)
	case scanner.TOKEN_SLASH:
		compiler.CurrentChunk.WriteCode(chunk.OP_DIV, compiler.token.Line)
	default: panic("don't know infix for '" + infixer.Lexeme + "'")
	}
}

func (compiler *Compiler) number() {
	f, err := strconv.ParseFloat(compiler.token.Lexeme, 64)
	if err != nil {
		panic(fmt.Sprintf("float parse error. %v", err))
	}
	compiler.CurrentChunk.WriteConstant(value.Value(f), compiler.token.Line)
}

func (compiler *Compiler) expression() {
	rule := compiler.getRule(compiler.token.Type)
	rule.prefix()

	for {
		compiler.token = compiler.Scnr.ScanToken()
		if compiler.token.Type == scanner.TOKEN_EOF {
			return
		}
		rule := compiler.getRule(compiler.token.Type)
		rule.infix()
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
