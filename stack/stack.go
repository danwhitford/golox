package stack

import "fmt"

type Stack[T any] []T

func (st *Stack[T]) Push(t T) {
	*st = append(*st, t)
}

func (st *Stack[T]) Pop() (T, error) {
	l := len(*st)
	if l < 1 {
		var t T
		return t, fmt.Errorf("stack underflow")
	}
	top := (*st)[l-1]
	*st = (*st)[:l-1]
	return top, nil
}

func (st *Stack[T]) Empty() bool {
	return len(*st) < 1
}
