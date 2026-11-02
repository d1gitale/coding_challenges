package internal_test

import (
	"os"
	"testing"
	"wc/internal"
)

func ByteCountingTest(t *testing.T) {
	buf := make([]byte, 32*1024)
	file, err := os.Open("tests/test.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// wc -c tests/test.txt
	expected := 342190
	got, err := internal.CountBytes(buf, file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if expected != got {
		t.Fatalf("bytes counted incorrectly: expected %d, got %d", expected, got)
	}
}
