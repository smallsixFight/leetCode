package kth_largest_element_in_an_array

import (
	"sort"
)

func FindKTHLargestThree(nums []int, k int) int {
	maxHeapSort(nums)
	return nums[len(nums)-k]
}

func maxHeapSort(nums []int) {
	for i := len(nums) / 2; i >= 0; i-- {
		maxHeapify(nums, i)
	}
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums[:i], 0)
	}
}

func maxHeapify(arr []int, i int) {
	l, r := 2*i+1, 2*i+2
	largest := 0
	if l < len(arr) && arr[l] > arr[i] {
		largest = l
	} else {
		largest = i
	}
	if r < len(arr) && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, largest)
	}
}

func FindKTHLargestTwo(nums []int, k int) int {
	nums = mergeSort(nums)
	return nums[len(nums)-k]
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	a1 := mergeSort(nums[:len(nums)/2])
	a2 := mergeSort(nums[len(nums)/2:])
	res := make([]int, len(nums))
	i, j := 0, 0
	idx := 0
	for i < len(a1) || j < len(a2) {
		if i == len(a1) {
			res[idx] = a2[j]
			j++
		} else if j == len(a2) {
			res[idx] = a1[i]
			i++
		} else if a1[i] < a2[j] {
			res[idx] = a1[i]
			i++
		} else {
			res[idx] = a2[j]
			j++
		}
		idx++
	}
	return res
}

func FindKTHLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
