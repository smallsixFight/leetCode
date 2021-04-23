package largest_divisible_subset

import "testing"

func Test_largestDivisibleSubset(t *testing.T) {
	t.Log(largestDivisibleSubset([]int{1, 2, 3}))       // {1, 2}
	t.Log(largestDivisibleSubset([]int{1, 2, 4, 8}))    // {1, 2, 4, 8}
	t.Log(largestDivisibleSubset([]int{1, 2, 4, 6, 8})) // {1, 2, 4, 8}
	t.Log(largestDivisibleSubset([]int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720}))
}
