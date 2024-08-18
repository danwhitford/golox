package debug

import (
	"fmt"
	"testing"

	"github.com/danwhitford/golox/chunk"
	"github.com/danwhitford/golox/value"
	"github.com/google/go-cmp/cmp"
)

func TestDissembleBasicChunk(t *testing.T) {
	var ch chunk.Chunk
	ch.WriteCode(chunk.OP_RETURN, 1)
	want := []string{
		"0000 1 OP_RETURN",
	}
	got := DissembleChunk(ch)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}

func TestDissembleConstantChunk(t *testing.T) {
	var ch chunk.Chunk

	constant := ch.AddConstant(72.2)
	if constant != 0 {
		t.Fatalf("wanted '0', got '%v'", constant)
	}
	ch.WriteCode(chunk.OP_CONSTANT, 1)
	ch.WriteChunk(constant, 1)
	ch.WriteCode(chunk.OP_RETURN, 3)

	want := []string{
		"0000 1 OP_CONSTANT\t72.2",
		"0002 3 OP_RETURN",
	}

	got := DissembleChunk(ch)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}

func TestLineNumbers(t *testing.T) {
	var ch chunk.Chunk

	constant := ch.AddConstant(72.2)
	if constant != 0 {
		t.Fatalf("wanted '0', got '%v'", constant)
	}
	ch.WriteCode(chunk.OP_CONSTANT, 1)
	ch.WriteChunk(constant, 1)
	ch.WriteCode(chunk.OP_RETURN, 1)

	want := []string{
		"0000 1 OP_CONSTANT\t72.2",
		"0002 | OP_RETURN",
	}

	got := DissembleChunk(ch)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}

func TestLongConstants(t *testing.T) {
	var ch chunk.Chunk

	for i := 0; i < 300; i++ {
		ch.WriteConstant(value.Value(i) * 7, 1)
	}

	var want []string
	line := "1"
	offset := 0
	for i := 0; i < 300; i++ {
		var wantLine string
		if i < 256 {
			wantLine = fmt.Sprintf("%04d %s OP_CONSTANT\t%g", offset, line, float64(i*7))
			offset += 2
		} else {
			wantLine = fmt.Sprintf("%04d %s OP_CONSTANT_LONG\t%g", offset, line, float64(i*7))
			offset += 9
		}
		want = append(want, wantLine)
		line = "|"
	}

	got := DissembleChunk(ch)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}
