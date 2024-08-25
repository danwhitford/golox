package vm

import (
	"bytes"
	"testing"

	"github.com/danwhitford/golox/chunk"
	"github.com/danwhitford/golox/value"
	"github.com/google/go-cmp/cmp"
)

func TestReturnOk(t *testing.T) {
	var ch chunk.Chunk
	ch.WriteCode(chunk.OP_RETURN, 1)

	var vm Vm
	got := vm.Interpret(ch)

	want := INTERPRET_OK

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}

func TestConstantOk(t *testing.T) {
	var ch chunk.Chunk
	ch.WriteConstant(72.0, 1)
	ch.WriteCode(chunk.OP_RETURN, 2)

	var buff bytes.Buffer
	var vm Vm
	vm.Out = &buff
	stat := vm.Interpret(ch)

	if stat != INTERPRET_OK {
		t.Errorf("not ok '%v'", stat)
	}

	want := "72\n"
	got := buff.String()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}

func TestDebugOk(t *testing.T) {
	var ch chunk.Chunk
	ch.WriteConstant(72.0, 1)
	ch.WriteCode(chunk.OP_RETURN, 2)

	var buff bytes.Buffer
	var vm Vm
	vm.DebugMode = true
	vm.Out = &buff
	stat := vm.Interpret(ch)

	if stat != INTERPRET_OK {
		t.Errorf("not ok '%v'", stat)
	}

	want := ` === 
 === 
OP_CONSTANT	72
 === 
 [ 72 ] 
 === 
OP_RETURN
72
`
	got := buff.String()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}

func TestNegateOk(t *testing.T) {
	var ch chunk.Chunk
	ch.WriteConstant(24.5, 1)
	ch.WriteCode(chunk.OP_NEGATE, 1)
	ch.WriteCode(chunk.OP_RETURN, 2)

	var buff bytes.Buffer
	var vm Vm
	vm.Out = &buff
	stat := vm.Interpret(ch)

	if stat != INTERPRET_OK {
		t.Errorf("not ok '%v'", stat)
	}

	want := "-24.5\n"

	got := buff.String()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}

func TestBinaryOpsOl(t *testing.T) {
	table := []struct {
		v1   value.Value
		v2   value.Value
		op   chunk.OpCode
		want string
	}{
		{
			2, 7, chunk.OP_ADD, "9",
		},
		{
			2, 7, chunk.OP_SUB, "-5",
		},
		{
			2, 7, chunk.OP_MULT, "14",
		},
		{
			14, 7, chunk.OP_DIV, "2",
		},
	}

	for _, tst := range table {
		var ch chunk.Chunk
		ch.WriteConstant(tst.v1, 1)
		ch.WriteConstant(tst.v2, 1)
		ch.WriteCode(tst.op, 1)
		ch.WriteCode(chunk.OP_RETURN, 2)

		var buff bytes.Buffer
		var vm Vm
		vm.Out = &buff
		stat := vm.Interpret(ch)

		if stat != INTERPRET_OK {
			t.Errorf("not ok '%v'", stat)
		}

		want := tst.want + "\n"
		got := buff.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("Mismatch (-want +got):\n%s", diff)
		}
	}
}
