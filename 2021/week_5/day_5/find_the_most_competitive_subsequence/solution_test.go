package find_the_most_competitive_subsequence

import "testing"

func Test_mostCompetitive(t *testing.T) {
	t.Log(mostCompetitive([]int{3, 5, 2, 6}, 2))
	t.Log(mostCompetitive([]int{2, 4, 3, 3, 5, 4, 9, 6}, 4))
}
