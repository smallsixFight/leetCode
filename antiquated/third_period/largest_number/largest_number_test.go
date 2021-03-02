package largest_number

import "testing"

func TestLargestNumber(t *testing.T) {
	t.Log(LargestNumber([]int{10, 2}))
	t.Log(LargestNumber([]int{3, 30, 34, 5, 9}))
	t.Log(LargestNumber([]int{0, 0}))
	t.Log(LargestNumber([]int{1, 0, 0}))
	t.Log(LargestNumber([]int{5827, 1331, 4889, 4514, 919, 7832}))
	t.Log(LargestNumber([]int{3, 43, 48, 94, 85, 33, 64, 32, 63, 66}))

}
