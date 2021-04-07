package design_a_stack_with_increment_operation

import "testing"

func TestConstructor(t *testing.T) {
	stack := Constructor(3)
	stack.Push(1)
	stack.Push(2)
	t.Log(stack.Pop())
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Increment(5, 100)
	stack.Increment(2, 100)
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
}
