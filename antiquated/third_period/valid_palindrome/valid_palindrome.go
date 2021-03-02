package valid_palindrome

func IsPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	p1, p2 := 0, len(s)-1
	for {
		for s[p1] < 48 || (s[p1] > 57 && s[p1] < 65) || (s[p1] > 90 && s[p1] < 97) || s[p1] > 122 {
			p1++
			if p1 == len(s) {
				break
			}
		}
		for s[p2] < 48 || (s[p2] > 57 && s[p2] < 65) || (s[p2] > 90 && s[p2] < 97) || s[p2] > 122 {
			p2--
			if p2 == 0 {
				break
			}
		}
		if p1 >= p2 {
			break
		}
		a, b := s[p1], s[p2]
		if s[p1] > 96 {
			a = a - byte(32)
		}
		if s[p2] > 96 {
			b = b - byte(32)
		}
		if a != b {
			return false
		}
		p1++
		p2--
	}
	return true
}
