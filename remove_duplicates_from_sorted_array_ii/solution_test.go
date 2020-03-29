package remove_duplicates_from_sorted_array_ii

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	res := RemoveDuplicates(nums)
	t.Log(nums[:res])
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	res = RemoveDuplicates(nums)
	t.Log(nums[:res])
}
