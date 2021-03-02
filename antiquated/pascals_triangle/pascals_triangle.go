package pascals_triangle

func Generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	} else if numRows == 1 {
		return [][]int{{1}}
	} else if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	res := make([][]int, 2)
	res[0] = []int{1}
	res[1] = []int{1, 1}
	for i := 3; i <= numRows; i++ {
		arr := make([]int, 0)
		for k := 0; k < i; k++ {
			if k == 0 || k == i-1 {
				arr = append(arr, 1)
			} else {
				arr = append(arr, res[i-2][k-1]+res[i-2][k])
			}
		}
		res = append(res, arr)
	}
	return res
}
