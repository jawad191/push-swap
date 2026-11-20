package operations

import "push-swap/internal/stack"

func Apply(op string, a, b *stack.Stack) bool {
	switch op {
	case "sa":
		a.Swap()
	case "sb":
		b.Swap()
	case "ss":
		a.Swap()
		b.Swap()
	case "pa":
		if n, ok := b.Pop(); ok {
			a.Push(n)
		}
	case "pb":
		if n, ok := a.Pop(); ok {
			b.Push(n)
		}
	case "ra":
		a.Rotate()
	case "rb":
		b.Rotate()
	case "rr":
		a.Rotate()
		b.Rotate()
	case "rra":
		a.ReverseRotate()
	case "rrb":
		b.ReverseRotate()
	case "rrr":
		a.ReverseRotate()
		b.ReverseRotate()
	default:
		return false
	}

	return true
}
