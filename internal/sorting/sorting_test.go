package sorting

import (
	"reflect"
	"testing"

	"push-swap/internal/stack"
)

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    bool
	}{
		{"sorted numbers", []int{1, 2, 3, 4}, true},
		{"unsorted numbers", []int{3, 1, 2}, false},
		{"one number", []int{5}, true},
		{"empty stack", []int{}, true},
		{"equal numbers", []int{2, 2, 2}, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := stack.Stack{Numbers: test.numbers}

			got := IsSorted(&a)

			if got != test.want {
				t.Errorf("expected %v, got %v", test.want, got)
			}
		})
	}
}

func TestRangs(t *testing.T) {
	numbers := []int{5, 2, 8, 1, 4}
	expected := []int{3, 1, 4, 0, 2}

	got := rangs(numbers)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}

	if !reflect.DeepEqual(numbers, []int{5, 2, 8, 1, 4}) {
		t.Errorf("original numbers were changed: %v", numbers)
	}
}

func TestSortAlreadySorted(t *testing.T) {
	a := stack.Stack{Numbers: []int{1, 2, 3}}
	b := stack.NewStack()

	instructions := Sort(&a, &b)

	if len(instructions) != 0 {
		t.Errorf("expected no instructions, got %v", instructions)
	}

	if !reflect.DeepEqual(a.Numbers, []int{1, 2, 3}) {
		t.Errorf("stack A was changed: %v", a.Numbers)
	}
}

func TestSortTwoNumbers(t *testing.T) {
	a := stack.Stack{Numbers: []int{2, 1}}
	b := stack.NewStack()

	instructions := Sort(&a, &b)

	if !reflect.DeepEqual(a.Numbers, []int{0, 1}) {
		t.Errorf("expected [0 1], got %v", a.Numbers)
	}

	if !b.IsEmpty() {
		t.Errorf("stack B should be empty: %v", b.Numbers)
	}

	if !reflect.DeepEqual(instructions, []string{"sa"}) {
		t.Errorf("expected [sa], got %v", instructions)
	}
}

func TestSortThreeNumbers(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    []int
	}{
		{"case 1", []int{2, 1, 3}, []int{0, 1, 2}},
		{"case 2", []int{3, 1, 2}, []int{0, 1, 2}},
		{"case 3", []int{3, 2, 1}, []int{0, 1, 2}},
		{"case 4", []int{1, 3, 2}, []int{0, 1, 2}},
		{"case 5", []int{2, 3, 1}, []int{0, 1, 2}},
		{"already sorted", []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := stack.Stack{Numbers: test.numbers}
			b := stack.NewStack()

			instructions := Sort(&a, &b)

			if !reflect.DeepEqual(a.Numbers, test.want) {
				t.Errorf("expected %v, got %v", test.want, a.Numbers)
			}

			if !b.IsEmpty() {
				t.Errorf("stack B should be empty: %v", b.Numbers)
			}

			if len(instructions) > 2 {
				t.Errorf("expected at most 2 instructions, got %d", len(instructions))
			}
		})
	}
}

func TestSortFourAndFiveNumbers(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
	}{
		{"four numbers", []int{4, 2, 3, 1}},
		{"five numbers", []int{5, 2, 9, 1, 7}},
		{"reverse order", []int{5, 4, 3, 2, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := stack.Stack{Numbers: test.numbers}
			b := stack.NewStack()

			Sort(&a, &b)

			expected := make([]int, len(test.numbers))
			for i := range expected {
				expected[i] = i
			}

			if !reflect.DeepEqual(a.Numbers, expected) {
				t.Errorf("expected %v, got %v", expected, a.Numbers)
			}

			if !b.IsEmpty() {
				t.Errorf("stack B should be empty: %v", b.Numbers)
			}
		})
	}
}

func TestSortMoreThanFiveNumbers(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
	}{
		{"six numbers", []int{6, 3, 5, 1, 4, 2}},
		{"eight numbers", []int{8, 3, 6, 1, 7, 2, 5, 4}},
		{"ten numbers", []int{10, 4, 8, 2, 6, 1, 9, 3, 7, 5}},
		{"twenty five numbers", []int{
			25, 12, 7, 19, 3, 22, 1, 15, 9, 5,
			18, 11, 24, 2, 14, 6, 20, 4, 17, 8,
			23, 10, 16, 13, 21,
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := stack.Stack{Numbers: test.numbers}
			b := stack.NewStack()

			Sort(&a, &b)

			expected := make([]int, len(test.numbers))
			for i := range expected {
				expected[i] = i
			}

			if !reflect.DeepEqual(a.Numbers, expected) {
				t.Errorf("expected %v, got %v", expected, a.Numbers)
			}

			if !b.IsEmpty() {
				t.Errorf("stack B should be empty: %v", b.Numbers)
			}
		})
	}
}

func TestTailleDeBloc(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{10, 15},
		{100, 15},
		{101, 30},
		{500, 30},
	}

	for _, test := range tests {
		got := tailleDeBloc(test.n)

		if got != test.want {
			t.Errorf("tailleDeBloc(%d): expected %d, got %d",
				test.n, test.want, got)
		}
	}
}

func TestPositionDuMinimumA(t *testing.T) {
	a := stack.Stack{Numbers: []int{4, 2, 5, 1, 3}}

	got := positionDuMinimumA(&a)

	if got != 3 {
		t.Errorf("expected position 3, got %d", got)
	}
}

func TestPositionDuMaximumB(t *testing.T) {
	b := stack.Stack{Numbers: []int{2, 5, 1, 4, 3}}

	got := positionDuMaximumB(&b)

	if got != 1 {
		t.Errorf("expected position 1, got %d", got)
	}
}
