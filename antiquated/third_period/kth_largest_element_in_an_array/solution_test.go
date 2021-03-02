package kth_largest_element_in_an_array

import "testing"

func TestFindKTHLargestThree(t *testing.T) {
	t.Log(FindKTHLargestThree([]int{3, 2, 1, 5, 6, 4}, 2))
	t.Log(FindKTHLargestThree([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func TestFindKTHLargestTwo(t *testing.T) {
	t.Log(FindKTHLargestTwo([]int{3, 2, 1, 5, 6, 4}, 2))
	t.Log(FindKTHLargestTwo([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func TestFindKTHLargest(t *testing.T) {
	t.Log(FindKTHLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	t.Log(FindKTHLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
