package combinations

import "testing"

func TestCombine(t *testing.T) {
	res := Combine(4, 2)
	t.Log(res)
	res = Combine(3, 2)
	t.Log(res)
	res = Combine(3, 3)
	t.Log(res)
	res = Combine(4, 4)
	t.Log(res)
	res = Combine(3, 0)
	t.Log(res)
	res = Combine(5, 2)
	t.Log(res)
	res = Combine(7, 4)
	t.Log(res)
}
