package internal_test

import (
	"os"
	"testing"
	"wc/internal"
)

func TestByteCounting(t *testing.T) {
	file := openTestFile(t, "../tests/test.txt")

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
	expected := []internal.Result{
		internal.Result{
			Bytes: 342190,
			Lines: 7145,
			Words: 58164,
			Chars: 339292,
		},
		internal.Result{
			Bytes: 285,
			Lines: 11,
			Words: 11,
			Chars: 113,
		},
	}

	for i, fn := range files {
		file := openTestFile(t, fn)
		expected := expected[i]
		got, err := internal.CountEverything(file)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if expected.Bytes != got.Bytes {
			t.Fatalf("bytes counted incorrectly: expected %d, got %d", expected.Bytes, got.Bytes)
		}
		if expected.Lines != got.Lines {
			t.Fatalf("lines counted incorrectly: expected %d, got %d", expected.Lines, got.Lines)
		}
		if expected.Words != got.Words {
			t.Fatalf("words counted incorrectly: expected %d, got %d", expected.Words, got.Words)
		}
		if expected.Chars != got.Chars {
			t.Fatalf("chars counted incorrectly: expected %d, got %d", expected.Chars, got.Chars)
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
