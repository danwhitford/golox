package chunk

import "testing"

func TestChunk(t *testing.T) {
	var chunk Chunk

	constant := chunk.AddConstant(72.2)
	chunk.WriteCode(OP_CONSTANT, 1)
	chunk.WriteChunk(constant, 1)
	chunk.WriteCode(OP_RETURN, 3)
}
