package chunk

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/danwhitford/golox/runlengthencoder"
	"github.com/danwhitford/golox/value"
)

//go:generate stringer -type=OpCode
type OpCode byte

const (
	OP_RETURN OpCode = iota
	OP_CONSTANT
	OP_CONSTANT_LONG
	OP_NEGATE
	OP_ADD
	OP_SUB
	OP_MULT
	OP_DIV
	OP_NIL
	OP_TRUE
	OP_FALSE
	OP_EQUAL
	OP_GREATER
	OP_LESS
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
	ch.Constants = append(ch.Constants, value)
	var constantIndex uint64 = uint64(len(ch.Constants) - 1)
	if constantIndex <= 255 {
		ch.WriteCode(OP_CONSTANT, line)
		ch.WriteChunk(byte(constantIndex), line)
	} else {
		ch.WriteCode(OP_CONSTANT_LONG, line)
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, constantIndex)
		if err != nil {
			panic(fmt.Sprintf("failed to write constant, '%v'", value))
		}
		for _, b := range buf.Bytes() {
			ch.WriteChunk(b, line)
		}
	}
}
