package compare_version_numbers

import "testing"

func TestCompareVersion(t *testing.T) {
	t.Log(-1 == CompareVersion("0.1", "1.1"))
	t.Log(1 == CompareVersion("1.0.1", "1"))
	t.Log(-1 == CompareVersion("7.5.2.4", "7.5.3"))
	t.Log(0 == CompareVersion("1.01", "1.001"))
	t.Log(0 == CompareVersion("1.0", "1.0.0"))
}
