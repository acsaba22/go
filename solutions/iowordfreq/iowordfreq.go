package iowordfreq

import (
	"github.com/acsaba22/go/solutions/wordfreq"
)

type IoWordFreq struct {
	wordfreq.WordFreq
}

func (w *IoWordFreq) Write(p []byte) (n int, err error) {
	w.AddWords(string(p))
	return len(p), nil
}
