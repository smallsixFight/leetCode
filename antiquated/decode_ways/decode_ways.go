package decode_ways

func NumDecodingsTwo(s string) int {
	var p1, p2 int
	if len(s) < 2 {
		if len(s) == 1 && s[0] == '0' {
			return 0
		}
		return len(s)
	}
	if s[0] != '0' {
		p1 = 1
		if s[1] == '0' {
			if s[0] == '1' || s[0] == '2' {
				p2 = 1
			}
		} else {
			if s[0] > '2' || (s[0] == '2' && s[1] > '6') {
				p2 = 1
			} else {
				p2 = 2
			}
		}
	}
	cur := p2
	for i := 2; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = p1
			} else {
				cur = 0
			}
		} else {
			if s[i-1] == '0' || s[i-1] > '2' || (s[i-1] == '2' && s[i] > '6') {
				cur = p2
			} else {
				cur = p1 + p2
			}
		}
		p1 = p2
		p2 = cur
	}
	return cur
}

func NumDecodings(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	res := make([]int, 0)
	if s[0] == '0' {
		res = append(res, 0)
		if s[1] == '0' {
			res = append(res, 0)
		} else {
			res = append(res, 1)
		}
	} else {
		res = append(res, 1)
		if s[1] == '0' {
			if s[0] == '2' {
				res = append(res, 0)
			} else {
				res = append(res, 1)
			}
		} else {
			if s[0] == '1' || (s[0] == '2' && s[1] < '7') {
				res = append(res, 2)
			} else {
				res = append(res, 1)
			}
		}
	}
	for i := 2; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] < '1' || s[i-1] > '2' {
				res = append(res, 0)
			} else if s[i-1] == '1' || s[i-1] == '2' {
				res = append(res, res[i-2])
			}
		} else {
			if s[i-1] == '1' || (s[i-1] == '2' && s[i] < '7') {
				res = append(res, res[i-1]+res[i-2])
			} else {
				res = append(res, res[i-1])
			}
		}
	}
	return res[len(res)-1]
}
