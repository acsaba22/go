package prof

import (
	"strings"
	"unicode"
)

type WordFreq struct {
	words map[string]int
}

func (wf *WordFreq) AddWords(s string) bool {
	if wf.words == nil {
		wf.words = make(map[string]int)
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
		w := strings.ToLower(s[start:len(s)])
		wf.words[w]++
	}
	return false
}

func (wf WordFreq) GetWordCount(word string) int {
	return wf.words[strings.ToLower(word)]
}

type FastWF struct {
	WordFreq
}

func (w *FastWF) Write(p []byte) (n int, err error) {
	s := string(p)
	w.AddWords(s)
	return len(p), nil
}
