package queries_on_a_permutation_with_key

func ProcessQueries(queries []int, m int) []int {
	p := make([]int, m)
	record := make(map[int]int)
	for i := 0; i < m; i++ {
		p[i] = i + 1
		record[i+1] = i
	}
	res := make([]int, 0)
	for i := 0; i < len(queries); i++ {
		idx := record[queries[i]]
		res = append(res, idx)
		tempArr := make([]int, idx)
		for k := 0; k < len(tempArr); k++ {
			tempArr[k] = p[k]
		}
		p[0] = p[idx]
		record[p[idx]] = 0
		for k := 1; k < len(tempArr)+1; k++ {
			p[k] = tempArr[k-1]
			record[p[k]] = k
		}
	}
	return res
}
