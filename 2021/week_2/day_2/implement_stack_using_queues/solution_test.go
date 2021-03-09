package implement_stack_using_queues

import "testing"

func TestMyStack(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	t.Log(stack.Top())
	t.Log(stack.Top())
	t.Log(stack.Empty())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Empty())
}

func TestMyStack2(t *testing.T) {
	stack := MyStack{}
	stack.Push(1)
	t.Log(stack.Pop())
	t.Log(stack.Empty())
}
