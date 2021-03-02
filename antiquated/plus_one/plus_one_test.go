package plus_one

import "testing"

func TestPlusOne(t *testing.T) {
	t.Log(PlusOne([]int{1, 2, 3}))
	t.Log(PlusOne([]int{4, 3, 2, 1}))
	t.Log(PlusOne([]int{4, 3, 2, 9}))
	t.Log(PlusOne([]int{9, 9, 9, 9}))
}
