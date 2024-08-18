package debug

import (
	"fmt"

	"github.com/danwhitford/golox/chunk"
)

func DissembleChunk(ch chunk.Chunk) []string {
	var dissembled []string

	offset := 0
	for offset < len(ch.Code) {
		line, shift := dissembleInstruction(ch, offset)
		line = fmt.Sprintf("%04d %s %s", offset, offsetString(ch, offset), line)
		dissembled = append(dissembled, line)
		offset += shift
	}
	return dissembled
}

func dissembleInstruction(ch chunk.Chunk, offset int) (string, int) {
	code := ch.Code[offset]
	switch (chunk.OpCode(code)) {
		case chunk.OP_RETURN: return "OP_RETURN", 1
		case chunk.OP_CONSTANT: return constantInstruction(ch, offset), 2
	}
	panic(fmt.Sprintf("instruction not recognised: '%v'", ch))
}

func constantInstruction(ch chunk.Chunk, offset int) string {
	chunkLoc := ch.Code[offset+1]
	cst := ch.Constants[chunkLoc]
	return fmt.Sprintf("OP_CONSTANT\t%g", cst)
}

func offsetString(ch chunk.Chunk, offset int) string {
	if offset > 0 && ch.Lines[offset] == ch.Lines[offset-1] {
		return "|"
	} else {
		return fmt.Sprintf("%d", ch.Lines[offset])
	}
}
