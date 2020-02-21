package roman_to_integer

func Solution2(s string) int {
	rv := func(s byte) int {
		switch s {
		case 'I':
			return 1
		case 'V':
			return 5
		case 'X':
			return 10
		case 'L':
			return 50
		case 'C':
			return 100
		case 'D':
			return 500
		case 'M':
			return 1000
		}
		return 0
	}

	target := 0
	for i, l := 0, len(s); ; i++ {
		if l == 1 {
			return rv(s[0])
		}

		if i+1 >= l {
			target += rv(s[i])
			break
		}

		if n := rv(s[i]); n >= rv(s[i+1]) {
			target += n
		} else {
			target -= n
		}
	}
	return target
}

func Solution(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == 'I' {
			if i+1 < len(s) && s[i+1] == 'V' {
				result += 4
				i += 1
			} else if i+1 < len(s) && s[i+1] == 'X' {
				result += 9
				i += 1
			} else {
				result += 1
			}
		} else if c == 'X' {
			if i+1 < len(s) && s[i+1] == 'L' {
				result += 40
				i += 1
			} else if i+1 < len(s) && s[i+1] == 'C' {
				result += 90
				i += 1

			} else {
				result += 10
			}
		} else if c == 'C' {
			if i+1 < len(s) && s[i+1] == 'D' {
				result += 400
				i += 1
			} else if i+1 < len(s) && s[i+1] == 'M' {
				result += 900
				i += 1

			} else {
				result += 100
			}
		} else if c == 'V' {
			result += 5
		} else if c == 'L' {
			result += 50
		} else if c == 'D' {
			result += 500
		} else {
			result += 1000
		}
	}
	return result
}
