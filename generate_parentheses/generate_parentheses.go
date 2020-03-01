package generate_parentheses

func GenerateParenthesis(n int) []string {
	res := make([]string, 0)
	depthExplore("(", 1, &res, n)
	return res
}

func depthExplore(preStr string, count int, res *[]string, n int) {
	if len(preStr) == 2*n {
		if count == 0 {
			*res = append(*res, preStr)
		}
		return
	}
	if count < n {
		depthExplore(preStr+"(", count+1, res, n)
	}
	if count > 0 {
		depthExplore(preStr+")", count-1, res, n)
	}
}
