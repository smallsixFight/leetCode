package sort_list

import "testing"

func Test_intersection(t *testing.T) {
	t.Log(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	t.Log(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
