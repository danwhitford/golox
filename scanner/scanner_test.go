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
	}

	for _, tst := range table {
		scnr := NewScanner(tst.code)
		for _, tkn := range tst.tokens {
			want := tkn
			got := scnr.ScanToken()
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("Mismatch (-want +got):\n%s", diff)
			}
		}
	}
}
