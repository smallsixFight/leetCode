package two_sum_ii_input_array_is_sorted

func TwoSumTwo(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}
	p1 := 0
	for p1 < len(numbers)-1 && numbers[p1] <= target {
		p2 := p1 + 1
		for p2 < len(numbers)-1 && numbers[p1]+numbers[p2] < target {
			p2++
		}
		if numbers[p1]+numbers[p2] == target {
			return []int{p1 + 1, p2 + 1}
		}
		p1++
	}
	return nil
}

func TwoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return nil
}
