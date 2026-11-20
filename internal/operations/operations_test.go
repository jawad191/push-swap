package operations

import (
	"testing"

	"push-swap/internal/stack"
)

func TestOperationsOnEmptyStacks(t *testing.T) {
	operations := []string{
		"sa", "sb", "ss",
		"pa", "pb",
		"ra", "rb", "rr",
		"rra", "rrb", "rrr",
	}

	for _, op := range operations {
		a := stack.NewStack()
		b := stack.NewStack()

		result := Apply(op, &a, &b)

		if !result {
			t.Errorf("operation %s failed", op)
		}

		if !a.IsEmpty() || !b.IsEmpty() {
			t.Errorf("operation %s changed an empty stack", op)
		}
	}
}

func TestInvalidOperation(t *testing.T) {
	a := stack.NewStack()
	b := stack.NewStack()

	result := Apply("invalid", &a, &b)

	if result {
		t.Error("invalid operation should return false")
	}
}

func TestOperationsOnSingleElementStacks(t *testing.T) {
	tests := []struct {
		name  string
		op    string
		a     []int
		b     []int
		wantA []int
		wantB []int
	}{
		{"sa", "sa", []int{5}, nil, []int{5}, nil},
		{"sb", "sb", nil, []int{8}, nil, []int{8}},
		{"ss", "ss", []int{5}, []int{8}, []int{5}, []int{8}},
		{"ra", "ra", []int{5}, nil, []int{5}, nil},
		{"rb", "rb", nil, []int{8}, nil, []int{8}},
		{"rr", "rr", []int{5}, []int{8}, []int{5}, []int{8}},
		{"rra", "rra", []int{5}, nil, []int{5}, nil},
		{"rrb", "rrb", nil, []int{8}, nil, []int{8}},
		{"rrr", "rrr", []int{5}, []int{8}, []int{5}, []int{8}},
		{"pa", "pa", nil, []int{8}, []int{8}, nil},
		{"pb", "pb", []int{5}, nil, nil, []int{5}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := stack.Stack{Numbers: test.a}
			b := stack.Stack{Numbers: test.b}

			if !Apply(test.op, &a, &b) {
				t.Fatalf("operation %s failed", test.op)
			}

			if !equal(a.Numbers, test.wantA) {
				t.Errorf("A = %v, expected %v", a.Numbers, test.wantA)
			}

			if !equal(b.Numbers, test.wantB) {
				t.Errorf("B = %v, expected %v", b.Numbers, test.wantB)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestOperationsWithMultipleElements(t *testing.T) {
	tests := []struct {
		name  string
		op    string
		a     []int
		b     []int
		wantA []int
		wantB []int
	}{
		{"sa", "sa", []int{3, 1}, []int{}, []int{1, 3}, []int{}},
		{"sb", "sb", []int{}, []int{3, 1}, []int{}, []int{1, 3}},
		{"ss", "ss", []int{3, 1}, []int{5, 2}, []int{1, 3}, []int{2, 5}},
		{"ra", "ra", []int{3, 1, 2}, []int{}, []int{1, 2, 3}, []int{}},
		{"rb", "rb", []int{}, []int{3, 1, 2}, []int{}, []int{1, 2, 3}},
		{"rr", "rr", []int{3, 1, 2}, []int{6, 4, 5}, []int{1, 2, 3}, []int{4, 5, 6}},
		{"rra", "rra", []int{3, 1, 2}, []int{}, []int{2, 3, 1}, []int{}},
		{"rrb", "rrb", []int{}, []int{3, 1, 2}, []int{}, []int{2, 3, 1}},
		{"rrr", "rrr", []int{3, 1, 2}, []int{6, 4, 5}, []int{2, 3, 1}, []int{5, 6, 4}},
		{"pa", "pa", []int{3, 1}, []int{5, 2}, []int{5, 3, 1}, []int{2}},
		{"pb", "pb", []int{3, 1}, []int{5, 2}, []int{1}, []int{3, 5, 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := stack.Stack{Numbers: test.a}
			b := stack.Stack{Numbers: test.b}

			Apply(test.op, &a, &b)

			if !equal(a.Numbers, test.wantA) {
				t.Errorf("A = %v, expected %v", a.Numbers, test.wantA)
			}

			if !equal(b.Numbers, test.wantB) {
				t.Errorf("B = %v, expected %v", b.Numbers, test.wantB)
			}
		})
	}
}
