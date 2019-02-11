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

func (sr *stringReader) Read(p []byte) (n int, err error) {
	if len(sr.s) <= sr.i {
		err = io.EOF
		return
	}

	n = len(sr.s) - sr.i
	if len(p) < n {
		n = len(p)
	}
	copy(p, sr.s[sr.i:sr.i+n])
	sr.i += n
	return
}

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
