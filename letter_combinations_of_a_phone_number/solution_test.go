package letter_combinations_of_a_phone_number

import "testing"

func TestSolution3(t *testing.T) {
	t.Log(Solution3("23"))
	t.Log(Solution2("213"))
	t.Log(Solution2("234"))
	t.Log(Solution2("222"))
}

func TestSolution2(t *testing.T) {
	t.Log(Solution2("23"))
	t.Log(Solution2("213"))
	t.Log(Solution2("234"))
	t.Log(Solution2("222"))
}

func TestSolution(t *testing.T) {
	t.Log(Solution("23"))
	t.Log(Solution("213"))
	t.Log(Solution("234"))
	t.Log(Solution("222"))
}
