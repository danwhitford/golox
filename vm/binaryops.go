package vm
import "github.com/danwhitford/golox/value"


func (vm *Vm) Add() error {
    v2, err := vm.Stack.Pop()
    if err != nil {
        return err
    }
    v1, err := vm.Stack.Pop()
    if err != nil {
        return err
    }
    vm.Stack.Push(
        value.NumberVal(v1.AsNumber() + v2.AsNumber()),
    )

    return nil
}

func (vm *Vm) Sub() error {
    v2, err := vm.Stack.Pop()
    if err != nil {
        return err
    }
    v1, err := vm.Stack.Pop()
    if err != nil {
        return err
    }
    vm.Stack.Push(
        value.NumberVal(v1.AsNumber() - v2.AsNumber()),
    )

    return nil
}

func (vm *Vm) Mult() error {
    v2, err := vm.Stack.Pop()
    if err != nil {
        return err
    }
    v1, err := vm.Stack.Pop()
    if err != nil {
        return err
    }
    vm.Stack.Push(
        value.NumberVal(v1.AsNumber() * v2.AsNumber()),
    )

    return nil
}

func (vm *Vm) Div() error {
    v2, err := vm.Stack.Pop()
    if err != nil {
        return err
    }
    v1, err := vm.Stack.Pop()
    if err != nil {
        return err
    }
    vm.Stack.Push(
        value.NumberVal(v1.AsNumber() / v2.AsNumber()),
    )

    return nil
}

