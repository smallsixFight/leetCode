package unique_binary_search_trees_II

import (
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	res := GenerateTrees(4)
	t.Log(len(res))
	for k := range res {
		t.Log(PreOrderPrint(res[k]))
	}
}
