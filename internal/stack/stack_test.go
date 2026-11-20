package stack

import "testing"

func TestNewStack(t *testing.T) {
	s := NewStack()

	if !s.IsEmpty() {
		t.Error("new stack should be empty")
	}
}

func TestPush(t *testing.T) {
	s := NewStack()
	s.Push(5)
	s.Push(3)

	if s.Numbers[0] != 3 || s.Numbers[1] != 5 {
		t.Errorf("expected [3 5], got %v", s.Numbers)
	}
}

func TestPop(t *testing.T) {
	s := NewStack()
	s.Push(5)
	s.Push(3)

	n, ok := s.Pop()

	if !ok {
		t.Fatal("Pop should return true")
	}

	if n != 3 {
		t.Errorf("expected 3, got %d", n)
	}

	if len(s.Numbers) != 1 || s.Numbers[0] != 5 {
		t.Errorf("expected [5], got %v", s.Numbers)
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := NewStack()

	_, ok := s.Pop()

	if ok {
		t.Error("Pop on empty stack should return false")
	}
}

func TestSwap(t *testing.T) {
	s := Stack{Numbers: []int{3, 5, 8}}
	s.Swap()

	if s.Numbers[0] != 5 || s.Numbers[1] != 3 {
		t.Errorf("expected [5 3 8], got %v", s.Numbers)
	}
}

func TestSwapOneElement(t *testing.T) {
	s := Stack{Numbers: []int{5}}
	s.Swap()

	if len(s.Numbers) != 1 || s.Numbers[0] != 5 {
		t.Errorf("unexpected result: %v", s.Numbers)
	}
}

func TestRotate(t *testing.T) {
	s := Stack{Numbers: []int{1, 2, 3}}
	s.Rotate()

	expected := []int{2, 3, 1}

	for i := range expected {
		if s.Numbers[i] != expected[i] {
			t.Errorf("expected %v, got %v", expected, s.Numbers)
			break
		}
	}
}

func TestReverseRotate(t *testing.T) {
	s := Stack{Numbers: []int{1, 2, 3}}
	s.ReverseRotate()

	expected := []int{3, 1, 2}

	for i := range expected {
		if s.Numbers[i] != expected[i] {
			t.Errorf("expected %v, got %v", expected, s.Numbers)
			break
		}
	}
}

func TestRotateOneElement(t *testing.T) {
	s := Stack{Numbers: []int{5}}
	s.Rotate()

	if s.Numbers[0] != 5 {
		t.Errorf("expected [5], got %v", s.Numbers)
	}
}

func TestReverseRotateOneElement(t *testing.T) {
	s := Stack{Numbers: []int{5}}
	s.ReverseRotate()

	if s.Numbers[0] != 5 {
		t.Errorf("expected [5], got %v", s.Numbers)
	}
}
