package main

import (
	"fmt"
	"os"

	"push-swap/internal/parser"
	"push-swap/internal/sorting"
	"push-swap/internal/stack"
)

func main() {
	// No arguments: do nothing
	if len(os.Args) == 1 {
		return
	}

	// Parse and validate numbers
	numbers, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	// Create stacks
	a := stack.NewStack()
	b := stack.NewStack()

	// Fill stack A
	for i := len(numbers) - 1; i >= 0; i-- {
		a.Push(numbers[i])
	}

	// Sort
	instructions := sorting.Sort(&a, &b)

	// Display instructions
	for _, instruction := range instructions {
		fmt.Println(instruction)
	}
}
