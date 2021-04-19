package sort_list

import "testing"

func Test_maximumGap2(t *testing.T) {
	t.Log(maximumGap2([]int{3, 6, 9, 1}))
	t.Log(maximumGap2([]int{10}))
}

func Test_maximumGap(t *testing.T) {
	t.Log(maximumGap([]int{3, 6, 9, 1}))
	t.Log(maximumGap([]int{10}))
	//t.Log(140%100/10)
}
