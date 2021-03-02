package divide_two_integers

import "testing"

func TestDivide(t *testing.T) {
	t.Log(3 == Divide(10, 3))
	t.Log(-2 == Divide(7, -3))
	t.Log(-3 == Divide(10, -3))
	t.Log(1 == Divide(1, 1))
	t.Log(2147483647 == Divide(-2147483648, -1))
	t.Log(-2147483648 == Divide(-2147483648, 1))
}
