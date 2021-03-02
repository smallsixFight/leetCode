package implement_strstr

func Strstr2(haystack, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if needle == "" {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) <= len(haystack) &&
			haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func Strstr(haystack, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if needle == "" {
		return 0
	}
	next := getNext(needle)
	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(needle) {
		return i - j
	}
	return -1
}

func getNext(s string) []int {
	next := make([]int, len(s))
	next[0] = -1
	k, j := -1, 0
	for j < len(s)-1 {
		if k == -1 || s[k] == s[j] {
			k++
			j++
			if s[k] == s[j] {
				next[j] = next[k]
			} else {
				next[j] = k
			}
		} else {
			k = next[k]
		}
	}
	return next
}
