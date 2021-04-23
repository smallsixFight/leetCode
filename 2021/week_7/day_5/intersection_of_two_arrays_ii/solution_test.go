package largest_divisible_subset

import "testing"

func Test_intersect2(t *testing.T) {
	t.Log(intersect2([]int{1, 2, 2, 1}, []int{2, 2}))       // {2, 2}
	t.Log(intersect2([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})) // { 4, 9}
}

func Test_intersect(t *testing.T) {
	t.Log(intersect([]int{1, 2, 2, 1}, []int{2, 2}))       // {2, 2}
	t.Log(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})) // { 4, 9}
}
