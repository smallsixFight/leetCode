package search_a_2d_matrix

func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}
	low, high := 0, m-1
	row := 0
	mid := 0
	for low <= high {
		mid = (low + high) >> 1
		if target == matrix[mid][0] || target == matrix[mid][n-1] {
			return true
		}
		if target > matrix[mid][0] && target < matrix[mid][n-1] {
			row = mid
			break
		} else if target > matrix[mid][n-1] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	low, high = 0, n-1
	for low <= high {
		mid = (low + high) >> 1
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
