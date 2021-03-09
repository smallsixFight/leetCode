package basic_calculator_ii

import "testing"

func Test_calculate(t *testing.T) {
	t.Log(calculate("3+2*2") == 7)
	t.Log(calculate("3/2") == 1)
	t.Log(calculate("3+5 / 2 ") == 5)
	t.Log(calculate("3+(5 / 2) ") == 5)
	t.Log(calculate("3+(5 - 2)/2 ") == 4)
	t.Log(calculate("42") == 42)
	t.Log(calculate("21+2*(1+2)") == 27)
	t.Log(calculate("1-1+1") == 1)
	t.Log(calculate("2*3+4") == 10)
	t.Log(calculate("1*2-3/4+5*6-7*8+9/10"))
}
