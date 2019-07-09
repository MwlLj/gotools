package zwebsocket

import (
	"testing"
)

func TestConnect(t *testing.T) {
	Connect("", 4096, 4096)
}
