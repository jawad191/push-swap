package stack

type Stack struct {
	Numbers []int
}

func NewStack() Stack {
	return Stack{
		Numbers: []int{},
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.Numbers) == 0
}

func (s *Stack) Push(n int) {
	s.Numbers = append([]int{n}, s.Numbers...)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	n := s.Numbers[0]
	s.Numbers = s.Numbers[1:]

	return n, true
}

func (s *Stack) Swap() {
	if len(s.Numbers) < 2 {
		return
	}

	s.Numbers[0], s.Numbers[1] = s.Numbers[1], s.Numbers[0]
}

func (s *Stack) Rotate() {
	if len(s.Numbers) < 2 {
		return
	}

	first := s.Numbers[0]
	s.Numbers = append(s.Numbers[1:], first)
}

func (s *Stack) ReverseRotate() {
	if len(s.Numbers) < 2 {
		return
	}

	last := s.Numbers[len(s.Numbers)-1]
	s.Numbers = append([]int{last}, s.Numbers[:len(s.Numbers)-1]...)
}
