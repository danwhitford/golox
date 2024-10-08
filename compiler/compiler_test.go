package compiler

import (
	"testing"

	"github.com/danwhitford/golox/chunk"
	"github.com/danwhitford/golox/value"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCompile(t *testing.T) {
	table := []struct {
		source string
		want   chunk.Chunk
	}{
		{
			"",
			chunk.Chunk{
				Code:      []byte{byte(chunk.OP_RETURN)},
				Constants: nil,
			},
		},
		{
			"4",
			chunk.Chunk{
				Code: []byte{
					byte(chunk.OP_CONSTANT),
					byte(0),
					byte(chunk.OP_RETURN),
				},
				Constants: []value.Value{4},
			},
		},
	}

	opts := cmpopts.IgnoreFields(chunk.Chunk{}, "Lines")
	for i, tst := range table {
		cmplr := Init(tst.source)
		got := cmplr.Compile(tst.source)
		if diff := cmp.Diff(tst.want, got, opts); diff != "" {
			t.Errorf("%d: Mismatch (-want +got):\n%s", i, diff)
		}
	}
}
