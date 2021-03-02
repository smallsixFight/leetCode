package palindrome_partitioning

var str string
var res [][]string

func Partition(s string) [][]string {
	str = s
	record := make(map[[2]int]bool)
	isPalindrome := func(startIdx, endIdx int) {
		sa, ea := startIdx, endIdx
		for startIdx <= endIdx {
			if s[startIdx] != s[endIdx] {
				return
			}
			startIdx++
			endIdx--
		}
		record[[2]int{sa, ea}] = true
	}
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			isPalindrome(i, j)
		}
	}
	res = make([][]string, 0)
	for i := 0; i < len(str); i++ {
		dfs(record, []string{}, 0, i)
	}
	return res
}

func dfs(record map[[2]int]bool, arr []string, a, b int) {
	if a == len(str) || b == len(str) || !record[[2]int{a, b}] {
		return
	}
	arr = append(arr, str[a:b+1])
	if b == len(str)-1 {
		temp := make([]string, len(arr))
		copy(temp, arr)
		res = append(res, temp)
		return
	}
	for i := b + 1; i < len(str); i++ {
		dfs(record, arr, b+1, i)
	}
}
