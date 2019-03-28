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
	start := 0
	finish := -1
	for i, r := range s {
		if unicode.IsLetter(r) {
			finish = i
		} else {
			if start <= finish {
				w := strings.ToLower(s[start : finish+1])
				wf.words[w]++
			}
			start = i + 1
		}
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
