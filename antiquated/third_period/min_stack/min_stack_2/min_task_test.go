package min_stack_2

import "testing"

func TestMinStack(t *testing.T) {
	stack := Constructor()
	//stack.Pop()
	stack.Push(2)
	stack.Push(1)
	t.Log(stack.GetMin())
	t.Log(stack.Top())
	stack.Pop()
	t.Log(stack.Top())
	t.Log(stack.GetMin())
	stack.Push(-1)
	t.Log(stack.Top())
	t.Log(stack.GetMin())
}
