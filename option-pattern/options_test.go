package options

import (
	"testing"
)

func TestOptions(t *testing.T) {
	conn, err := NewConnection("localhost:8080", WithTimeout(100), WithCaching())
	if err != nil {
		t.Fatalf("Create Connection Failed!,err=%v", err)
	}

	if !(conn.timeout == 100 && conn.cache) {
		t.Errorf("Options are not setting properly.")
	}
}
