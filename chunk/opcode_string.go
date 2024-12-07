// Code generated by "stringer -type=OpCode"; DO NOT EDIT.

package chunk

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OP_RETURN-0]
	_ = x[OP_CONSTANT-1]
	_ = x[OP_CONSTANT_LONG-2]
	_ = x[OP_NEGATE-3]
	_ = x[OP_ADD-4]
	_ = x[OP_SUB-5]
	_ = x[OP_MULT-6]
	_ = x[OP_DIV-7]
	_ = x[OP_NIL-8]
	_ = x[OP_TRUE-9]
	_ = x[OP_FALSE-10]
	_ = x[OP_EQUAL-11]
	_ = x[OP_GREATER-12]
	_ = x[OP_LESS-13]
}

const _OpCode_name = "OP_RETURNOP_CONSTANTOP_CONSTANT_LONGOP_NEGATEOP_ADDOP_SUBOP_MULTOP_DIVOP_NILOP_TRUEOP_FALSEOP_EQUALOP_GREATEROP_LESS"

var _OpCode_index = [...]uint8{0, 9, 20, 36, 45, 51, 57, 64, 70, 76, 83, 91, 99, 109, 116}

func (i OpCode) String() string {
	if i >= OpCode(len(_OpCode_index)-1) {
		return "OpCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OpCode_name[_OpCode_index[i]:_OpCode_index[i+1]]
}
