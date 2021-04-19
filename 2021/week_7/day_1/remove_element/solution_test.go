package sort_list

import (
	"fmt"
	"testing"
)

func Test_removeElement(t *testing.T) {
	arr := []int{3, 2, 2, 3}
	l := removeElement(arr, 3)
	for i := 0; i < l; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
	arr = []int{0, 1, 2, 2, 3, 0, 4, 2}
	l = removeElement(arr, 2)
	for i := 0; i < l; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
