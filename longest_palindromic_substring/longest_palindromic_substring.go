package longest_palindromic_substring

// 中心扩展算法
func LongestPalindrome2(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		l := 0
		if len1 > len2 {
			l = len1
		} else {
			l = len2
		}
		if l > end-start {
			start = i - (l-1)/2 // 奇数时减少的值相同
			end = i + l/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	record := make([][]int, 0)
	res := s[0:1]
	for i := 0; i < len(s); i++ {
		record = append(record, []int{i, i})
		if i < len(s)-1 && s[i] == s[i+1] {
			temp := s[i : i+2]
			if len(res) == 1 {
				res = temp
			}
			record = append(record, []int{i, i + 1})
		}
	}
	for len(record) > 0 {
		v := record[0]
		record = record[1:]
		if v[0]-1 >= 0 && v[1]+1 < len(s) {
			if s[v[0]-1] == s[v[1]+1] {
				record = append(record, []int{v[0] - 1, v[1] + 1})
				temp := s[v[0]-1 : v[1]+2]
				if len(temp) > len(res) {
					res = temp
				}
			}
		}
	}
	return res
}
