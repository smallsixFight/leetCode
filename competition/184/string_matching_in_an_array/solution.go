package string_matching_in_an_array

func StringMatching(words []string) []string {
	if len(words) < 2 {
		return nil
	}
	res := make([]string, 0)
	m := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if len(words[j]) < len(words[i]) {
				for k := 0; k < len(words[i]) && (len(words[i])-k) >= len(words[j]) && !m[words[j]]; k++ {
					if words[i][k:k+len(words[j])] == words[j] {
						res = append(res, words[j])
						m[words[j]] = true
						break
					}
				}
			}
		}
	}
	return res
}
