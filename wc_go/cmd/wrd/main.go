package main

import (
	"fmt"
	"log"
	"os"
	"wc/cmd"
	"wc/internal"
	"wc/pkg"
)

func main() {
	log.SetFlags(0)

	flags := cmd.ProcessFlags()

	file, err := pkg.OpenFileOrStdin()
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	defer file.Close()

	var output string

	if file == os.Stdin {
		res, err := internal.CountEverything(file)
		if err != nil {
			log.Fatalf("failed to count everything: %v", err)
		}

		if *flags.CountBytes {
			output += fmt.Sprintf("%d ", res["bytes"])
		}
		if *flags.CountLines {
			output += fmt.Sprintf("%d ", res["lines"])
		}
		if *flags.CountWords {
			output += fmt.Sprintf("%d ", res["words"])
		}
		if *flags.CountChars {
			output += fmt.Sprintf("%d ", res["chars"])
		}
	} else {
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
	}

	output += fmt.Sprintf("%s\n", file.Name())
	fmt.Print(output)
}
