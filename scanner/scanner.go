package scanner

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
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

var identifierMap map[string]TokenType = map[string]TokenType{
	"and":    TOKEN_AND,
	"class":  TOKEN_CLASS,
	"else":   TOKEN_ELSE,
	"if":     TOKEN_IF,
	"nil":    TOKEN_NIL,
	"or":     TOKEN_OR,
	"print":  TOKEN_PRINT,
	"return": TOKEN_RETURN,
	"super":  TOKEN_SUPER,
	"var":    TOKEN_VAR,
	"while":  TOKEN_WHILE,
	"false":  TOKEN_FALSE,
	"for":    TOKEN_FOR,
	"fun":    TOKEN_FUN,
	"this":   TOKEN_THIS,
	"true":   TOKEN_TRUE,
}

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
	err := scnr.skipWhitespace()
	if err != nil {
		return scnr.makeToken(
			TOKEN_ERROR,
			string(err.Error()),
		)
	}
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

	if unicode.IsLetter(r) {
		err := scnr.Source.UnreadRune()
		if err != nil {
			return Token{}
		}
		return scnr.identifier()
	}

	if unicode.IsDigit(r) {
		err := scnr.Source.UnreadRune()
		if err != nil {
			return Token{}
		}
		return scnr.number()
	}

	switch r {
	case '(':
		return scnr.makeToken(TOKEN_LEFT_PAREN, string(r))
	case ')':
		return scnr.makeToken(TOKEN_RIGHT_PAREN, string(r))
	case '{':
		return scnr.makeToken(TOKEN_LEFT_BRACE, string(r))
	case '}':
		return scnr.makeToken(TOKEN_RIGHT_BRACE, string(r))
	case ';':
		return scnr.makeToken(TOKEN_SEMICOLON, string(r))
	case ',':
		return scnr.makeToken(TOKEN_COMMA, string(r))
	case '.':
		return scnr.makeToken(TOKEN_DOT, string(r))
	case '-':
		return scnr.makeToken(TOKEN_MINUS, string(r))
	case '+':
		return scnr.makeToken(TOKEN_PLUS, string(r))
	case '/':
		return scnr.makeToken(TOKEN_SLASH, string(r))
	case '*':
		return scnr.makeToken(TOKEN_STAR, string(r))
	case '!':
		if b, err := scnr.match('='); b {
			return scnr.makeToken(TOKEN_BANG_EQUAL, "!=")
		} else {
			if err != nil {
				return scnr.makeToken(
					TOKEN_ERROR,
					string(err.Error()),
				)
			}
			return scnr.makeToken(TOKEN_BANG, "!")
		}
	case '=':
		if b, err := scnr.match('='); b {
			return scnr.makeToken(TOKEN_EQUAL_EQUAL, "==")
		} else {
			if err != nil {
				return scnr.makeToken(
					TOKEN_ERROR,
					string(err.Error()),
				)
			}
			return scnr.makeToken(TOKEN_EQUAL, "=")
		}
	case '<':
		if b, err := scnr.match('='); b {
			return scnr.makeToken(TOKEN_LESS_EQUAL, "<=")
		} else {
			if err != nil {
				return scnr.makeToken(
					TOKEN_ERROR,
					string(err.Error()),
				)
			}
			return scnr.makeToken(TOKEN_LESS, "<")
		}
	case '>':
		if b, err := scnr.match('='); b {
			return scnr.makeToken(TOKEN_GREATER_EQUAL, ">=")
		} else {
			if err != nil {
				return scnr.makeToken(
					TOKEN_ERROR,
					string(err.Error()),
				)
			}
			return scnr.makeToken(TOKEN_GREATER, ">")
		}
	case '"':
		return scnr.string()
	}

	return scnr.makeToken(
		TOKEN_ERROR,
		fmt.Sprintf("do not understand token '%c'", r),
	)
}

func (scnr *Scanner) identifier() Token {
	var sb strings.Builder
	for {
		r, _, err := scnr.Source.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return Token{}
			}
		}
		if !unicode.In(r, unicode.Number, unicode.Letter) {
			break
		}
		sb.WriteRune(r)
	}

	return scnr.makeToken(
		identifierType(sb.String()),
		sb.String(),
	)
}

func identifierType(word string) TokenType {
	m, prs := identifierMap[word]
	if prs {
		return m
	} else {
		return TOKEN_IDENTIFIER
	}
}

func (scnr *Scanner) number() Token {
	var sb strings.Builder
	for {
		r, _, err := scnr.Source.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return scnr.makeToken(
					TOKEN_ERROR,
					string(err.Error()),
				)
			}
		}
		if r != '.' && !unicode.IsDigit(r) {
			err := scnr.Source.UnreadRune()
			if err != nil {
				panic(err)
			}
			break
		}
		sb.WriteRune(r)
	}

	return scnr.makeToken(
		TOKEN_NUMBER,
		sb.String(),
	)
}

func (scnr *Scanner) string() Token {
	var sb strings.Builder
	for {
		r, _, err := scnr.Source.ReadRune()
		if err != nil {
			if err == io.EOF {
				return scnr.makeToken(
					TOKEN_ERROR,
					"unterminated string",
				)
			} else {
				return Token{}
			}
		}
		if r == '"' {
			break
		}
		sb.WriteRune(r)
	}

	return scnr.makeToken(
		TOKEN_STRING,
		sb.String(),
	)
}

func (scnr *Scanner) skipWhitespace() error {
	for {
		next, _, err := scnr.Source.ReadRune()
		if err != nil {
			if err == io.EOF {
				return nil
			} else {
				return err
			}
		}
		if !unicode.IsSpace(next) {
			err := scnr.Source.UnreadRune()
			if err != nil {
				return err
			}
			return nil
		}
	}
}

func (scnr *Scanner) match(r rune) (bool, error) {
	next, _, err := scnr.Source.ReadRune()
	if err != nil {
		if err == io.EOF {
			return false, nil
		} else {
			return false, err
		}
	}
	if r != next {
		err = scnr.Source.UnreadRune()
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (scnr *Scanner) advance() (rune, int, error) {
	return scnr.Source.ReadRune()
}

func (scnr *Scanner) makeToken(tt TokenType, lexeme string) Token {
	return Token{
		tt,
		lexeme,
		scnr.Line,
	}
}
