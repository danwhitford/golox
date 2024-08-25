package debug

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/danwhitford/golox/chunk"
)

func DissembleChunk(ch chunk.Chunk) []string {
	var dissembled []string

	offset := 0
	for offset < len(ch.Code) {
		line, shift := DissembleInstruction(ch, offset)
		line = fmt.Sprintf("%04d %s %s", offset, offsetString(ch, offset), line)
		dissembled = append(dissembled, line)
		offset += shift
	}
	return dissembled
}

func DissembleInstruction(ch chunk.Chunk, offset int) (string, int) {
	code := ch.Code[offset]
	switch chunk.OpCode(code) {
	case chunk.OP_RETURN:
		return "OP_RETURN", 1
	case chunk.OP_CONSTANT:
		return constantInstruction(ch, offset), 2
	case chunk.OP_CONSTANT_LONG:
		return constantLongInstruction(ch, offset), 9
	case chunk.OP_NEGATE:
		return "OP_NEGATE", 1
	case chunk.OP_ADD:
		return "OP_ADD", 1
	}
	panic(fmt.Sprintf("instruction not recognised: '%v'", ch.Code[offset]))
}

func constantLongInstruction(ch chunk.Chunk, offset int) string {
	var constIdx uint64
	bb := ch.Code[offset+1 : offset+9]
	buf := bytes.NewReader(bb)
	err := binary.Read(buf, binary.BigEndian, &constIdx)
	if err != nil {
		panic("failed to convert bytes to float64")
	}
	val := ch.Constants[constIdx]
	return fmt.Sprintf("OP_CONSTANT_LONG\t%g", val)
}

func constantInstruction(ch chunk.Chunk, offset int) string {
	chunkLoc := ch.Code[offset+1]
	cst := ch.Constants[chunkLoc]
	return fmt.Sprintf("OP_CONSTANT\t%g", cst)
}

func offsetString(ch chunk.Chunk, offset int) string {
	if offset > 0 && ch.Lines.Get(offset) == ch.Lines.Get(offset-1) {
		return "|"
	} else {
		return fmt.Sprintf("%d", ch.Lines.Get(offset))
	}
}
