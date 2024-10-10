package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/danwhitford/golox/vm"
)

func repl(vm *vm.Vm) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(">>> ")
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {
			fmt.Print(">>> ")
			continue
		}

		// fmt.Fprintf(vm.Out, "[ GOT '%s' ]\n", line)
		res := vm.Interpret(line)
		fmt.Println(res)
		fmt.Print(">>> ")
	}
}

func main() {
	fmt.Println("ARGS", os.Args)

	vm := vm.InitVm()
	vm.DebugMode = true
	repl(vm)
}
