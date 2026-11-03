package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"wc/internal"
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
		fn = flag.Arg(0)

		file, err = os.Open(fn)
		if err != nil {
			log.Fatalf("failed to open file %s: %v", fn, err)
		}
		defer file.Close()
	} else {
		file = os.Stdin
		fn = "stdin"
	}

	var output string
	buf := make([]byte, 32*1024)

	if *countBytes {
		size, err := internal.CountBytes(buf, file)
		if err != nil {
			log.Fatalf("failed to count bytes: %v", err)
		}

		output += fmt.Sprintf("%d ", size)
	}
	if *countLines {
		lines, err := internal.CountLines(buf, file)
		if err != nil {
			log.Fatalf("failed to count lines: %v", err)
		}

		output += fmt.Sprintf("%d ", lines)
	}
	if *countWords {
		words, err := internal.CountWords(file)
		if err != nil {
			log.Fatalf("failed to count words: %v", err)
		}

		output += fmt.Sprintf("%d ", words)
	}
	if *countChars {
		chars, err := internal.CountChars(file)
		if err != nil {
			log.Fatalf("failed to count chars: %v", err)
		}

		output += fmt.Sprintf("%d ", chars)
	}

	output += fmt.Sprintf("%s\n", fn)
	fmt.Print(output)
}
