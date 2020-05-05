package course_schedule_ii

import "testing"

func TestFindOrder(t *testing.T) {
	//t.Log(FindOrder(2, [][]int{{1, 0}}))
	t.Log(FindOrder(2, [][]int{{0, 1}, {1, 0}}))
	//t.Log(FindOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
