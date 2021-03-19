package backspace_string_compare

import "testing"

func Test_backspaceCompare2(t *testing.T) {
	t.Log(backspaceCompare2("ab#c", "ad#c"))
	t.Log(backspaceCompare2("ab##", "c#d#"))
	t.Log(backspaceCompare2("a##c", "#a#c"))
	t.Log(backspaceCompare2("a#c", "b"))

	t.Log(backspaceCompare2("y#fo##f", "y#f#o##f"))
	t.Log(backspaceCompare2("bxj##tw", "bxj###tw"))
	t.Log(backspaceCompare2("j##yc##bs#srqpfzantto###########i#mwb", "j##yc##bs#srqpf#zantto###########i#mwb"))
}

func Test_backspaceCompare(t *testing.T) {
	t.Log(backspaceCompare("ab#c", "ad#c"))
	t.Log(backspaceCompare("ab##", "c#d#"))
	t.Log(backspaceCompare("a##c", "#a#c"))
	t.Log(backspaceCompare("a#c", "b"))

	t.Log(backspaceCompare("y#fo##f", "y#f#o##f"))
	t.Log(backspaceCompare("bxj##tw", "bxj###tw"))
	t.Log(backspaceCompare("j##yc##bs#srqpfzantto###########i#mwb", "j##yc##bs#srqpf#zantto###########i#mwb"))
}
