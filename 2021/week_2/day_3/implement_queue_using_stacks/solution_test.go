package implement_queue_using_stacks

import "testing"

func TestMyQueue(t *testing.T) {
	mq := Constructor()
	mq.Push(1)
	mq.Push(2)
	mq.Push(3)
	t.Log(mq.Peek())
	t.Log(mq.Pop())
	t.Log(mq.Peek())
	t.Log(mq.Peek())
	t.Log(mq.Empty())
}
