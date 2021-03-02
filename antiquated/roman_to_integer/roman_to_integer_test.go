package roman_to_integer

import "testing"

func TestSolution(t *testing.T) {
	t.Log(Solution("III"), Solution("III") == 3)
	t.Log(Solution("IV"), Solution("IV") == 4)
	t.Log(Solution("IX"), Solution("IX") == 9)
	t.Log(Solution("LVIII"), Solution("LVIII") == 58)
	t.Log(Solution("MCMXCIV"), Solution("MCMXCIV") == 1994)
}

func TestSolution2(t *testing.T) {
	t.Log(Solution("III"), Solution2("III") == 3)
	t.Log(Solution("IV"), Solution2("IV") == 4)
	t.Log(Solution("IX"), Solution2("IX") == 9)
	t.Log(Solution("LVIII"), Solution2("LVIII") == 58)
	t.Log(Solution("MCMXCIV"), Solution2("MCMXCIV") == 1994)
}
