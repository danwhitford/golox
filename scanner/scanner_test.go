package scanner

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestScanToken(t *testing.T) {
	table := []struct {
		code   string
		tokens []Token
	}{
		{
			"",
			[]Token{
				{
					Type:   TOKEN_EOF,
					Lexeme: "",
					Line:   1,
				},
			},
		},
		{
			"(",
			[]Token{
				{
					Type:   TOKEN_LEFT_PAREN,
					Lexeme: "(",
					Line:   1,
				},
			},
		},
		{
			")",
			[]Token{
				{
					Type:   TOKEN_RIGHT_PAREN,
					Lexeme: ")",
					Line:   1,
				},
			},
		},
		{
			"=",
			[]Token{
				{
					Type:   TOKEN_EQUAL,
					Lexeme: "=",
					Line:   1,
				},
			},
		},
		{
			"==",
			[]Token{
				{
					Type:   TOKEN_EQUAL_EQUAL,
					Lexeme: "==",
					Line:   1,
				},
			},
		},
		{
			`"foo bar"`,
			[]Token{
				{
					Type:   TOKEN_STRING,
					Lexeme: "foo bar",
					Line:   1,
				},
			},
		},
		{
			`14`,
			[]Token{
				{
					Type:   TOKEN_NUMBER,
					Lexeme: "14",
					Line:   1,
				},
			},
		},
		{
			`14.5078`,
			[]Token{
				{
					Type:   TOKEN_NUMBER,
					Lexeme: "14.5078",
					Line:   1,
				},
			},
		},
		{
			`foo`,
			[]Token{
				{
					Type:   TOKEN_IDENTIFIER,
					Lexeme: "foo",
					Line:   1,
				},
			},
		},
		{
			`fun`,
			[]Token{
				{
					Type:   TOKEN_FUN,
					Lexeme: "fun",
					Line:   1,
				},
			},
		},
		{
			`= `,
			[]Token{
				{
					Type:   TOKEN_EQUAL,
					Lexeme: "=",
					Line:   1,
				},
			},
		},
		{
			`<<`,
			[]Token{
				{
					Type:   TOKEN_LESS,
					Lexeme: "<",
					Line:   1,
				},
				{
					Type:   TOKEN_LESS,
					Lexeme: "<",
					Line:   1,
				},
			},
		},
		{
			`1+2`,
			[]Token{
				{
					Type:   TOKEN_NUMBER,
					Lexeme: "1",
					Line:   1,
				},
				{
					Type:   TOKEN_PLUS,
					Lexeme: "+",
					Line:   1,
				},
				{
					Type:   TOKEN_NUMBER,
					Lexeme: "2",
					Line:   1,
				},
			},
		},
		{
			`(1+2)`,
			[]Token{
				{
					Type:   TOKEN_LEFT_PAREN,
					Lexeme: "(",
					Line:   1,
				},
				{
					Type:   TOKEN_NUMBER,
					Lexeme: "1",
					Line:   1,
				},
				{
					Type:   TOKEN_PLUS,
					Lexeme: "+",
					Line:   1,
				},
				{
					Type:   TOKEN_NUMBER,
					Lexeme: "2",
					Line:   1,
				},
				{
					Type:   TOKEN_RIGHT_PAREN,
					Lexeme: ")",
					Line:   1,
				},
			},
		},
		{
			`(1 + 2) - 3`,
			[]Token{
				{
					Type:   TOKEN_LEFT_PAREN,
					Lexeme: "(",
					Line:   1,
				},
				{
					Type:   TOKEN_NUMBER,
					Lexeme: "1",
					Line:   1,
				},
				{
					Type:   TOKEN_PLUS,
					Lexeme: "+",
					Line:   1,
				},
				{
					Type:   TOKEN_NUMBER,
					Lexeme: "2",
					Line:   1,
				},
				{
					Type:   TOKEN_RIGHT_PAREN,
					Lexeme: ")",
					Line:   1,
				},
				{
					Type:   TOKEN_MINUS,
					Lexeme: "-",
					Line:   1,
				},
				{
					Type:   TOKEN_NUMBER,
					Lexeme: "3",
					Line:   1,
				},
			},
		},
	}

	for i, tst := range table {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			scnr := NewScanner(tst.code)
			for _, tkn := range tst.tokens {
				want := tkn
				got := scnr.ScanToken()
				if got.Type == TOKEN_ERROR {
					t.Fatalf("%d: %v", i, got.Lexeme)
				}
				if diff := cmp.Diff(want, got); diff != "" {
					t.Fatalf("%d: Mismatch (-want +got):\n%s", i, diff)
				}
			}
		})
	}
}
