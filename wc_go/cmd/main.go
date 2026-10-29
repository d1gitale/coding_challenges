package main

import (
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
		var (
			err  error
			read int
		)
		buf := make([]byte, 32*1024)
		lines := 0
		for {
			read, err = file.Read(buf)
			if err != nil {
				break
			}

			lines += bytes.Count(buf[:read], []byte("\n"))
		}

		if err == io.EOF {
			fmt.Printf("%d %s\n", lines, fn)
		} else {
			log.Fatalf("unexpected error: %v", err)
		}
	}
}
