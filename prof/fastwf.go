package prof

import (
	"bytes"
	"strings"
	"unicode"
)

type WordFreq struct {
	m map[string]int
}

func (wf *WordFreq) AddWords(s string) {
	if wf.m == nil {
		wf.m = map[string]int{}
	}
	buf := bytes.Buffer{}
	for _, r := range s {
		if unicode.IsLetter(r) {
			buf.WriteRune(unicode.ToLower(r))
		} else {
			if buf.Len() != 0 {
				wf.m[buf.String()]++
				buf.Reset()
			}
		}
	}
	if buf.Len() != 0 {
		wf.m[buf.String()]++
		buf.Reset()
	}
}

func (wf *WordFreq) GetWordCount(word string) int {
	return wf.m[strings.ToLower(word)]
}

func (w *WordFreq) Write(p []byte) (n int, err error) {
	s := string(p)
	w.AddWords(s)
	return len(p), nil
}

type FastWF struct {
	words map[string]int
}

func (wf *FastWF) AddWords(s string) bool {
	if wf.words == nil {
		wf.words = map[string]int{}
	}
	start := -1
	for i, r := range s {
		if unicode.IsLetter(r) {
			if start < 0 {
				start = i
			}
		} else {
			if start >= 0 {
				w := strings.ToLower(s[start:i])
				wf.words[w]++
			}
			start = -1
		}
	}
	if start >= 0 {
		w := strings.ToLower(s[start:])
		wf.words[w]++
	}
	return false
}

func (wf FastWF) GetWordCount(word string) int {
	return wf.words[strings.ToLower(word)]
}

func (w *FastWF) Write(p []byte) (n int, err error) {
	s := string(p)
	w.AddWords(s)
	return len(p), nil
}
