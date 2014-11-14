package wtgo

import (
	"testing"
)

// TODO need to cleanup all WiredTiger files created during testing
func TestOpen(t *testing.T) {
	conn, err := Open(".", "create,cache_size=500M")
	if err != nil {
		t.Errorf("Open failed with %v", err)
	} else {
		conn.Close()
	}
}

func TestOpenInvalidOpts(t *testing.T) {
	conn, err := Open(".", "create,cache_size=500M,meow=mooooooo")
	t.Logf("err == nil: %v", err == nil)
	if err == nil {
		defer conn.Close()
		t.Fatalf("Open should have failed with invalid opts.")
	}
}

func TestOpenSession(t *testing.T) {
	conn, err := Open(".", "create")
	if err != nil {
		t.Errorf("Open failed with %v", err)
	}
	defer conn.Close()
	_, err = conn.OpenSession("isolation=snapshot")
	if err != nil {
		t.Errorf("Failed to open session: %v", err)
	}
}

func TestOpenSessionInvalidOpts(t *testing.T) {
	conn, err := Open(".", "create")
	if err != nil {
		t.Errorf("Open failed with %v", err)
	}
	defer conn.Close()
	_, err = conn.OpenSession("isolation=bananas")
	if err == nil {
		t.Errorf("Open session should have failed")
	}
}
