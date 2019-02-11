package wordfreq

import (
	"bytes"
	"strings"
	"unicode"
)

// Make:
// type WordFreq struct {...}
//
// And make the test compile, then pass.
//
// To split a string into pieces you can iterate over:
// buf := bytes.Buffer{}
// for over string
//   if unicode.IsLetter(r)
//      buf.WriteRune(unicode.ToLower(r))
//   else
//      we have one token in buf.String
//      buf.Reset

type WordFreq struct {
	words map[string]int
}

func (wf *WordFreq) AddWords(sentence string) bool {
	if wf.words == nil {
		wf.words = make(map[string]int)
	}
	words := splitToWords(sentence)
	for _, w := range words {
		wf.words[w]++
	}
	return false
}

func (wf WordFreq) GetWordCount(word string) int {
	return wf.words[strings.ToLower(word)]
}

func splitToWords(s string) []string {
	buf := bytes.Buffer{}
	ret := []string{}
	for _, r := range s {
		switch {
		case unicode.IsLetter(r):
			buf.WriteRune(unicode.ToLower(r))
		case 0 < buf.Len():
			ret = append(ret, buf.String())
			buf.Reset()
		}
	}
	return ret
}
