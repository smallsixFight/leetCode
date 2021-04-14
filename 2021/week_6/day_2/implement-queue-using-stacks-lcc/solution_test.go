package implement_queue_using_stacks_lcc

import "testing"

func TestConstructor(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Empty())
}
