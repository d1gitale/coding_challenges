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

	var read int

	buf := make([]byte, 32*1024)

	switch {
	case *countBytes:
		size := fileStat.Size()

		fmt.Printf("%d %s\n", size, fn)
	case *countLines:
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
	case *countWords:
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		words := 0
		for scanner.Scan() {
			words++
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("failed to read the file %s: %v", fn, err)
		}

		fmt.Printf("%d %s\n", words, fn)
	}
}
