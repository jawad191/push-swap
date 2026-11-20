package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"push-swap/internal/operations"
	"push-swap/internal/parser"
	"push-swap/internal/stack"
)

func main() {
	// No arguments: do nothing.
	if len(os.Args) == 1 {
		return
	}

	numbers, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		fail()
	}

	a, b := stack.FromSlice(numbers), stack.NewStack()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		op := strings.TrimSuffix(scanner.Text(), "\r")
		if !operations.Apply(op, a, b) {
			fail()
		}
	}
	if scanner.Err() != nil {
		fail()
	}

	if a.IsSorted() && b.IsEmpty() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

// fail prints Error on the error output and exits.
func fail() {
	fmt.Fprintln(os.Stderr, "Error")
	os.Exit(1)
}
