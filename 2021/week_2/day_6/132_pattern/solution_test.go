package _32_pattern

import "testing"

func Test_find132pattern(t *testing.T) {
	t.Log(find132pattern([]int{1, 2, 3, 4}))
	t.Log(find132pattern([]int{3, 1, 4, 2}))
	t.Log(find132pattern([]int{-1, 3, 2, 0}))
}
