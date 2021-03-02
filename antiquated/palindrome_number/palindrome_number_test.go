package palindrome_number

import "testing"

func TestFirstMethod(t *testing.T) {
	t.Log(FirstMethod(131))
	t.Log(FirstMethod(-131))
	t.Log(FirstMethod(-10))
}

func TestSecondMethod(t *testing.T) {
	t.Log(SecondMethod(131))
	t.Log(SecondMethod(113311))
	t.Log(SecondMethod(-131))
	t.Log(SecondMethod(-10))
}
