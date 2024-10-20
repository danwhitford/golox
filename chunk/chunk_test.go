package chunk

import (
	"testing"

	"github.com/danwhitford/golox/value"
)

func TestChunk(t *testing.T) {
	var chunk Chunk

	constant := chunk.AddConstant(value.NumberVal(72.2))
	chunk.WriteCode(OP_CONSTANT, 1)
	chunk.WriteChunk(constant, 1)
	chunk.WriteCode(OP_RETURN, 3)
}

func TestWriteLongConstant(t *testing.T) {
	var chunk Chunk

	for i := 0; i < 300; i++ {
		chunk.WriteConstant(value.NumberVal(float64(i)), i)
	}
}
