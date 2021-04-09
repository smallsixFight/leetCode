package crawler_log_folder

import "testing"

func Test_minOperations(t *testing.T) {
	t.Log(minOperations([]string{"d1/", "d2/", "../", "d21/", "./"}))
	t.Log(minOperations([]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}))
	t.Log(minOperations([]string{"d1/", "../", "../", "../"}))
}
