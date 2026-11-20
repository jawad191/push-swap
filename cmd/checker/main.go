package main

import (
	"bufio"
	"fmt"
	"os"

	"push-swap/internal/operations"
	"push-swap/internal/parser"
	"push-swap/internal/stack"
)

func main() {
	// No arguments: do nothing
	if len(os.Args) == 1 {
		return
	}

	// Parse numbers
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

	// Read instructions
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		op := scanner.Text()

		if !operations.Apply(op, &a, &b) {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	// Check if stack A is sorted
	for i := 0; i < len(a.Numbers)-1; i++ {
		if a.Numbers[i] > a.Numbers[i+1] {
			fmt.Println("KO")
			return
		}
	}

	// Check if stack B is empty
	if !b.IsEmpty() {
		fmt.Println("KO")
		return
	}

	fmt.Println("OK")
}
