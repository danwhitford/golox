package scanner

import (
	"bufio"
	"io"
	"strings"
)

type TokenType byte

//go:generate stringer -type=TokenType
const (
	TOKEN_EOF TokenType = iota
	TOKEN_LEFT_PAREN
	TOKEN_RIGHT_PAREN
	TOKEN_LEFT_BRACE
	TOKEN_RIGHT_BRACE
	TOKEN_COMMA
	TOKEN_DOT
	TOKEN_MINUS
	TOKEN_PLUS
	TOKEN_SEMICOLON
	TOKEN_SLASH
	TOKEN_STAR
	TOKEN_BANG
	TOKEN_BANG_EQUAL
	TOKEN_EQUAL
	TOKEN_EQUAL_EQUAL
	TOKEN_GREATER
	TOKEN_GREATER_EQUAL
	TOKEN_LESS
	TOKEN_LESS_EQUAL
	TOKEN_IDENTIFIER
	TOKEN_STRING
	TOKEN_NUMBER
	TOKEN_AND
	TOKEN_CLASS
	TOKEN_ELSE
	TOKEN_FALSE
	TOKEN_FOR
	TOKEN_FUN
	TOKEN_IF
	TOKEN_NIL
	TOKEN_OR
	TOKEN_PRINT
	TOKEN_RETURN
	TOKEN_SUPER
	TOKEN_THIS
	TOKEN_TRUE
	TOKEN_VAR
	TOKEN_WHILE
	TOKEN_ERROR
)

type Scanner struct {
	Source               *bufio.Reader
	Start, Current, Line int
}

func NewScanner(source string) *Scanner {
	var scner Scanner
	scner.Source = bufio.NewReader(strings.NewReader(source))
	scner.Line = 1
	scner.Start = 0
	scner.Current = 0
	return &scner
}

type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
}

func (scnr *Scanner) ScanToken() Token {
	r, _, err := scnr.advance()
	if err != nil {
		if err == io.EOF {
			return Token{
				TOKEN_EOF,
				"",
				scnr.Line,
			}
		}
		return Token{
			TOKEN_ERROR,
			string(r),
			scnr.Line,
		}
	}

	switch r {
		case '(': return scnr.makeToken(TOKEN_LEFT_PAREN, r);
		case ')': return scnr.makeToken(TOKEN_RIGHT_PAREN, r);
		case '{': return scnr.makeToken(TOKEN_LEFT_BRACE, r);
		case '}': return scnr.makeToken(TOKEN_RIGHT_BRACE, r);
		case ';': return scnr.makeToken(TOKEN_SEMICOLON, r);
		case ',': return scnr.makeToken(TOKEN_COMMA, r);
		case '.': return scnr.makeToken(TOKEN_DOT, r);
		case '-': return scnr.makeToken(TOKEN_MINUS, r);
		case '+': return scnr.makeToken(TOKEN_PLUS, r);
		case '/': return scnr.makeToken(TOKEN_SLASH, r);
		case '*': return scnr.makeToken(TOKEN_STAR, r);
	}

	panic("end of tokens dunno what to do")
}

func (scnr *Scanner) advance() (rune, int, error) {
	return scnr.Source.ReadRune()
}

func (scnr *Scanner) makeToken(tt TokenType, r rune) Token {
	return Token{
		tt,
		string(r),
		scnr.Line,
	}
}