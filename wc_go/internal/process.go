package internal

import (
	"fmt"
	"io"
	"os"
)

func CountBytes(buf []byte, file *os.File) (int, error) {
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
