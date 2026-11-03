package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func CountBytes(buf []byte, file *os.File) (int, error) {
	file.Seek(0, io.SeekStart)
	size := 0
	var (
		n   int
		err error
	)

	for {
		n, err = file.Read(buf)
		if err != nil {
			break
		}

		size += n
	}

	if err == io.EOF {
		return size, nil
	} else {
		return 0, fmt.Errorf("failed to read file %s: %v", file.Name(), err)
	}
}

func CountLines(buf []byte, file *os.File) (int, error) {
	file.Seek(0, io.SeekStart)

	var (
		read int
		err  error
	)

	lines := 0
	for {
		read, err = file.Read(buf)
		if err != nil {
			break
		}

		lines += bytes.Count(buf[:read], []byte("\n"))
	}

	if err == io.EOF {
		return lines, nil
	} else {
		return 0, fmt.Errorf("failed to read file %s: %v", file.Name(), err)
	}
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
