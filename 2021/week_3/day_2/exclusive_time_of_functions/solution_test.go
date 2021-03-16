package exclusive_time_of_functions

import "testing"

func Test_exclusiveTime(t *testing.T) {
	t.Log(exclusiveTime(2, []string{"0:start:0",
		"1:start:2",
		"1:end:5",
		"0:end:6"}))
	t.Log(exclusiveTime(2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:7", "1:end:7", "0:end:8"}))
}
