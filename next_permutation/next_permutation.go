package next_permutation

func NextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}
	p1, p2 := -1, 0
	for i := len(nums) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] && (j > p1 || p1 == -1) {
				p1 = j
				p2 = i
				break
			}
		}
	}
	if p1 != -1 {
		nums[p1], nums[p2] = nums[p2], nums[p1]
	}
	//sort.Ints(nums[p1+1:])
	maxHeapSort(nums[p1+1:])
}

func maxHeapSort(nums []int) {
	for i := len(nums) / 2; i >= 0; i-- {
		maxHeapify(nums, i)
	}
	for lastIdx := len(nums) - 1; lastIdx >= 0; lastIdx-- {
		nums[0], nums[lastIdx] = nums[lastIdx], nums[0]
		maxHeapify(nums[0:lastIdx], 0)
	}
}

func maxHeapify(arr []int, i int) {
	l, r := 2*i+1, 2*i+2
	var largest int
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
