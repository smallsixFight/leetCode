package course_schedule

import "testing"

func TestCanFinish(t *testing.T) {
	t.Log(true == CanFinish(3, [][]int{{1, 0}}))
	t.Log(true == CanFinish(2, [][]int{{1, 0}}))
	t.Log(false == CanFinish(2, [][]int{{1, 0}, {0, 1}}))
	t.Log(false == CanFinish(3, [][]int{{0, 2}, {1, 2}, {2, 0}}))
}
