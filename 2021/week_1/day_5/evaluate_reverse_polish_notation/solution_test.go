package evaluate_reverse_polish_notation

import "testing"

func Test_evalRPN(t *testing.T) {
	t.Log(evalRPN([]string{"2", "1", "+", "3", "*"}))
	t.Log(evalRPN([]string{"4", "13", "5", "/", "+"}))
	t.Log(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
