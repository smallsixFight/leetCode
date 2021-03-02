package find_first_and_last_position_of_element_in_sorted_array

import "testing"

func TestSearchRange(t *testing.T) {
	t.Log(SearchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	t.Log(SearchRange([]int{5, 7, 7, 8, 8, 10}, 7))
	t.Log(SearchRange([]int{5, 7, 7, 8, 8, 10}, 0))
	t.Log(SearchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	t.Log(SearchRange([]int{2, 2}, 2))
}
