package remove_duplicates_from_sorted_array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	t.Log(1 == RemoveDuplicates([]int{1}))
	t.Log(2 == RemoveDuplicates([]int{1, 1, 2}))
	t.Log(5 == RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	t.Log(6 == RemoveDuplicates([]int{1, 2, 3, 4, 4, 5, 6, 6, 6}))
}
