package number_of_atoms

import "testing"

func Test_countOfAtoms(t *testing.T) {
	t.Log(countOfAtoms("H2O"))
	t.Log(countOfAtoms("Mg(OH)2"))
	t.Log(countOfAtoms("K4(ON(SO3)2)2"))
	t.Log(countOfAtoms("Be32"))
}
