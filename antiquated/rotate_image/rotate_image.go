package rotate_image

func Rotate(matrix [][]int) {
	n := len(matrix)
	for l := 0; l < n/2; l++ {
		for s := l; s < n-l-1; s++ {
			i, j := l, s
			t := matrix[i][s]
			for k := 0; k < 4; k++ {
				i, j = j, n-i-1
				temp := matrix[i][j]
				matrix[i][j] = t
				t = temp
			}

		}
	}
}
