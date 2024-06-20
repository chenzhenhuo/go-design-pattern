package proxy

import "testing"

func TestAa(t *testing.T) {
	NewUserV2Proxy(&UserV2{}).Hello()
}
