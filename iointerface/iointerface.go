package main

import (
	"bufio"
	"fmt"
	"io"
)

type stringReader struct {
	s string
	i int
}

// go doc io.Reader
// copy(dest, src) has a special case for copy(byte[], string)

func toReader(s string) io.Reader {
	return &stringReader{s, 0}
	// return strings.NewReader(s)
}

func main() {
	var ior io.Reader
	ior = toReader("Let's make tokens from this sentence.")
	scanner := bufio.NewScanner(ior)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
