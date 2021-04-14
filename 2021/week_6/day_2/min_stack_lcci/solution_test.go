package min_stack_lcci

import "testing"

func TestConstructor(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	t.Log(stack.GetMin())
	stack.Pop()
	t.Log(stack.Top())
	t.Log(stack.GetMin())
}
