package chunk

import (
	"github.com/danwhitford/golox/runlengthencoder"
	"github.com/danwhitford/golox/value"
)

type OpCode byte

const (
	OP_RETURN OpCode = iota
	OP_CONSTANT
	OP_CONSTANT_LONG
)

type Chunk struct {
	Code      []byte
	Constants []value.Value
	Lines     runlengthencoder.RunLengthEncoder
}

func (ch *Chunk) WriteCode(code OpCode, line int) {
	ch.WriteChunk(byte(code), line)
}

func (ch *Chunk) WriteChunk(b byte, line int) {
	ch.Code = append(ch.Code, b)
	ch.Lines.Append(line)
}

func (ch *Chunk) AddConstant(value value.Value) byte {
	ch.Constants = append(ch.Constants, value)
	return byte(len(ch.Constants) - 1)
}

func (ch *Chunk) WriteConstant(value value.Value, line int) {
	panic("cba tbh")
}
