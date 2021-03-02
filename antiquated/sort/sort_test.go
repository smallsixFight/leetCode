package sort

import "testing"

func TestMergeSort(t *testing.T) {
	a1 := []int{109, 89, 9, 33, 56, 90, 3, 4, 5, 3, 1, 2, 65}
	MergeSort(a1)
	t.Log(a1)
}

func TestMaxHeapSort(t *testing.T) {
	a1 := []int{109, 89, 9, 33, 56, 90, 3, 4, 5, 3, 1, 2, 65}
	MaxHeapSort(a1)
	t.Log(a1)
}
