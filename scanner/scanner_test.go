package scanner

import (
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
	}

	for _, tst := range table {
		scnr := NewScanner(tst.code)
		for i, tkn := range tst.tokens {
			want := tkn
			got := scnr.ScanToken()
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("%d: Mismatch (-want +got):\n%s", i, diff)
			}
		}
	}
}
