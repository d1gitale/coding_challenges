package cmd

import "flag"

type Flags struct {
	CountBytes *bool
	CountLines *bool
	CountWords *bool
	CountChars *bool
}

func ProcessFlags() *Flags {
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

	return &Flags{
		countBytes,
		countLines,
		countWords,
		countChars,
	}
}
