package pascals_triangle_ii

// 符合杨辉三角定义的函数
func GetRowTwo(rowIndex int) []int {
	if rowIndex == 0 {
		return nil
	} else if rowIndex == 1 {
		return []int{1}
	}
	res := make([]int, rowIndex)
	res[0], res[1] = 1, 1
	for i := 2; i < rowIndex; i++ {
		res[i] = 1
		for k := i - 1; k > 0; k-- {
			res[k] += res[k-1]
		}
	}
	return res
}

func GetRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	}
	res := make([]int, rowIndex+1)
	res[0], res[1], res[2] = 1, 2, 1
	for i := 3; i < rowIndex+1; i++ {
		res[i] = 1
		for k := i - 1; k > 0; k-- {
			res[k] += res[k-1]
		}
	}
	return res
}
