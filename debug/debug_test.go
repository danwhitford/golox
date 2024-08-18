package debug

import (
	"testing"

	"github.com/danwhitford/golox/chunk"
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
