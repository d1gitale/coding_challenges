package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	countBytes := flag.Bool("c", false, "wrd -c test.txt")
	flag.Parse()

	fn := flag.Arg(0)

	if *countBytes {
		file, err := os.Open(fn)
		if err != nil {
			log.Fatalf("failed to open file %s: %v", fn, err)
		}
		f, err := file.Stat()
		if err != nil {
			log.Fatalf("unexpected error: %v", err)
		}
		size := f.Size()

		fmt.Printf("%d %s\n", size, fn)
	}
}
