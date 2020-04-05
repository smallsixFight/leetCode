package restore_ip_addresses

import "testing"

func TestRestoreIPAddresses(t *testing.T) {
	res := RestoreIPAddresses("25525511135")
	t.Log(res)
	t.Log(RestoreIPAddresses("0000"))
	t.Log(RestoreIPAddresses("010010"))
	t.Log(RestoreIPAddresses("172162541"))
}
