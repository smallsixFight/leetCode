package submissions

import "testing"

func Test_isUgly(t *testing.T) {
	t.Log(isUgly(6))
	t.Log(isUgly(8))
	t.Log(isUgly(14))
	t.Log(isUgly(1))
}
