package guan_mai

import "testing"

func Test_find_seat(t *testing.T) {
	a1 := []int{0, 1, 0, 0, 0, 0, 1}
	t.Log(3 == findSeat(a1))
	t.Log(5 == findSeat([]int{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0}))
	t.Log(3 == findSeat([]int{1, 0, 0, 0, 0, 0, 0, 1, 0, 0}))
	t.Log(10 == findSeat([]int{1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}))
	t.Log(12 == findSeat([]int{1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0}))
	t.Log(5 == findSeat([]int{1, 0, 1, 0, 1, 0, 0, 1}))
	t.Log(6 == findSeat([]int{1, 0, 1, 0, 1, 0, 0, 0, 1}))
}
