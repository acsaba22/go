package iowordfreq

import (
	"fmt"
	"testing"
)

func ExampleIoWordFreq() {
	iowf := IoWordFreq{}
	iowf.AddWords("hello friend")
	fmt.Println(iowf.GetWordCount("hello"))

	fmt.Fprintln(&iowf, "and hello again")
	fmt.Println(iowf.GetWordCount("hello"))

	fmt.Fprintln(&iowf, "and hello again again")
	fmt.Println(iowf.GetWordCount("hello"))
	// Output:
	// 1
	// 2
	// 3
}

func TestMapRelocation(t *testing.T) {
	iowf := IoWordFreq{}
	fmt.Fprintln(&iowf, "one")
	if iowf.GetWordCount("one") != 1 {
		t.Errorf("one problem")
	}
}
