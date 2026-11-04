package internal_test

import (
	"os"
	"testing"
	"wc/internal"
)

func TestByteCounting(t *testing.T) {
	file := openTestFile(t, "../tests/test.txt")

	// wc -c tests/test.txt
	expected := 342190
	got, err := internal.CountBytes(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if expected != got {
		t.Fatalf("bytes counted incorrectly: expected %d, got %d", expected, got)
	}
}

func TestLineCounting(t *testing.T) {
	file := openTestFile(t, "../tests/test.txt")

	// wc -l tests/test.txt
	expected := 7145
	got, err := internal.CountLines(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if expected != got {
		t.Fatalf("lines counted incorrectly: expected %d, got %d", expected, got)
	}
}

func TestWordCounting(t *testing.T) {
	file := openTestFile(t, "../tests/test.txt")

	// wc -w tests/test.txt
	expected := 58164
	got, err := internal.CountWords(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if expected != got {
		t.Fatalf("words counted incorrectly: expected %d, got %d", expected, got)
	}
}

func TestCharCounting(t *testing.T) {
	files := []string{"../tests/test.txt", "../tests/multibyte.txt"}
	// wc -m tests/test.txt
	// wc -m tests/multibyte.txt
	expected := []int{339292, 113}

	for i, fn := range files {
		file := openTestFile(t, fn)
		expected := expected[i]
		got, err := internal.CountChars(file)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if expected != got {
			t.Fatalf("chars counted incorrectly: expected %d, got %d", expected, got)
		}
	}
}

func TestEverythingCounting(t *testing.T) {
	files := []string{"../tests/test.txt", "../tests/multibyte.txt"}
	// wc -clwm tests/test.txt
	// wc -clwm tests/multibyte.txt
	expected := []internal.Result{
		internal.Result{
			"bytes": 342190,
			"lines": 7145,
			"words": 58164,
			"chars": 339292,
		},
		internal.Result{
			"bytes": 285,
			"lines": 11,
			"words": 11,
			"chars": 113,
		},
	}

	for i, fn := range files {
		file := openTestFile(t, fn)
		expected := expected[i]
		got, err := internal.CountEverything(file)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for k, _ := range expected {
			if expected[k] != got[k] {
				t.Fatalf("%s counted incorrectly: expected %d, got %d", k, expected[k], got[k])
			}
		}
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
