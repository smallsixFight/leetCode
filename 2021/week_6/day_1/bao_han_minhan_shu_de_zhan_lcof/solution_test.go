package bao_han_minhan_shu_de_zhan_lcof

import "testing"

func TestConstructor(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	t.Log(stack.Min())
	stack.Pop()
	t.Log(stack.Top())
	t.Log(stack.Min())
}
