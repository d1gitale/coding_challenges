package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	countBytes := flag.Bool("c", false, "wrd -c test.txt")
	countLines := flag.Bool("l", false, "wrd -l test.txt")
	countWords := flag.Bool("w", false, "wrd -w test.txt")
	countChars := flag.Bool("m", false, "wrd -m test.txt")
	flag.Parse()

	if flag.NFlag() == 0 {
		*countBytes = true
		*countLines = true
		*countWords = true
	}

	var (
		file *os.File
		fn   string
		err  error
	)
	if flag.NArg() != 0 {
		fn := flag.Arg(0)

		file, err = os.Open(fn)
		if err != nil {
			log.Fatalf("failed to open file %s: %v", fn, err)
		}
		defer file.Close()
	} else {
		file = os.Stdin
		fn = "stdin"
	}

	var (
		read   int
		output string
	)

	buf := make([]byte, 32*1024)

	if *countBytes {
		size := 0
		var n int

		for {
			n, err = file.Read(buf)
			if err != nil {
				break
			}

			size += n
		}

		if err == io.EOF {
			output += fmt.Sprintf("%d ", size)
		} else {
			log.Fatalf("failed to read file %s: %v", fn, err)
		}
	}
	if *countLines {
		file.Seek(0, io.SeekStart)
		lines := 0
		for {
			read, err = file.Read(buf)
			if err != nil {
				break
			}

			lines += bytes.Count(buf[:read], []byte("\n"))
		}

		if err == io.EOF {
			output += fmt.Sprintf("%d ", lines)
		} else {
			log.Fatalf("failed to read file %s: %v", fn, err)
		}
	}
	if *countWords {
		file.Seek(0, io.SeekStart)

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		words := 0
		for scanner.Scan() {
			words++
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("failed to read the file %s: %v", fn, err)
		}

		output += fmt.Sprintf("%d ", words)
	}
	if *countChars {
		file.Seek(0, io.SeekStart)
		// TODO: change the implementation to chunk-based reading
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanRunes)

		chars := 0
		for scanner.Scan() {
			chars++
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("failed to read the file %s: %v", fn, err)
		}

		output += fmt.Sprintf("%d ", chars)
	}

	output += fmt.Sprintf("%s\n", fn)
	fmt.Print(output)
}
