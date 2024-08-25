package vm

//go:generate go run generators/generate_binary_ops.go

import (
	"fmt"
	"io"

	"github.com/danwhitford/golox/chunk"
	"github.com/danwhitford/golox/debug"
	"github.com/danwhitford/golox/stack"
	"github.com/danwhitford/golox/value"
)

type InterpretResult byte

const (
	INTERPRET_OK InterpretResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

type Vm struct {
	Chunk     chunk.Chunk
	Ip        int
	DebugMode bool
	Out       io.Writer
	Stack     stack.Stack[value.Value]
}

func (vm *Vm) Interpret(ch chunk.Chunk) InterpretResult {
	vm.Chunk = ch
	vm.Ip = 0
	return vm.Run()
}

func (vm *Vm) Run() InterpretResult {
	for {
		if vm.DebugMode {
			fmt.Fprintln(vm.Out, " === ")
			for _, v := range vm.Stack {
				fmt.Fprintf(vm.Out, " [ %v ] \n", v)
			}
			fmt.Fprintln(vm.Out, " === ")
			line, _ := debug.DissembleInstruction(vm.Chunk, vm.Ip)
			fmt.Fprintln(vm.Out, line)
		}
		instruction := vm.readByte()
		switch instruction {
		case chunk.OP_CONSTANT:
			{
				constant := vm.readConstant()
				vm.Stack.Push(constant)
			}
		case chunk.OP_ADD:
			{
				err := vm.Add()
				if err != nil {
					return INTERPRET_RUNTIME_ERROR
				}
			}
		case chunk.OP_SUB:
			{
				err := vm.Sub()
				if err != nil {
					return INTERPRET_RUNTIME_ERROR
				}
			}
		case chunk.OP_MULT:
			{
				err := vm.Mult()
				if err != nil {
					return INTERPRET_RUNTIME_ERROR
				}
			}
		case chunk.OP_DIV:
			{
				err := vm.Div()
				if err != nil {
					return INTERPRET_RUNTIME_ERROR
				}
			}
		case chunk.OP_NEGATE:
			{
				v, err := vm.Stack.Pop()
				if err != nil {
					return INTERPRET_RUNTIME_ERROR
				}
				vm.Stack.Push(-v)
			}
		case chunk.OP_RETURN:
			if !vm.Stack.Empty() {
				v, err := vm.Stack.Pop()
				if err != nil {
					return INTERPRET_RUNTIME_ERROR
				}
				fmt.Fprintln(vm.Out, v)
			}
			return INTERPRET_OK
		}
	}
}

func (vm *Vm) readByte() chunk.OpCode {
	instruction := chunk.OpCode(vm.Chunk.Code[vm.Ip])
	vm.Ip++
	return instruction
}

func (vm *Vm) readConstant() value.Value {
	return vm.Chunk.Constants[vm.readByte()]
}
