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
				Constants: []value.Value{value.NumberVal(4)},
			},
		},
		{
			"-42",
			chunk.Chunk{
				Code: []byte{
					byte(chunk.OP_CONSTANT),
					byte(0),
					byte(chunk.OP_NEGATE),
					byte(chunk.OP_RETURN),
				},
				Constants: []value.Value{value.NumberVal(42)},
			},
		},
		{
			"1 + 2",
			chunk.Chunk{
				Code: []byte{
					byte(chunk.OP_CONSTANT),
					byte(0),
					byte(chunk.OP_CONSTANT),
					byte(1),
					byte(chunk.OP_ADD),
					byte(chunk.OP_RETURN),
				},
				Constants: []value.Value{
					value.NumberVal(1),
					value.NumberVal(2),
				},
			},
		},
		{
			"3 - 4",
			chunk.Chunk{
				Code: []byte{
					byte(chunk.OP_CONSTANT),
					byte(0),
					byte(chunk.OP_CONSTANT),
					byte(1),
					byte(chunk.OP_SUB),
					byte(chunk.OP_RETURN),
				},
				Constants: []value.Value{
					value.NumberVal(3),
					value.NumberVal(4),
				},
			},
		},
		{
			"10 * 5",
			chunk.Chunk{
				Code: []byte{
					byte(chunk.OP_CONSTANT),
					byte(0),
					byte(chunk.OP_CONSTANT),
					byte(1),
					byte(chunk.OP_MULT),
					byte(chunk.OP_RETURN),
				},
				Constants: []value.Value{
					value.NumberVal(10),
					value.NumberVal(5),
				},
			},
		},
		{
			"100 / 3",
			chunk.Chunk{
				Code: []byte{
					byte(chunk.OP_CONSTANT),
					byte(0),
					byte(chunk.OP_CONSTANT),
					byte(1),
					byte(chunk.OP_DIV),
					byte(chunk.OP_RETURN),
				},
				Constants: []value.Value{
					value.NumberVal(100),
					value.NumberVal(3),
				},
			},
		},
		{
			"(1 + 2) - 3",
			chunk.Chunk{
				Code: []byte{
					byte(chunk.OP_CONSTANT),
					byte(0),
					byte(chunk.OP_CONSTANT),
					byte(1),
					byte(chunk.OP_ADD),
					byte(chunk.OP_CONSTANT),
					byte(2),
					byte(chunk.OP_SUB),
					byte(chunk.OP_RETURN),
				},
				Constants: []value.Value{
					value.NumberVal(1),
					value.NumberVal(2),
					value.NumberVal(3),
				},
			},
		},
		{
			"1 + (2 - 3)",
			chunk.Chunk{
				Code: []byte{
					byte(chunk.OP_CONSTANT),
					byte(0),
					byte(chunk.OP_CONSTANT),
					byte(1),
					byte(chunk.OP_CONSTANT),
					byte(2),
					byte(chunk.OP_SUB),
					byte(chunk.OP_ADD),
					byte(chunk.OP_RETURN),
				},
				Constants: []value.Value{
					value.NumberVal(1),
					value.NumberVal(2),
					value.NumberVal(3),
				},
			},
		},
		{
			"-2 + 5", // 2 NEGATE 5 +
			chunk.Chunk{
				Code: []byte{
					byte(chunk.OP_CONSTANT),
					byte(0),
					byte(chunk.OP_NEGATE), // 0x03
					byte(chunk.OP_CONSTANT),
					byte(1),
					byte(chunk.OP_ADD),
					byte(chunk.OP_RETURN),
				},
				Constants: []value.Value{
					value.NumberVal(2),
					value.NumberVal(5),
				},
			},
		},
	}

	opts := cmpopts.IgnoreFields(chunk.Chunk{}, "Lines")
	for i, tst := range table {
		t.Logf("running %d =[ %s ]=", i, tst.source)
		cmplr := Init(tst.source)
		got := cmplr.Compile(tst.source)
		if diff := cmp.Diff(tst.want, got, opts); diff != "" {
			t.Fatalf("%d: Mismatch (-want +got):\n%s", i, diff)
		}
		t.Logf("passed %d", i)
	}
}
