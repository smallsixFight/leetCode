package min_stack

import "testing"

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(3)
	minStack.Push(2)
	minStack.Push(1)
	t.Log(minStack.GetMin())
	minStack.Pop()
	t.Log(minStack.GetMin())
}
