package pkg

import (
	"flag"
	"fmt"
	"os"
)

func OpenFileOrStdin() (*os.File, error) {
	var (
		file *os.File
		err  error
	)

	if flag.NArg() != 0 {
		fn := flag.Arg(0)

		file, err = os.Open(fn)
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s: %v", fn, err)
		}
	} else {
		file = os.Stdin
	}

	return file, nil
}
