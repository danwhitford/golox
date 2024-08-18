package chunk

import "github.com/danwhitford/golox/value"

type OpCode byte

const (
	OP_RETURN OpCode = iota
	OP_CONSTANT
)

type Chunk struct {
	Code []byte
	Constants []value.Value
	Lines []int
}

func (ch *Chunk) WriteCode(code OpCode, line int) {
	ch.WriteChunk(byte(code), line)
}

func (ch *Chunk) WriteChunk(b byte, line int) {
	ch.Code = append(ch.Code, b)
	ch.Lines = append(ch.Lines, line)
}

func (ch *Chunk) AddConstant(value value.Value) byte {
	ch.Constants = append(ch.Constants, value)
	return byte(len(ch.Constants) -1)
}
