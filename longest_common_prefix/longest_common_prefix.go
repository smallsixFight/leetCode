package longest_common_prefix

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	commonPrefix := func(s1, s2 string) string {
		var min int
		if len(s1) < len(s2) {
			min = len(s1)
		} else {
			min = len(s2)
		}
		result := ""
		for i := 0; i < min; i++ {
			if s1[i] == s2[i] {
				result = s1[:i+1]
			} else {
				return result
			}
		}
		return result
	}
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result = commonPrefix(result, strs[i])
		if result == "" {
			break
		}
	}
	return result
}

//func commonPrefix(s1, s2 string) string {
//
//}
