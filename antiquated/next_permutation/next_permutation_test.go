package next_permutation

import "testing"

func TestNextPermutation(t *testing.T) {
	nums1 := []int{1, 2, 3}
	NextPermutation(nums1)
	t.Log(nums1)
	nums2 := []int{3, 2, 1}
	NextPermutation(nums2)
	t.Log(nums2)
	nums3 := []int{1, 1, 5}
	NextPermutation(nums3)
	t.Log(nums3)
	nums4 := []int{1, 4, 5, 3, 2, 6}
	NextPermutation(nums4)
	t.Log(nums4)
	nums5 := []int{6, 5, 4, 3, 2, 1, 0}
	NextPermutation(nums5)
	t.Log(nums5)
	nums6 := []int{1, 3, 2}
	NextPermutation(nums6)
	t.Log(nums6)
}
