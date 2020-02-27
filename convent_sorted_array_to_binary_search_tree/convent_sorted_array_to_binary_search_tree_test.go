package convent_sorted_array_to_binary_search_tree

import (
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	arr := []int{-10, -3, 0, 5, 9}
	res := SortedArrayToBST(arr)
	//fmt.Println(res)
	printBinaryTree(res)
}
