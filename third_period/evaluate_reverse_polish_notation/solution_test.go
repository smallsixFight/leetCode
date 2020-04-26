package evaluate_reverse_polish_notation

import "testing"

func TestEvalRPN(t *testing.T) {
	t.Log(EvalRPN([]string{"2", "1", "+", "3", "*"}))
	t.Log(EvalRPN([]string{"4", "13", "5", "/", "+"}))
	t.Log(EvalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
