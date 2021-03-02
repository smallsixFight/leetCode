package word_break

func WordBreak(s string, wordDict []string) bool {
	if s == "" || len(wordDict) == 0 {
		return false
	}
	m := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = true
	}
	m2 := make(map[int]bool)
	m2[-1] = true
	for i := 0; i < len(s); i++ {
		for k := range m2 {
			if m[s[k+1:i+1]] {
				m2[i] = true
				break
			}
		}
	}
	return m2[len(s)-1]
}
