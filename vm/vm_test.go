package vm

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/danwhitford/golox/chunk"
	"github.com/danwhitford/golox/value"
	"github.com/google/go-cmp/cmp"
)

func TestReturnOk(t *testing.T) {
	var vm Vm
	got := vm.Interpret("")

	want := INTERPRET_OK

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}

func TestConstantOk(t *testing.T) {
	var buff bytes.Buffer
	var vm Vm
	vm.Out = &buff
	stat := vm.Interpret("72.0")

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
	var buff bytes.Buffer
	var vm Vm
	vm.DebugMode = true
	vm.Out = &buff
	stat := vm.Interpret("72.0")

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
	var buff bytes.Buffer
	var vm Vm
	vm.Out = &buff
	stat := vm.Interpret("-24.5")

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
		var buff bytes.Buffer
		var vm Vm
		vm.Out = &buff
		stat := vm.Interpret(fmt.Sprintf("%v %v %v", tst.v1, tst.op, tst.v2))

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
