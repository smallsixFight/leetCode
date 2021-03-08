package basic_calculator

import "testing"

func Test_calculate(t *testing.T) {
	t.Log(calculate("1 + 1"))
	t.Log(calculate("2-1 + 2 "))
	t.Log(calculate("(1+(4+5+2)-3)+(6+8)"))
	t.Log(calculate("(1)"))
	t.Log(calculate("-2+ 1"))
	t.Log(calculate("- (3 + (4 + 5))"))
	t.Log(calculate("- (3 + (4 + 5 + 6))"))
}
