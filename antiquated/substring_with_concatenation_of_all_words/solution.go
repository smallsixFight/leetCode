package substring_with_concatenation_of_all_words

func FindSubstringThree(s string, words []string) []int {
	res := make([]int, 0)
	if s == "" || len(words) == 0 {
		return res
	}
	dupMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		dupMap[words[i]]++
	}
	wordLen := len(words[0])
	for i := 0; i < wordLen; i++ {
		left, right := i, i
		count := 0
		remainMap := make(map[string]int)
		for right+wordLen <= len(s) {
			if _, ok := dupMap[s[right:right+wordLen]]; ok {
				remainMap[s[right:right+wordLen]]++
				count++
				for remainMap[s[right:right+wordLen]] > dupMap[s[right:right+wordLen]] {
					remainMap[s[left:left+wordLen]]--
					left += wordLen
					count--
				}
			} else {
				right += wordLen
				left = right
				if count != 0 {
					remainMap = make(map[string]int)
					count = 0
				}
				continue
			}
			right += wordLen
			if count == len(words) {
				res = append(res, left)
			}
		}
	}
	return res
}

func FindSubstringTwo(s string, words []string) []int {
	res := make([]int, 0)
	if s == "" || len(words) == 0 {
		return res
	}
	dupMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		dupMap[words[i]]++
	}
	wordsLen := len(words[0])
	remain := make(map[string]int, len(dupMap))
	reset := func() {
		for k, v := range dupMap {
			remain[k] = v
		}
	}

	for i := 0; i < wordsLen; i++ {
		count := 0
		reset()
		left, right := i, i
		for right+wordsLen <= len(s) {
			if _, ok := remain[s[right:right+wordsLen]]; ok {
				remain[s[right:right+wordsLen]]--
				count++
				if remain[s[right:right+wordsLen]] < 0 {
					for remain[s[right:right+wordsLen]] < 0 {
						remain[s[left:left+wordsLen]]++
						left += wordsLen
						count--
					}
				}
			} else {
				left = right + wordsLen
				right = left
				if count != 0 {
					reset()
					count = 0
				}
				continue
			}
			right += wordsLen
			if count == len(words) {
				res = append(res, left)
			}
		}
	}
	return res
}

func FindSubstring(s string, words []string) []int {
	res := make([]int, 0)
	if s == "" || len(words) == 0 {
		return res
	}
	dupMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		dupMap[words[i]]++
	}
	wordsLen := len(words[0])
	recordMap := make(map[int]string, len(s))
	for i := 0; i < len(s)-wordsLen+1; i++ {
		for j := 0; j < len(words); j++ {
			if s[i:i+wordsLen] == words[j] {
				recordMap[i] = words[j]
				break
			}
		}
	}
	for k := range recordMap {
		count := 0
		tempMap := make(map[string]int)
		for a := 0; k+a < len(s) && count < len(words); a += wordsLen {
			if v, ok := recordMap[k+a]; ok {
				tempMap[v]++
				if tempMap[v] > dupMap[v] {
					break
				}
				count++
				continue
			}
			break
		}
		if count == len(words) {
			res = append(res, k)
		}
	}
	return res
}
