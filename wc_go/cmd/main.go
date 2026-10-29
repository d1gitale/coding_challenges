package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	countBytes := flag.Bool("c", false, "wrd -c test.txt")
	countLines := flag.Bool("l", false, "wrd -l test.txt")
	flag.Parse()

	fn := flag.Arg(0)

	file, err := os.Open(fn)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", fn, err)
	}
	defer file.Close()
	fileStat, err := file.Stat()
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	switch {
	case *countBytes:
		size := fileStat.Size()

		fmt.Printf("%d %s\n", size, fn)
	case *countLines:
		data, err := os.ReadFile(fn)
		if err != nil {
			log.Fatalf("failed to read file %s: %v", fn, err)
		}
		lines := bytes.Count(data, []byte("\n"))

		fmt.Printf("%d %s\n", lines, fn)
	}
}
