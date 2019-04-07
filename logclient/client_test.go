package logclient

import (
	"testing"
	"time"
)

func TestPrintln(t *testing.T) {
	client := new("localhost", 50005, "tests", "1.0", "1")
	for {
		client.Println("hello, %v", "liujun")
		time.Sleep(1 * time.Second)
	}
}
