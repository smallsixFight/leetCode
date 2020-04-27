package find_minimum_in_rotated_sorted_array

import "testing"

func TestFindMin(t *testing.T) {
	t.Log(FindMin([]int{2, 1}))
	t.Log(FindMin([]int{3, 4, 5, 1, 2}))
	t.Log(FindMin([]int{4, 5, 6, 7, 0, 1, 2}))
}
