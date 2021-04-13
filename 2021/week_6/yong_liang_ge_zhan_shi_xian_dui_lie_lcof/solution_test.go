package yong_liang_ge_zhan_shi_xian_dui_lie_lcof

import "testing"

func TestConstructor(t *testing.T) {
	q := Constructor()
	q.AppendTail(3)
	t.Log(q.DeleteHead())
	t.Log(q.DeleteHead())

	t.Log("----------")
	q1 := Constructor()
	t.Log(q1.DeleteHead())
	q1.AppendTail(5)
	q1.AppendTail(2)
	t.Log(q1.DeleteHead())
	t.Log(q1.DeleteHead())
}
