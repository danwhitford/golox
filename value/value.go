package value

type ValueType byte

const (
	VAL_BOOL ValueType = iota
	VAL_NIL
	VAL_NUMBER
)

type Value struct {
	T  ValueType
	As any
}

func BoolVal(b bool) Value {
	return Value{VAL_BOOL, b}
}

func NilVal() Value {
	return Value{VAL_NIL, nil}
}

func NumberVal(f float64) Value {
	return Value{VAL_NUMBER, f}
}

func (v Value) AsBool() bool {
	return v.As.(bool)
}

func (v Value) AsNumber() float64 {
	return v.As.(float64)
}
