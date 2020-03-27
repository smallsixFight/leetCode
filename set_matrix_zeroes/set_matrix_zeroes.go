package set_matrix_zeroes

func SetZeroes(matrix [][]int) {
	s1, s2 := 1, 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					s1 = 0
				}
				if j == 0 {
					s2 = 0
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := s2; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
	if s1 == 0 {
		for i := 1; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
}
