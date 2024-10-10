package vm


func (vm *Vm) Add() error {
    v2, err := vm.Stack.Pop()
    if err != nil {
        return err
    }
    v1, err := vm.Stack.Pop()
    if err != nil {
        return err
    }
    vm.Stack.Push(v1 + v2)

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
    vm.Stack.Push(v1 - v2)

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
    vm.Stack.Push(v1 * v2)

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
    vm.Stack.Push(v1 / v2)

    return nil
}

