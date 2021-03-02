package two_sum

func FirstMethod(arr []int, target int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == target-arr[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

func SecondMethod(arr []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if _, ok := m[target-arr[i]]; ok {
			return []int{m[target-arr[i]], i}
		}
		m[arr[i]] = i
	}
	return nil
}
