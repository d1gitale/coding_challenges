package internal_test

import (
	"os"
	"testing"
	"wc/internal"
)

func TestByteCounting(t *testing.T) {
	buf := make([]byte, 32*1024)
	file := openTestFile(t, "../tests/test.txt")

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

func TestLineCounting(t *testing.T) {
	buf := make([]byte, 32*1024)
	file := openTestFile(t, "../tests/test.txt")

	// wc -l tests/test.txt
	expected := 7145
	got, err := internal.CountLines(buf, file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if expected != got {
		t.Fatalf("lines counted incorrectly: expected %d, got %d", expected, got)
	}
}

func openTestFile(t *testing.T, filename string) *os.File {
	t.Helper()

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}

	t.Cleanup(func() {
		file.Close()
	})

	return file
}
