package vm

//go:generate go run generators/generate_binary_ops.go

import (
	"fmt"
	"io"
	"os"

	"github.com/danwhitford/golox/chunk"
	"github.com/danwhitford/golox/compiler"
	"github.com/danwhitford/golox/debug"
	"github.com/danwhitford/golox/stack"
	"github.com/danwhitford/golox/value"
)

//go:generate stringer -type=InterpretResult
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

func InitVm() *Vm {
	var vm Vm
	vm.Out = os.Stdout
	return &vm
}

func (vm *Vm) Interpret(source string) InterpretResult {
	cmp := compiler.Init(source)
	vm.Chunk = cmp.Compile(source)
	vm.Ip = 0
	res := vm.Run()
	return res
}

func (vm *Vm) Run() InterpretResult {
	for {
		if vm.DebugMode {
			fmt.Fprintln(vm.Out, " === ")
			for _, v := range vm.Stack {
				fmt.Fprintf(vm.Out, " [ %v ] \n", v.As)
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
				switch v.T {
				case value.VAL_NUMBER:
					v.As = -(v.AsNumber())
				case value.VAL_BOOL:
					v.As = !(v.AsBool())
				default:
					panic(fmt.Sprintf("don't know how to negate '%v'", v))
				}
				vm.Stack.Push(v)
			}
		case chunk.OP_RETURN:
			if !vm.Stack.Empty() {
				v, err := vm.Stack.Pop()
				if err != nil {
					return INTERPRET_RUNTIME_ERROR
				}
				switch v.T {
				case value.VAL_NIL:
					fmt.Fprintln(vm.Out, "nil")
				default:
					fmt.Fprintln(vm.Out, v.As)
				}
			}
			return INTERPRET_OK
		case chunk.OP_FALSE:
			vm.Stack = append(vm.Stack, value.BoolVal(false))
		case chunk.OP_TRUE:
			vm.Stack = append(vm.Stack, value.BoolVal(true))
		case chunk.OP_NIL:
			vm.Stack = append(vm.Stack, value.NilVal())
		case chunk.OP_LESS:
			b, err := vm.Stack.Pop()
			if err != nil {
				panic(err)
			}
			a, err := vm.Stack.Pop()
			if err != nil {
				panic(err)
			}
			res := a.AsNumber() < b.AsNumber()
			vm.Stack.Push(value.BoolVal(res))
		case chunk.OP_GREATER:
			b, err := vm.Stack.Pop()
			if err != nil {
				panic(err)
			}
			a, err := vm.Stack.Pop()
			if err != nil {
				panic(err)
			}
			res := a.AsNumber() > b.AsNumber()
			vm.Stack.Push(value.BoolVal(res))
		case chunk.OP_EQUAL:
			b, err := vm.Stack.Pop()
			if err != nil {
				panic(err)
			}
			a, err := vm.Stack.Pop()
			if err != nil {
				panic(err)
			}
			res := a.As == b.As
			vm.Stack.Push(value.BoolVal(res))
		default:
			panic(fmt.Sprintf("don't know how to handle OP CODE '%v'", instruction))
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
