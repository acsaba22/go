package wordfreq

import (
	"bytes"
	"sort"
	"strings"
	"unicode"
)

// Options:
// - make zero viable value
// - make a new operator and document to use it
// - make the type lowercase and make an uppercase interface

type WordFreq struct {
	words map[string]int
}

func splitToWords(s string) []string {
	buf := bytes.Buffer{}
	ret := []string{}
	for _, r := range s {
		switch {
		case unicode.IsLetter(r) || unicode.IsNumber(r):
			buf.WriteRune(unicode.ToLower(r))
		case 0 < buf.Len():
			ret = append(ret, buf.String())
			buf.Reset()
		}
	}
	return ret
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

type Word struct {
	w string
	c int
}

// TODO add unit test
func (wf WordFreq) WordsByCount() (s []Word) {
	for k, v := range wf.words {
		s = append(s, Word{k, v})
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].c > s[j].c
	})
	return s
}

// type WordCounts []struct {
// 	W string
// 	C int
// }

// func (wf WordFreq) WordCounts() WordCounts {
// 	for
// 	words := splitToWords(sentence)
// 	for _, w := range words {
// 		wf.words[w]++
// 	}
// 	return
// }
