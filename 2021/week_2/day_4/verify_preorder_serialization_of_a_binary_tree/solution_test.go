package verify_preorder_serialization_of_a_binary_tree

import "testing"

func Test_isValidSerialization(t *testing.T) {
	t.Log(isValidSerialization(""))
	t.Log(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	t.Log(isValidSerialization("1,#"))
	t.Log(isValidSerialization("9,#,#,1"))
	t.Log(isValidSerialization("#"))
}

func Test_isValidSerialization2(t *testing.T) {
	t.Log(isValidSerialization2(""))
	t.Log(isValidSerialization2("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	t.Log(isValidSerialization2("1,#"))
	t.Log(isValidSerialization2("9,#,#,1"))
	t.Log(isValidSerialization2("#"))
}

func Test_isValidSerialization3(t *testing.T) {
	t.Log(isValidSerialization3(""))
	t.Log(isValidSerialization3("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	t.Log(isValidSerialization3("1,#"))
	t.Log(isValidSerialization3("9,#,#,1"))
	t.Log(isValidSerialization3("#"))
}
