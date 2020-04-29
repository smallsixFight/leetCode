package fraction_to_recurring_decimal

import "testing"

func TestFractionToDecimal(t *testing.T) {

	t.Log(FractionToDecimal(1, 2))
	t.Log(FractionToDecimal(1, 3))
	t.Log(FractionToDecimal(2, 3))
	t.Log("0.(064516129032258)" == FractionToDecimal(2, 31))
	t.Log(FractionToDecimal(10, 30))
	t.Log(FractionToDecimal(1, 30))
	t.Log(FractionToDecimal(1, 33))
	t.Log(FractionToDecimal(1, 333))
	t.Log(FractionToDecimal(-50, 8))
	t.Log(FractionToDecimal(7, -12))
	t.Log(FractionToDecimal(89, 270))
	t.Log(FractionToDecimal(1, 214748364))
	t.Log(FractionToDecimal(2147483647, 370000))
	t.Log(7.0 / -12.0)
}
