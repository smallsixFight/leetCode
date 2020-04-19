package word_ladder

func LadderLength(beginWord, endWord string, wordList []string) int {
	isExist := false
	m := make(map[string]bool)
	commonStatus := make(map[string][]string)
	for i := 0; i < len(wordList); i++ {
		if wordList[i] == endWord {
			isExist = true
		}
		if wordList[i] == beginWord {
			continue
		}
		for j := 0; j < len(wordList[i]); j++ {
			newWord := wordList[i][:j] + "*" + wordList[i][j+1:]
			if _, ok := commonStatus[newWord]; ok {
				commonStatus[newWord] = append(commonStatus[newWord], wordList[i])
			} else {
				commonStatus[newWord] = []string{wordList[i]}
			}
		}
		m[wordList[i]] = false
	}
	if !isExist {
		return 0
	}
	queue := []string{beginWord}
	res := -1
	for len(queue) > 0 {
		res += 1
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i] == endWord {
				return res + 1
			}
			for j := 0; j < len(beginWord); j++ {
				newWord := queue[i][:j] + "*" + queue[i][j+1:]
				if v, ok := commonStatus[newWord]; ok {
					for k := 0; k < len(v); k++ {
						if !m[v[k]] {
							queue = append(queue, v[k])
							m[v[k]] = true
						}
					}
				}
			}
		}
		queue = queue[l:]
	}
	return 0
}
