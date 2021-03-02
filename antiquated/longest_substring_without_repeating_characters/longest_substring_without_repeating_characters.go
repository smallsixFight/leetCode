package longest_substring_without_repeating_characters

func LengthOfLengestSubstring(s string) int {
	m := make(map[byte]int)
	maxLength, temp := 0, 0
	for p1, p2 := 0, 0; p2 < len(s); p2++ {
		if _, ok := m[s[p2]]; !ok || m[s[p2]] < p1 {
			m[s[p2]] = p2
			temp++
		} else {
			p1 = m[s[p2]] + 1
			temp = p2 - m[s[p2]]
			m[s[p2]] = p2
		}
		if temp > maxLength {
			maxLength = temp
		}
	}
	return maxLength
}
