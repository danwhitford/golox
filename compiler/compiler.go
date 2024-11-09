package compiler

import (
	"fmt"
	"strconv"

	"github.com/danwhitford/golox/chunk"
	"github.com/danwhitford/golox/scanner"
	"github.com/danwhitford/golox/value"
)

type Compiler struct {
	Scnr          *scanner.Scanner
	CurrentChunk  chunk.Chunk
	currentToken  scanner.Token
	previousToken scanner.Token
}

type parseRule struct {
	prefix     func()
	infix      func()
	precedence precedence
}

type precedence int

const (
	PREC_NONE       precedence = iota
	PREC_ASSIGNMENT            // =
	PREC_OR                    // or
	PREC_AND                   // and
	PREC_EQUALITY              // == !=
	PREC_COMPARISON            // < > <= >=
	PREC_TERM                  // + -
	PREC_FACTOR                // * /
	PREC_UNARY                 // ! -
	PREC_CALL                  // . ()
	PREC_PRIMARY
)

func Init(source string) *Compiler {
	return &Compiler{
		Scnr: scanner.NewScanner(source),
	}
}

func (compiler *Compiler) advance() {
	compiler.previousToken = compiler.currentToken
	compiler.currentToken = compiler.Scnr.ScanToken()
}

func (compiler *Compiler) getRule(t scanner.TokenType) parseRule {
	switch t {
	case scanner.TOKEN_MINUS:
		return parseRule{compiler.unary, compiler.binary, PREC_TERM}
	case scanner.TOKEN_NUMBER:
		return parseRule{compiler.number, nil, PREC_NONE}
	case scanner.TOKEN_PLUS:
		return parseRule{nil, compiler.binary, PREC_TERM}
	case scanner.TOKEN_STAR, scanner.TOKEN_SLASH:
		return parseRule{nil, compiler.binary, PREC_FACTOR}
	case scanner.TOKEN_LEFT_PAREN:
		return parseRule{compiler.grouping, nil, PREC_NONE}
	case scanner.TOKEN_EOF, scanner.TOKEN_RIGHT_PAREN:
		return parseRule{nil, nil, PREC_NONE}
	case scanner.TOKEN_FALSE, scanner.TOKEN_TRUE, scanner.TOKEN_NIL:
		return parseRule{compiler.literal, nil, PREC_NONE}
	case scanner.TOKEN_BANG:
		return parseRule{compiler.unary, nil, PREC_NONE}
	case scanner.TOKEN_EQUAL_EQUAL, scanner.TOKEN_BANG_EQUAL:
		return parseRule{nil, compiler.binary, PREC_EQUALITY}
	case scanner.TOKEN_GREATER, scanner.TOKEN_GREATER_EQUAL, scanner.TOKEN_LESS, scanner.TOKEN_LESS_EQUAL:
		return parseRule{nil, compiler.binary, PREC_COMPARISON}
	}
	fmt.Printf("%#v\n", compiler)
	panic("don't know rule for " + t.String())
}

func (compiler *Compiler) unary() {
	compiler.parseWithPrecedence(PREC_UNARY)
	compiler.CurrentChunk.WriteCode(chunk.OP_NEGATE, compiler.currentToken.Line)
}

func (compiler *Compiler) binary() {
	infixer := compiler.previousToken.Type
	rule := compiler.getRule(infixer)
	compiler.parseWithPrecedence(rule.precedence + 1)

	switch infixer {
	case scanner.TOKEN_PLUS:
		compiler.writeOpCode(chunk.OP_ADD)
	case scanner.TOKEN_MINUS:
		compiler.writeOpCode(chunk.OP_SUB)
	case scanner.TOKEN_STAR:
		compiler.writeOpCode(chunk.OP_MULT)
	case scanner.TOKEN_SLASH:
		compiler.writeOpCode(chunk.OP_DIV)
	case scanner.TOKEN_BANG_EQUAL:
		compiler.writeOpCodes(chunk.OP_EQUAL, chunk.OP_NEGATE)
	case scanner.TOKEN_LESS:
		compiler.writeOpCode(chunk.OP_LESS)
	case scanner.TOKEN_GREATER:
		compiler.writeOpCode(chunk.OP_GREATER)
	case scanner.TOKEN_EQUAL_EQUAL:
		compiler.writeOpCode(chunk.OP_EQUAL)
	case scanner.TOKEN_GREATER_EQUAL:
		compiler.writeOpCodes(chunk.OP_LESS, chunk.OP_NEGATE)
	case scanner.TOKEN_LESS_EQUAL:
		compiler.writeOpCodes(chunk.OP_GREATER, chunk.OP_NEGATE)
	default:
		panic("don't know infix for '" + infixer.String() + "'")
	}
}

func (compiler *Compiler) number() {
	f, err := strconv.ParseFloat(compiler.previousToken.Lexeme, 64)
	if err != nil {
		panic(fmt.Sprintf("float parse error. %v", err))
	}
	compiler.CurrentChunk.WriteConstant(value.NumberVal(f), compiler.currentToken.Line)
}

func (compiler *Compiler) grouping() {
	compiler.expression()
	compiler.consume(scanner.TOKEN_RIGHT_PAREN)
}

func (compiler *Compiler) literal() {
	switch compiler.previousToken.Type {
	case scanner.TOKEN_FALSE:
		compiler.writeOpCode(chunk.OP_FALSE)
	case scanner.TOKEN_TRUE:
		compiler.writeOpCode(chunk.OP_TRUE)
	case scanner.TOKEN_NIL:
		compiler.writeOpCode(chunk.OP_NIL)
	default:
		panic(fmt.Sprintf("'%v' is not a literal type", compiler.previousToken.Type))
	}
}

func (compiler *Compiler) writeOpCode(op chunk.OpCode) {
	compiler.CurrentChunk.WriteCode(op, compiler.currentToken.Line)
}

func (compiler *Compiler) writeOpCodes(a, b chunk.OpCode) {
	compiler.writeOpCode(a)
	compiler.writeOpCode(b)
}

func (compiler *Compiler) consume(t scanner.TokenType) {
	if compiler.currentToken.Type == t {
		compiler.advance()
		return
	}
	panic("wanted '" + t.String() + "' but got '" + compiler.currentToken.Type.String() + "'")
}

func (compiler *Compiler) expression() {
	compiler.parseWithPrecedence(PREC_ASSIGNMENT)
}

func (compiler *Compiler) parseWithPrecedence(prec precedence) {
	compiler.advance()
	rule := compiler.getRule(compiler.previousToken.Type)
	if rule.prefix == nil {
		panic(fmt.Sprintf("no prefix rule for '%v'", compiler.previousToken.Type))
	}
	rule.prefix()

	for prec <= compiler.getRule(compiler.currentToken.Type).precedence {
		compiler.advance()
		rule := compiler.getRule(compiler.previousToken.Type)
		rule.infix()
	}
}

func (compiler *Compiler) Compile(source string) chunk.Chunk {
	if len(source) < 1 {
		compiler.CurrentChunk.WriteCode(chunk.OP_RETURN, compiler.currentToken.Line)
		return compiler.CurrentChunk
	}

	compiler.advance()    // prime the pump
	compiler.expression() // read the expression

	if compiler.currentToken.Type == scanner.TOKEN_EOF {
		compiler.CurrentChunk.WriteCode(chunk.OP_RETURN, compiler.currentToken.Line)
		return compiler.CurrentChunk
	} else {
		panic("expected 'EOF' but got '" + compiler.currentToken.Type.String() + "'")
	}
}
