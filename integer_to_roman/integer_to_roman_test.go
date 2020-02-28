package integer_to_roman

import "testing"

func TestIntToRoman2(t *testing.T) {
	t.Log("III" == IntToRoman2(3))
	t.Log("IV" == IntToRoman2(4))
	t.Log("IX" == IntToRoman2(9))
	t.Log("LVIII" == IntToRoman2(58))
	t.Log("MCMXCIV" == IntToRoman2(1994))
}

func TestIntToRoman(t *testing.T) {
	t.Log(IntToRoman(10))
	t.Log(IntToRoman(11))
	t.Log(IntToRoman(1901))

	t.Log("III" == IntToRoman(3))
	t.Log("IV" == IntToRoman(4))
	t.Log("IX" == IntToRoman(9))
	t.Log("LVIII" == IntToRoman(58))
	t.Log("MCMXCIV" == IntToRoman(1994))
}
