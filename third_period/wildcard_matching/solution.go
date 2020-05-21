package wildcard_matching

func IsMatch(s, p string) bool {
	p1, p2 := 0, 0
	backIdx := -1
	sBackIdx := 0
	for p1 < len(s) {
		if p2 < len(p) && (s[p1] == p[p2] || p[p2] == '?') {
			p1++
			p2++
		} else if p2 < len(p) && p[p2] == '*' {
			backIdx = p2
			p2++
			sBackIdx = p1
		} else if backIdx == -1 {
			return false
		} else {
			p1 = sBackIdx + 1
			p2 = backIdx + 1
			sBackIdx = p1
		}
	}
	for p2 < len(p) {
		if p[p2] != '*' {
			return false
		}
		p2++
	}
	return true
}
