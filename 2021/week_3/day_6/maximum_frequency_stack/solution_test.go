package maximum_frequency_stack

import "testing"

func TestFreqStack(t *testing.T) {
	s := Constructor()
	s.Push(5)
	s.Push(7)
	s.Push(5)
	s.Push(7)
	s.Push(4)
	s.Push(5)
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
}
