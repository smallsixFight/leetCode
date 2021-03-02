package evaluate_reverse_polish_notation

func EvalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "*", "/", "+", "-":
			v1 := stack[len(stack)-1]
			v2 := stack[len(stack)-2]
			if tokens[i] == "*" {
				stack[len(stack)-2] = v2 * v1
			} else if tokens[i] == "/" {
				stack[len(stack)-2] = v2 / v1
			} else if tokens[i] == "+" {
				stack[len(stack)-2] = v2 + v1
			} else {
				stack[len(stack)-2] = v2 - v1
			}
			stack = stack[:len(stack)-1]
		default:
			isNe := false
			num := 0
			for j := 0; j < len(tokens[i]); j++ {
				if tokens[i][j] == '-' {
					isNe = true
				} else if tokens[i][j] >= '0' && tokens[i][j] <= '9' {
					num = num*10 + int(tokens[i][j]-'0')
				}
			}
			if isNe {
				stack = append(stack, -num)
			} else {
				stack = append(stack, num)
			}
		}
	}
	return stack[0]
}
