package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"wc/cmd"
	"wc/internal"
)

func main() {
	log.SetFlags(0)

	flags := cmd.ProcessFlags()

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

	if *flags.CountBytes {
		size, err := internal.CountBytes(file)
		if err != nil {
			log.Fatalf("failed to count bytes: %v", err)
		}

		output += fmt.Sprintf("%d ", size)
	}
	if *flags.CountLines {
		lines, err := internal.CountLines(file)
		if err != nil {
			log.Fatalf("failed to count lines: %v", err)
		}

		output += fmt.Sprintf("%d ", lines)
	}
	if *flags.CountWords {
		words, err := internal.CountWords(file)
		if err != nil {
			log.Fatalf("failed to count words: %v", err)
		}

		output += fmt.Sprintf("%d ", words)
	}
	if *flags.CountChars {
		chars, err := internal.CountChars(file)
		if err != nil {
			log.Fatalf("failed to count chars: %v", err)
		}

		output += fmt.Sprintf("%d ", chars)
	}

	output += fmt.Sprintf("%s\n", fn)
	fmt.Print(output)
}
