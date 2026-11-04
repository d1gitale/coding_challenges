package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type Result struct {
	Bytes int
	Lines int
	Words int
	Chars int
}

func CountBytes(file *os.File) (int, error) {
	file.Seek(0, io.SeekStart)

	size := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		size += len(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed to read file %s: %v", file.Name(), err)
	}

	return size, nil
}

func CountLines(file *os.File) (int, error) {
	file.Seek(0, io.SeekStart)

	lines := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed to read file %s: %v", file.Name(), err)
	}

	return lines, nil
}

func CountWords(file *os.File) (int, error) {
	file.Seek(0, io.SeekStart)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	words := 0
	for scanner.Scan() {
		words++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed to read the file %s: %v", file.Name(), err)
	}

	return words, nil
}

func CountChars(file *os.File) (int, error) {
	file.Seek(0, io.SeekStart)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	chars := 0
	for scanner.Scan() {
		chars++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed to read the file %s: %v", file.Name(), err)
	}

	return chars, nil
}

func CountEverything(file *os.File) (Result, error) {
	file.Seek(0, io.SeekStart)

	reader := bufio.NewReader(file)

	bytes := 0
	lines := 0
	words := 0
	chars := 0
	inWord := false

	for {
		r, n, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return Result{}, fmt.Errorf("failed to read rune: %v", err)
		}

		bytes += n

		if r == '\n' {
			lines++
		}

		if unicode.IsSpace(r) {
			inWord = false
		} else {
			if !inWord {
				words++
				inWord = true
			}
		}

		chars++
	}

	return Result{
		Bytes: bytes,
		Lines: lines,
		Words: words,
		Chars: chars,
	}, nil
}
