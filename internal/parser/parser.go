package parser

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseArgs(args []string) ([]int, error) {
	var numbers []int
	seen := make(map[int]bool)

	for _, arg := range args {
		values := strings.Fields(arg)

		for _, value := range values {
			n, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("invalid number: %s", value)
			}

			if seen[n] {
				return nil, fmt.Errorf("duplicate number: %d", n)
			}

			seen[n] = true
			numbers = append(numbers, n)
		}
	}

	if len(numbers) == 0 {
		return nil, fmt.Errorf("no numbers provided")
	}

	return numbers, nil
}
