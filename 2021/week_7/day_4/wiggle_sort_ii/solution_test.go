package sort_list

import "testing"

func Test_wiggleSort(t *testing.T) {
	arr := []int{1, 5, 1, 1, 6, 4}
	//wiggleSort(arr)
	//t.Log(arr)
	//arr = []int{1, 3, 2, 2, 3, 1}
	//wiggleSort(arr)
	//t.Log(arr)
	arr = []int{1, 4, 3, 4, 1, 2, 1, 3, 1, 3, 2, 3, 3}
	wiggleSort(arr)
	t.Log(arr)
}
